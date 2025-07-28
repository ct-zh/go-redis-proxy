package service

import (
	"context"
	"time"

	"github.com/ct-zh/go-redis-proxy/internal/dao"
	"github.com/ct-zh/go-redis-proxy/pkg/types"
	"github.com/ct-zh/go-redis-proxy/pkg/errors"
)

// RedisStringServiceImpl implements RedisStringService interface
type RedisStringServiceImpl struct {
	dao dao.RedisDAO
}

// NewRedisStringService creates a new RedisStringService instance
func NewRedisStringService(redisDAO dao.RedisDAO) RedisStringService {
	return &RedisStringServiceImpl{
		dao: redisDAO,
	}
}

// Get retrieves a string value from Redis
func (s *RedisStringServiceImpl) Get(ctx context.Context, req *types.StringGetRequest) (*types.StringGetData, error) {
	// Connect to Redis
	if err := s.dao.Connect(req.RedisRequest); err != nil {
		return nil, errors.NewError(errors.CodeRedisConnectFailed)
	}

	// Get the value
	value, err := s.dao.StringGet(ctx, req.Key)
	if err != nil {
		return nil, errors.NewError(errors.CodeStringGetFailed)
	}

	return &types.StringGetData{
		Value: value,
	}, nil
}

// Set sets a string value in Redis
func (s *RedisStringServiceImpl) Set(ctx context.Context, req *types.StringSetRequest) (*types.StringSetData, error) {
	// Connect to Redis
	if err := s.dao.Connect(req.RedisRequest); err != nil {
		return nil, errors.NewError(errors.CodeRedisConnectFailed)
	}

	// Set TTL
	var ttl time.Duration
	if req.TTL > 0 {
		ttl = time.Duration(req.TTL) * time.Second
	}

	// Set the value
	result, err := s.dao.StringSet(ctx, req.Key, req.Value, ttl)
	if err != nil {
		return nil, errors.NewError(errors.CodeStringSetFailed)
	}

	return &types.StringSetData{
		Result: result,
	}, nil
}

// Del deletes a key from Redis
func (s *RedisStringServiceImpl) Del(ctx context.Context, req *types.StringDelRequest) (*types.StringDelData, error) {
	// Connect to Redis
	if err := s.dao.Connect(req.RedisRequest); err != nil {
		return nil, errors.NewError(errors.CodeRedisConnectFailed)
	}

	// Delete the key
	deleted, err := s.dao.StringDel(ctx, req.Key)
	if err != nil {
		return nil, errors.NewError(errors.CodeStringDelFailed)
	}

	return &types.StringDelData{
		Deleted: deleted,
	}, nil
}

// Exists checks if a key exists in Redis
func (s *RedisStringServiceImpl) Exists(ctx context.Context, req *types.StringExistsRequest) (*types.StringExistsData, error) {
	// Connect to Redis
	if err := s.dao.Connect(req.RedisRequest); err != nil {
		return nil, errors.NewError(errors.CodeRedisConnectFailed)
	}

	// Check if key exists
	exists, err := s.dao.StringExists(ctx, req.Key)
	if err != nil {
		return nil, errors.NewError(errors.CodeStringGetFailed)
	}

	return &types.StringExistsData{
		Exists: exists,
	}, nil
}

// Incr increments the integer value of a key by 1
func (s *RedisStringServiceImpl) Incr(ctx context.Context, req *types.StringIncrRequest) (*types.StringIncrData, error) {
	// Connect to Redis
	if err := s.dao.Connect(req.RedisRequest); err != nil {
		return nil, errors.NewError(errors.CodeRedisConnectFailed)
	}

	// Increment the value
	value, err := s.dao.StringIncr(ctx, req.Key)
	if err != nil {
		return nil, errors.NewError(errors.CodeStringIncrFailed)
	}

	return &types.StringIncrData{
		Value: value,
	}, nil
}

// Decr decrements the integer value of a key by 1
func (s *RedisStringServiceImpl) Decr(ctx context.Context, req *types.StringDecrRequest) (*types.StringDecrData, error) {
	// Connect to Redis
	if err := s.dao.Connect(req.RedisRequest); err != nil {
		return nil, errors.NewError(errors.CodeRedisConnectFailed)
	}

	// Decrement the value
	value, err := s.dao.StringDecr(ctx, req.Key)
	if err != nil {
		return nil, errors.NewError(errors.CodeStringDecrFailed)
	}

	return &types.StringDecrData{
		Value: value,
	}, nil
}

// Expire sets TTL for a key
func (s *RedisStringServiceImpl) Expire(ctx context.Context, req *types.StringExpireRequest) (*types.StringExpireData, error) {
	// Connect to Redis
	if err := s.dao.Connect(req.RedisRequest); err != nil {
		return nil, errors.NewError(errors.CodeRedisConnectFailed)
	}

	// Set TTL
	ttl := time.Duration(req.TTL) * time.Second
	success, err := s.dao.StringExpire(ctx, req.Key, ttl)
	if err != nil {
		return nil, errors.NewError(errors.CodeStringExpireFailed)
	}

	return &types.StringExpireData{
		Success: success,
	}, nil
}