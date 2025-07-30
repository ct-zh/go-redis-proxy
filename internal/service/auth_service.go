package service

import (
	"context"

	"github.com/ct-zh/go-redis-proxy/pkg/types"
)

// AuthService 定义认证服务接口
type AuthService interface {
	// Connect 建立Redis连接并返回token
	Connect(ctx context.Context, req *types.ConnectRequest) (*types.ConnectResponse, error)
	
	// RefreshToken 刷新token的过期时间
	RefreshToken(ctx context.Context, req *types.TokenRefreshRequest) (*types.ConnectResponse, error)
	
	// Disconnect 断开连接并清理资源
	Disconnect(ctx context.Context, req *types.DisconnectRequest) error
	
	// GetStats 获取连接池统计信息
	GetStats(ctx context.Context) (*types.ConnectionStats, error)
	
	// ValidateToken 验证token是否有效
	ValidateToken(ctx context.Context, token string) error
}

// TokenRedisService 定义基于token的Redis服务接口
type TokenRedisService interface {
	// String operations with token
	TokenStringGet(ctx context.Context, req *types.TokenStringGetRequest) (*types.StringGetData, error)
	TokenStringSet(ctx context.Context, req *types.TokenStringSetRequest) (*types.StringSetData, error)
	TokenStringDel(ctx context.Context, req *types.TokenStringDelRequest) (*types.StringDelData, error)
	TokenStringExists(ctx context.Context, req *types.TokenStringExistsRequest) (*types.StringExistsData, error)
	TokenStringIncr(ctx context.Context, req *types.TokenStringIncrRequest) (*types.StringIncrData, error)
	TokenStringDecr(ctx context.Context, req *types.TokenStringDecrRequest) (*types.StringDecrData, error)
	TokenStringExpire(ctx context.Context, req *types.TokenStringExpireRequest) (*types.StringExpireData, error)
	
	// List operations with token
	TokenListLPush(ctx context.Context, req *types.TokenListLPushRequest) (*types.ListLPushData, error)
	
	// Hash operations with token
	TokenHashHSet(ctx context.Context, req *types.TokenHashHSetRequest) (*types.HashHSetData, error)
	TokenHashHGet(ctx context.Context, req *types.TokenHashHGetRequest) (*types.HashHGetData, error)
}