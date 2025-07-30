package connection

import (
	"context"

	"github.com/go-redis/redis/v8"
)

// Service 连接服务，提供统一的连接管理接口
type Service struct {
	resolver *CompositeResolver
	manager  *ConnectionManager
}

// NewService 创建连接服务
func NewService() *Service {
	// 创建token存储
	tokenStore := NewMemoryTokenStore()
	
	// 创建解析器
	redisResolver := NewRedisRequestResolver()
	tokenResolver := NewTokenResolver(tokenStore)
	compositeResolver := NewCompositeResolver(redisResolver, tokenResolver)
	
	// 创建连接管理器
	manager := NewConnectionManager()
	
	return &Service{
		resolver: compositeResolver,
		manager:  manager,
	}
}

// GetConnection 从请求中解析连接配置并获取Redis连接
func (s *Service) GetConnection(ctx context.Context, req interface{}) (*redis.Client, error) {
	// 解析连接配置
	config, err := s.resolver.ResolveConnection(ctx, req)
	if err != nil {
		return nil, err
	}
	
	// 获取连接
	return s.manager.GetConnection(config)
}

// CreateToken 根据连接配置创建token
func (s *Service) CreateToken(ctx context.Context, req interface{}) (string, error) {
	// 解析连接配置
	config, err := s.resolver.ResolveConnection(ctx, req)
	if err != nil {
		return "", err
	}
	
	// 生成token
	token := GenerateToken(config)
	
	// 存储token和配置的映射
	tokenResolver := s.getTokenResolver()
	if tokenResolver != nil {
		err = tokenResolver.tokenStore.StoreConnectionConfig(ctx, token, config)
		if err != nil {
			return "", err
		}
	}
	
	return token, nil
}

// DeleteToken 删除token
func (s *Service) DeleteToken(ctx context.Context, token string) error {
	tokenResolver := s.getTokenResolver()
	if tokenResolver != nil {
		return tokenResolver.tokenStore.DeleteToken(ctx, token)
	}
	return nil
}

// IsValidToken 检查token是否有效
func (s *Service) IsValidToken(ctx context.Context, token string) bool {
	tokenResolver := s.getTokenResolver()
	if tokenResolver != nil {
		return tokenResolver.tokenStore.IsValidToken(ctx, token)
	}
	return false
}

// GetStats 获取连接服务统计信息
func (s *Service) GetStats() map[string]interface{} {
	stats := make(map[string]interface{})
	
	// 连接管理器统计
	stats["connection_manager"] = s.manager.GetStats()
	
	// Token存储统计
	tokenResolver := s.getTokenResolver()
	if tokenResolver != nil {
		if memoryStore, ok := tokenResolver.tokenStore.(*MemoryTokenStore); ok {
			stats["token_store"] = memoryStore.GetStats()
		}
	}
	
	return stats
}

// HealthCheck 健康检查
func (s *Service) HealthCheck(ctx context.Context) map[string]interface{} {
	result := make(map[string]interface{})
	
	// 连接管理器健康检查
	result["connections"] = s.manager.HealthCheck(ctx)
	
	// Token存储健康检查
	tokenResolver := s.getTokenResolver()
	if tokenResolver != nil {
		if memoryStore, ok := tokenResolver.tokenStore.(*MemoryTokenStore); ok {
			result["token_store"] = map[string]interface{}{
				"status": "healthy",
				"stats":  memoryStore.GetStats(),
			}
		}
	}
	
	return result
}

// Close 关闭连接服务
func (s *Service) Close() error {
	return s.manager.CloseAll()
}

// getTokenResolver 获取TokenResolver实例
func (s *Service) getTokenResolver() *TokenResolver {
	for _, resolver := range s.resolver.resolvers {
		if tokenResolver, ok := resolver.(*TokenResolver); ok {
			return tokenResolver
		}
	}
	return nil
}

// SupportsRequest 检查是否支持该请求类型
func (s *Service) SupportsRequest(req interface{}) bool {
	return s.resolver.SupportsRequest(req)
}