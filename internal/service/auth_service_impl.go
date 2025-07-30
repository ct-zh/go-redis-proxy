package service

import (
	"context"
	"time"

	"github.com/ct-zh/go-redis-proxy/pkg/errors"
	"github.com/ct-zh/go-redis-proxy/pkg/types"
)

// AuthServiceImpl 实现认证服务
type AuthServiceImpl struct {
	connectionManager *ConnectionManager
}

// NewAuthService 创建新的认证服务实例
func NewAuthService(connectionManager *ConnectionManager) AuthService {
	return &AuthServiceImpl{
		connectionManager: connectionManager,
	}
}

// Connect 建立Redis连接并返回token
func (s *AuthServiceImpl) Connect(ctx context.Context, req *types.ConnectRequest) (*types.ConnectResponse, error) {
	return s.connectionManager.Connect(ctx, req)
}

// RefreshToken 刷新token的过期时间
func (s *AuthServiceImpl) RefreshToken(ctx context.Context, req *types.TokenRefreshRequest) (*types.ConnectResponse, error) {
	return s.connectionManager.RefreshToken(req)
}

// Disconnect 断开连接并清理资源
func (s *AuthServiceImpl) Disconnect(ctx context.Context, req *types.DisconnectRequest) error {
	return s.connectionManager.Disconnect(req.Token)
}

// GetStats 获取连接池统计信息
func (s *AuthServiceImpl) GetStats(ctx context.Context) (*types.ConnectionStats, error) {
	return s.connectionManager.GetStats(), nil
}

// ValidateToken 验证token是否有效
func (s *AuthServiceImpl) ValidateToken(ctx context.Context, token string) error {
	_, err := s.connectionManager.GetClient(token)
	return err
}

// TokenRedisServiceImpl 实现基于token的Redis服务
type TokenRedisServiceImpl struct {
	connectionManager *ConnectionManager
}

// NewTokenRedisService 创建新的基于token的Redis服务实例
func NewTokenRedisService(connectionManager *ConnectionManager) TokenRedisService {
	return &TokenRedisServiceImpl{
		connectionManager: connectionManager,
	}
}

// TokenStringGet 基于token的GET操作
func (s *TokenRedisServiceImpl) TokenStringGet(ctx context.Context, req *types.TokenStringGetRequest) (*types.StringGetData, error) {
	client, err := s.connectionManager.GetClient(req.Token)
	if err != nil {
		return nil, err
	}

	result := client.Get(ctx, req.Key)
	if result.Err() != nil {
		if result.Err().Error() == "redis: nil" {
			return &types.StringGetData{Value: nil}, nil
		}
		return nil, errors.NewError(errors.CodeStringGetFailed)
	}

	return &types.StringGetData{Value: result.Val()}, nil
}

// TokenStringSet 基于token的SET操作
func (s *TokenRedisServiceImpl) TokenStringSet(ctx context.Context, req *types.TokenStringSetRequest) (*types.StringSetData, error) {
	client, err := s.connectionManager.GetClient(req.Token)
	if err != nil {
		return nil, err
	}

	var ttl time.Duration
	if req.TTL > 0 {
		ttl = time.Duration(req.TTL) * time.Second
	}

	result := client.Set(ctx, req.Key, req.Value, ttl)
	if result.Err() != nil {
		return nil, errors.NewError(errors.CodeStringSetFailed)
	}

	return &types.StringSetData{Result: result.Val()}, nil
}

// TokenStringDel 基于token的DEL操作
func (s *TokenRedisServiceImpl) TokenStringDel(ctx context.Context, req *types.TokenStringDelRequest) (*types.StringDelData, error) {
	client, err := s.connectionManager.GetClient(req.Token)
	if err != nil {
		return nil, err
	}

	result := client.Del(ctx, req.Key)
	if result.Err() != nil {
		return nil, errors.NewError(errors.CodeStringDelFailed)
	}

	return &types.StringDelData{Deleted: result.Val()}, nil
}

