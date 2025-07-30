package connection

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net"
	"reflect"
	"strconv"

	"github.com/ct-zh/go-redis-proxy/pkg/types"
)

// RedisRequestResolver RedisRequest类型的解析器
type RedisRequestResolver struct{}

// NewRedisRequestResolver 创建RedisRequest解析器
func NewRedisRequestResolver() *RedisRequestResolver {
	return &RedisRequestResolver{}
}

// SupportsRequest 检查是否支持该请求类型
func (r *RedisRequestResolver) SupportsRequest(req interface{}) bool {
	// 使用反射检查请求是否包含RedisRequest字段
	v := reflect.ValueOf(req)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return false
	}
	
	// 检查是否有RedisRequest字段
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		if field.Type == reflect.TypeOf(types.RedisRequest{}) {
			return true
		}
	}
	return false
}

// ResolveConnection 从RedisRequest中解析连接配置
func (r *RedisRequestResolver) ResolveConnection(ctx context.Context, req interface{}) (*ConnectionConfig, error) {
	redisReq, err := r.extractRedisRequest(req)
	if err != nil {
		return nil, fmt.Errorf("failed to extract RedisRequest: %w", err)
	}
	
	// 解析Addr字段，格式通常为 "host:port"
	host, port, err := parseAddr(redisReq.Addr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse addr %s: %w", redisReq.Addr, err)
	}
	
	config := &ConnectionConfig{
		Host:         host,
		Port:         port,
		Password:     redisReq.Password,
		DB:           redisReq.DB,
		PoolSize:     10, // 默认连接池大小
		MinIdleConns: 2,  // 默认最小空闲连接
		MaxRetries:   3,  // 默认重试次数
	}
	
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("invalid connection config: %w", err)
	}
	
	return config, nil
}

// extractRedisRequest 从请求中提取RedisRequest
func (r *RedisRequestResolver) extractRedisRequest(req interface{}) (*types.RedisRequest, error) {
	v := reflect.ValueOf(req)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("request must be a struct")
	}
	
	// 查找RedisRequest字段
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := v.Type().Field(i)
		
		if fieldType.Type == reflect.TypeOf(types.RedisRequest{}) {
			redisReq := field.Interface().(types.RedisRequest)
			return &redisReq, nil
		}
	}
	
	return nil, fmt.Errorf("RedisRequest field not found")
}

// TokenResolver Token类型的解析器
type TokenResolver struct {
	tokenStore TokenStore
}

// TokenStore Token存储接口
type TokenStore interface {
	// GetConnectionConfig 根据token获取连接配置
	GetConnectionConfig(ctx context.Context, token string) (*ConnectionConfig, error)
	
	// StoreConnectionConfig 存储token和连接配置的映射
	StoreConnectionConfig(ctx context.Context, token string, config *ConnectionConfig) error
	
	// DeleteToken 删除token
	DeleteToken(ctx context.Context, token string) error
	
	// IsValidToken 检查token是否有效
	IsValidToken(ctx context.Context, token string) bool
}

// NewTokenResolver 创建Token解析器
func NewTokenResolver(tokenStore TokenStore) *TokenResolver {
	return &TokenResolver{
		tokenStore: tokenStore,
	}
}

// SupportsRequest 检查是否支持该请求类型
func (r *TokenResolver) SupportsRequest(req interface{}) bool {
	// 使用反射检查请求是否包含TokenRequest字段
	v := reflect.ValueOf(req)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return false
	}
	
	// 检查是否有TokenRequest字段
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		if field.Type == reflect.TypeOf(types.TokenRequest{}) {
			return true
		}
	}
	return false
}

// ResolveConnection 从Token中解析连接配置
func (r *TokenResolver) ResolveConnection(ctx context.Context, req interface{}) (*ConnectionConfig, error) {
	tokenReq, err := r.extractTokenRequest(req)
	if err != nil {
		return nil, fmt.Errorf("failed to extract TokenRequest: %w", err)
	}
	
	if tokenReq.Token == "" {
		return nil, fmt.Errorf("token is required")
	}
	
	// 从token store中获取连接配置
	config, err := r.tokenStore.GetConnectionConfig(ctx, tokenReq.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to get connection config for token: %w", err)
	}
	
	return config, nil
}

// extractTokenRequest 从请求中提取TokenRequest
func (r *TokenResolver) extractTokenRequest(req interface{}) (*types.TokenRequest, error) {
	v := reflect.ValueOf(req)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("request must be a struct")
	}
	
	// 查找TokenRequest字段
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := v.Type().Field(i)
		
		if fieldType.Type == reflect.TypeOf(types.TokenRequest{}) {
			tokenReq := field.Interface().(types.TokenRequest)
			return &tokenReq, nil
		}
	}
	
	return nil, fmt.Errorf("TokenRequest field not found")
}

// GenerateToken 根据连接配置生成唯一token
func GenerateToken(config *ConnectionConfig) string {
	// 生成随机部分确保token唯一性
	randomBytes := make([]byte, 16)
	rand.Read(randomBytes)
	randomPart := hex.EncodeToString(randomBytes)
	
	// 使用连接配置的hash作为token的一部分，便于识别
	configHash := config.Hash()[:8]
	
	return fmt.Sprintf("token_%s_%s", configHash, randomPart)
}

// parseAddr 解析地址字符串，返回host和port
func parseAddr(addr string) (string, int, error) {
	host, portStr, err := net.SplitHostPort(addr)
	if err != nil {
		// 如果没有端口，默认使用6379
		return addr, 6379, nil
	}
	
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return "", 0, fmt.Errorf("invalid port: %s", portStr)
	}
	
	return host, port, nil
}