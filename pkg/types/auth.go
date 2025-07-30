package types

import "time"

// ConnectRequest 定义建立Redis连接的请求体
type ConnectRequest struct {
	RedisRequest
	TokenTTL int `json:"token_ttl,omitempty"` // Token有效期，单位秒，默认3600秒
}

// ConnectResponse 定义连接响应体
type ConnectResponse struct {
	Token     string    `json:"token"`      // 连接token
	ExpiresAt time.Time `json:"expires_at"` // token过期时间
	ConnID    string    `json:"conn_id"`    // 连接ID，用于调试和监控
}

// TokenRequest 定义使用token的请求基础结构
type TokenRequest struct {
	Token string `json:"token" binding:"required"`
}

// TokenRefreshRequest 定义token刷新请求
type TokenRefreshRequest struct {
	Token    string `json:"token" binding:"required"`
	TokenTTL int    `json:"token_ttl,omitempty"` // 新的TTL，单位秒
}

// DisconnectRequest 定义断开连接的请求体
type DisconnectRequest struct {
	Token string `json:"token" binding:"required"`
}

// ConnectionInfo 存储连接信息的内部结构
type ConnectionInfo struct {
	RedisConfig RedisRequest `json:"redis_config"`
	CreatedAt   time.Time    `json:"created_at"`
	ExpiresAt   time.Time    `json:"expires_at"`
	ConnID      string       `json:"conn_id"`
	LastUsed    time.Time    `json:"last_used"`
}

// ConnectionStats 连接统计信息
type ConnectionStats struct {
	TotalConnections   int `json:"total_connections"`
	ActiveConnections  int `json:"active_connections"`
	PoolSize          int `json:"pool_size"`
	IdleConnections   int `json:"idle_connections"`
}