// TokenStringExists 基于token的EXISTS操作
func (s *TokenRedisServiceImpl) TokenStringExists(ctx context.Context, req *types.TokenStringExistsRequest) (*types.StringExistsData, error) {
	client, err := s.connectionManager.GetClient(req.Token)
	if err != nil {
		return nil, err
	}

	result := client.Exists(ctx, req.Key)
	if result.Err() != nil {
		return nil, errors.NewError(errors.CodeStringGetFailed)
	}

	return &types.StringExistsData{Exists: result.Val() > 0}, nil
}

// TokenStringIncr 基于token的INCR操作
func (s *TokenRedisServiceImpl) TokenStringIncr(ctx context.Context, req *types.TokenStringIncrRequest) (*types.StringIncrData, error) {
	client, err := s.connectionManager.GetClient(req.Token)
	if err != nil {
		return nil, err
	}

	result := client.Incr(ctx, req.Key)
	if result.Err() != nil {
		return nil, errors.NewError(errors.CodeStringIncrFailed)
	}

	return &types.StringIncrData{Value: result.Val()}, nil
}

// TokenStringDecr 基于token的DECR操作
func (s *TokenRedisServiceImpl) TokenStringDecr(ctx context.Context, req *types.TokenStringDecrRequest) (*types.StringDecrData, error) {
	client, err := s.connectionManager.GetClient(req.Token)
	if err != nil {
		return nil, err
	}

	result := client.Decr(ctx, req.Key)
	if result.Err() != nil {
		return nil, errors.NewError(errors.CodeStringDecrFailed)
	}

	return &types.StringDecrData{Value: result.Val()}, nil
}

// TokenStringExpire 基于token的EXPIRE操作
func (s *TokenRedisServiceImpl) TokenStringExpire(ctx context.Context, req *types.TokenStringExpireRequest) (*types.StringExpireData, error) {
	client, err := s.connectionManager.GetClient(req.Token)
	if err != nil {
		return nil, err
	}

	ttl := time.Duration(req.TTL) * time.Second
	result := client.Expire(ctx, req.Key, ttl)
	if result.Err() != nil {
		return nil, errors.NewError(errors.CodeStringExpireFailed)
	}

	return &types.StringExpireData{Success: result.Val()}, nil
}

// TokenListLPush 基于token的LPUSH操作
func (s *TokenRedisServiceImpl) TokenListLPush(ctx context.Context, req *types.TokenListLPushRequest) (*types.ListLPushData, error) {
	client, err := s.connectionManager.GetClient(req.Token)
	if err != nil {
		return nil, err
	}

	if len(req.Values) == 0 {
		return &types.ListLPushData{Length: 0}, nil
	}

	// Convert []string to []interface{}
	interfaces := make([]interface{}, len(req.Values))
	for i, v := range req.Values {
		interfaces[i] = v
	}

	result := client.LPush(ctx, req.Key, interfaces...)
	if result.Err() != nil {
		return nil, errors.NewError(errors.CodeListPushFailed)
	}

	return &types.ListLPushData{Length: result.Val()}, nil
}

// TokenHashHSet 基于token的HSET操作
func (s *TokenRedisServiceImpl) TokenHashHSet(ctx context.Context, req *types.TokenHashHSetRequest) (*types.HashHSetData, error) {
	client, err := s.connectionManager.GetClient(req.Token)
	if err != nil {
		return nil, err
	}

	// Convert map[string]string to []interface{}
	fields := make([]interface{}, 0, len(req.Fields)*2)
	for field, value := range req.Fields {
		fields = append(fields, field, value)
	}

	result := client.HSet(ctx, req.Key, fields...)
	if result.Err() != nil {
		return nil, errors.NewError(errors.CodeHashSetFailed)
	}

	return &types.HashHSetData{Set: result.Val()}, nil
}

// TokenHashHGet 基于token的HGET操作
func (s *TokenRedisServiceImpl) TokenHashHGet(ctx context.Context, req *types.TokenHashHGetRequest) (*types.HashHGetData, error) {
	client, err := s.connectionManager.GetClient(req.Token)
	if err != nil {
		return nil, err
	}

	result := client.HGet(ctx, req.Key, req.Field)
	if result.Err() != nil {
		if result.Err().Error() == "redis: nil" {
			return &types.HashHGetData{Value: nil}, nil
		}
		return nil, errors.NewError(errors.CodeHashGetFailed)
	}

	return &types.HashHGetData{Value: result.Val()}, nil
}