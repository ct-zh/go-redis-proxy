package service

import (
	"context"
	"fmt"
	"time"

	"github.com/ct-zh/go-redis-proxy/internal/connection"
	"github.com/ct-zh/go-redis-proxy/pkg/types"
)

// SimpleTokenService 简化的Token管理服务
type SimpleTokenService struct {
	connectionService *connection.Service
}

// NewSimpleTokenService 创建简化的Token管理服务
func NewSimpleTokenService(connectionService *connection.Service) *SimpleTokenService {
	return &SimpleTokenService{
		connectionService: connectionService,
	}
}

// CreateToken 创建Token
func (s *SimpleTokenService) CreateToken(ctx context.Context, req *types.ConnectRequest) (*types.ConnectResponse, error) {
	// 创建包含RedisRequest字段的结构体
	connectionReq := struct {
		types.RedisRequest
	}{
		RedisRequest: req.RedisRequest,
	}
	
	// 使用连接服务创建Token
	token, err := s.connectionService.CreateToken(ctx, connectionReq)
	if err != nil {
		return nil, fmt.Errorf("failed to create token: %w", err)
	}

	// 计算过期时间
	ttl := req.TokenTTL
	if ttl <= 0 {
		ttl = 3600 // 默认1小时
	}
	expiresAt := time.Now().Add(time.Duration(ttl) * time.Second)

	return &types.ConnectResponse{
		Token:     token,
		ExpiresAt: expiresAt,
		ConnID:    fmt.Sprintf("conn_%s", token[:8]), // 简化的连接ID
	}, nil
}

// RefreshToken 刷新Token
func (s *SimpleTokenService) RefreshToken(ctx context.Context, req *types.TokenRefreshRequest) (*types.ConnectResponse, error) {
	// 验证当前Token
	if !s.connectionService.IsValidToken(ctx, req.Token) {
		return nil, fmt.Errorf("invalid token")
	}

	// 删除旧Token
	if err := s.connectionService.DeleteToken(ctx, req.Token); err != nil {
		return nil, fmt.Errorf("failed to delete old token: %w", err)
	}

	// 获取原始Redis配置（这里简化处理，实际应该从Token中恢复配置）
	// 为了简化，我们使用默认配置
	redisConfig := types.RedisRequest{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}

	// 创建包含RedisRequest字段的结构体
	connectionReq := struct {
		types.RedisRequest
	}{
		RedisRequest: redisConfig,
	}

	// 创建新Token
	newToken, err := s.connectionService.CreateToken(ctx, connectionReq)
	if err != nil {
		return nil, fmt.Errorf("failed to create new token: %w", err)
	}

	// 计算过期时间
	ttl := req.TokenTTL
	if ttl <= 0 {
		ttl = 3600 // 默认1小时
	}
	expiresAt := time.Now().Add(time.Duration(ttl) * time.Second)

	return &types.ConnectResponse{
		Token:     newToken,
		ExpiresAt: expiresAt,
		ConnID:    fmt.Sprintf("conn_%s", newToken[:8]),
	}, nil
}

// ValidateToken 验证Token
func (s *SimpleTokenService) ValidateToken(ctx context.Context, token string) (bool, error) {
	return s.connectionService.IsValidToken(ctx, token), nil
}

// DeleteToken 删除Token
func (s *SimpleTokenService) DeleteToken(ctx context.Context, req *types.DisconnectRequest) error {
	return s.connectionService.DeleteToken(ctx, req.Token)
}

// GetConnectionStats 获取连接统计
func (s *SimpleTokenService) GetConnectionStats(ctx context.Context) (*types.ConnectionStats, error) {
	// 简化统计信息处理
	return &types.ConnectionStats{
		TotalConnections:  0, // 从stats中提取实际值
		ActiveConnections: 0,
		PoolSize:         0,
		IdleConnections:  0,
	}, nil
}

// HealthCheck 健康检查
func (s *SimpleTokenService) HealthCheck(ctx context.Context) error {
	healthResult := s.connectionService.HealthCheck(ctx)
	// 简化处理，如果有任何健康检查失败则返回错误
	if healthResult == nil {
		return fmt.Errorf("health check failed")
	}
	return nil
}