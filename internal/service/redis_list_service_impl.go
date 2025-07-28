package service

import (
	"context"

	"github.com/ct-zh/go-redis-proxy/internal/dao"
	"github.com/ct-zh/go-redis-proxy/pkg/types"
	"github.com/ct-zh/go-redis-proxy/pkg/errors"
)

// RedisListServiceImpl implements RedisListService interface
type RedisListServiceImpl struct {
	dao dao.RedisDAO
}

// NewRedisListService creates a new RedisListService instance
func NewRedisListService(redisDAO dao.RedisDAO) RedisListService {
	return &RedisListServiceImpl{
		dao: redisDAO,
	}
}

// LPush pushes values to the left of a list
func (s *RedisListServiceImpl) LPush(ctx context.Context, req *types.ListLPushRequest) (*types.ListLPushData, error) {
	// Connect to Redis
	if err := s.dao.Connect(req.RedisRequest); err != nil {
		return nil, errors.NewError(errors.CodeRedisConnectFailed)
	}

	// Push values to the left
	length, err := s.dao.ListLPush(ctx, req.Key, req.Values)
	if err != nil {
		return nil, errors.NewError(errors.CodeListPushFailed)
	}

	return &types.ListLPushData{
		Length: length,
	}, nil
}

// RPush pushes values to the right of a list
func (s *RedisListServiceImpl) RPush(ctx context.Context, req *types.ListRPushRequest) (*types.ListRPushData, error) {
	// Connect to Redis
	if err := s.dao.Connect(req.RedisRequest); err != nil {
		return nil, errors.NewError(errors.CodeRedisConnectFailed)
	}

	// Push values to the right
	length, err := s.dao.ListRPush(ctx, req.Key, req.Values)
	if err != nil {
		return nil, errors.NewError(errors.CodeListPushFailed)
	}

	return &types.ListRPushData{
		Length: length,
	}, nil
}

// LPop pops a value from the left of a list
func (s *RedisListServiceImpl) LPop(ctx context.Context, req *types.ListLPopRequest) (*types.ListLPopData, error) {
	// Connect to Redis
	if err := s.dao.Connect(req.RedisRequest); err != nil {
		return nil, errors.NewError(errors.CodeRedisConnectFailed)
	}

	// Pop value from the left
	value, err := s.dao.ListLPop(ctx, req.Key)
	if err != nil {
		return nil, errors.NewError(errors.CodeListPopFailed)
	}

	return &types.ListLPopData{
		Value: value,
	}, nil
}

// RPop pops a value from the right of a list
func (s *RedisListServiceImpl) RPop(ctx context.Context, req *types.ListRPopRequest) (*types.ListRPopData, error) {
	// Connect to Redis
	if err := s.dao.Connect(req.RedisRequest); err != nil {
		return nil, errors.NewError(errors.CodeRedisConnectFailed)
	}

	// Pop value from the right
	value, err := s.dao.ListRPop(ctx, req.Key)
	if err != nil {
		return nil, errors.NewError(errors.CodeListPopFailed)
	}

	return &types.ListRPopData{
		Value: value,
	}, nil
}

// LRem removes elements from a list
func (s *RedisListServiceImpl) LRem(ctx context.Context, req *types.ListLRemRequest) (*types.ListLRemData, error) {
	// Connect to Redis
	if err := s.dao.Connect(req.RedisRequest); err != nil {
		return nil, errors.NewError(errors.CodeRedisConnectFailed)
	}

	// Remove elements
	removed, err := s.dao.ListLRem(ctx, req.Key, req.Count, req.Value)
	if err != nil {
		return nil, errors.NewError(errors.CodeListRemoveFailed)
	}

	return &types.ListLRemData{
		Removed: removed,
	}, nil
}

// LIndex gets an element from a list by index
func (s *RedisListServiceImpl) LIndex(ctx context.Context, req *types.ListLIndexRequest) (*types.ListLIndexData, error) {
	// Connect to Redis
	if err := s.dao.Connect(req.RedisRequest); err != nil {
		return nil, errors.NewError(errors.CodeRedisConnectFailed)
	}

	// Get element by index
	value, err := s.dao.ListLIndex(ctx, req.Key, req.Index)
	if err != nil {
		return nil, errors.NewError(errors.CodeListIndexOutOfRange)
	}

	return &types.ListLIndexData{
		Value: value,
	}, nil
}

// LRange gets a range of elements from a list
func (s *RedisListServiceImpl) LRange(ctx context.Context, req *types.ListLRangeRequest) (*types.ListLRangeData, error) {
	// Connect to Redis
	if err := s.dao.Connect(req.RedisRequest); err != nil {
		return nil, errors.NewError(errors.CodeRedisConnectFailed)
	}

	// Get range of elements
	values, err := s.dao.ListLRange(ctx, req.Key, req.Start, req.Stop)
	if err != nil {
		return nil, errors.NewError(errors.CodeListIndexOutOfRange)
	}

	return &types.ListLRangeData{
		Values: values,
	}, nil
}

// LLen gets the length of a list
func (s *RedisListServiceImpl) LLen(ctx context.Context, req *types.ListLLenRequest) (*types.ListLLenData, error) {
	// Connect to Redis
	if err := s.dao.Connect(req.RedisRequest); err != nil {
		return nil, errors.NewError(errors.CodeRedisConnectFailed)
	}

	// Get list length
	length, err := s.dao.ListLLen(ctx, req.Key)
	if err != nil {
		return nil, errors.NewError(errors.CodeStringGetFailed)
	}

	return &types.ListLLenData{
		Length: length,
	}, nil
}

// LTrim trims a list to a specified range
func (s *RedisListServiceImpl) LTrim(ctx context.Context, req *types.ListLTrimRequest) (*types.ListLTrimData, error) {
	// Connect to Redis
	if err := s.dao.Connect(req.RedisRequest); err != nil {
		return nil, errors.NewError(errors.CodeRedisConnectFailed)
	}

	// Trim the list
	result, err := s.dao.ListLTrim(ctx, req.Key, req.Start, req.Stop)
	if err != nil {
		return nil, errors.NewError(errors.CodeListTrimFailed)
	}

	return &types.ListLTrimData{
		Result: result,
	}, nil
}