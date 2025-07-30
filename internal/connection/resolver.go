package connection

import (
	"context"
	"crypto/md5"
	"fmt"
)

// ConnectionConfig 统一的Redis连接配置
type ConnectionConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	DB       int    `json:"db"`
	
	// 连接池配置
	PoolSize     int `json:"pool_size,omitempty"`
	MinIdleConns int `json:"min_idle_conns,omitempty"`
	MaxRetries   int `json:"max_retries,omitempty"`
}

// Hash 生成配置的唯一标识，用于连接池复用
func (c *ConnectionConfig) Hash() string {
	data := fmt.Sprintf("%s:%d:%s:%d:%d:%d:%d", 
		c.Host, c.Port, c.Password, c.DB, 
		c.PoolSize, c.MinIdleConns, c.MaxRetries)
	return fmt.Sprintf("%x", md5.Sum([]byte(data)))
}

// Validate 验证连接配置的有效性
func (c *ConnectionConfig) Validate() error {
	if c.Host == "" {
		return fmt.Errorf("host is required")
	}
	if c.Port <= 0 || c.Port > 65535 {
		return fmt.Errorf("invalid port: %d", c.Port)
	}
	if c.DB < 0 {
		return fmt.Errorf("invalid db: %d", c.DB)
	}
	return nil
}

// ConnectionResolver 连接配置解析器接口
type ConnectionResolver interface {
	// ResolveConnection 从请求中解析连接配置
	ResolveConnection(ctx context.Context, req interface{}) (*ConnectionConfig, error)
	
	// SupportsRequest 检查是否支持该类型的请求
	SupportsRequest(req interface{}) bool
}

// CompositeResolver 组合解析器，支持多种请求类型
type CompositeResolver struct {
	resolvers []ConnectionResolver
}

// NewCompositeResolver 创建组合解析器
func NewCompositeResolver(resolvers ...ConnectionResolver) *CompositeResolver {
	return &CompositeResolver{
		resolvers: resolvers,
	}
}

// ResolveConnection 尝试使用各个解析器解析连接配置
func (r *CompositeResolver) ResolveConnection(ctx context.Context, req interface{}) (*ConnectionConfig, error) {
	for _, resolver := range r.resolvers {
		if resolver.SupportsRequest(req) {
			return resolver.ResolveConnection(ctx, req)
		}
	}
	return nil, fmt.Errorf("unsupported request type: %T", req)
}

// SupportsRequest 检查是否有解析器支持该请求类型
func (r *CompositeResolver) SupportsRequest(req interface{}) bool {
	for _, resolver := range r.resolvers {
		if resolver.SupportsRequest(req) {
			return true
		}
	}
	return false
}