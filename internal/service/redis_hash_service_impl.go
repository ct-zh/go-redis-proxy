package service

import (
	"context"

	"github.com/ct-zh/go-redis-proxy/internal/dao"
	"github.com/ct-zh/go-redis-proxy/pkg/errors"
	"github.com/ct-zh/go-redis-proxy/pkg/types"
)

// RedisHashServiceImpl implements the RedisHashService interface
type RedisHashServiceImpl struct {
	redisDAO dao.RedisDAO
}

// NewRedisHashService creates a new RedisHashServiceImpl instance
func NewRedisHashService(redisDAO dao.RedisDAO) RedisHashService {
	return &RedisHashServiceImpl{
		redisDAO: redisDAO,
	}
}

// HSet sets field-value pairs in a hash
func (s *RedisHashServiceImpl) HSet(ctx context.Context, req *types.HashHSetRequest) (*types.HashHSetData, error) {
	// Connect to Redis
	if err := s.redisDAO.Connect(req.RedisRequest); err != nil {
		return nil, errors.NewError(errors.CodeRedisConnectFailed)
	}
	defer s.redisDAO.Close()

	// Call DAO layer
	set, err := s.redisDAO.HashHSet(ctx, req.Key, req.Fields)
	if err != nil {
		return nil, errors.NewError(errors.CodeHashSetFailed)
	}

	return &types.HashHSetData{Set: set}, nil
}

// HGet gets the value of a field in a hash
func (s *RedisHashServiceImpl) HGet(ctx context.Context, req *types.HashHGetRequest) (*types.HashHGetData, error) {
	// Connect to Redis
	if err := s.redisDAO.Connect(req.RedisRequest); err != nil {
		return nil, errors.NewError(errors.CodeRedisConnectFailed)
	}
	defer s.redisDAO.Close()

	// Call DAO layer
	value, err := s.redisDAO.HashHGet(ctx, req.Key, req.Field)
	if err != nil {
		return nil, errors.NewError(errors.CodeHashGetFailed)
	}

	return &types.HashHGetData{Value: value}, nil
}

// HMGet gets values of multiple fields in a hash
func (s *RedisHashServiceImpl) HMGet(ctx context.Context, req *types.HashHMGetRequest) (*types.HashHMGetData, error) {
	// Connect to Redis
	if err := s.redisDAO.Connect(req.RedisRequest); err != nil {
		return nil, errors.NewError(errors.CodeRedisConnectFailed)
	}
	defer s.redisDAO.Close()

	// Call DAO layer
	values, err := s.redisDAO.HashHMGet(ctx, req.Key, req.Fields)
	if err != nil {
		return nil, errors.NewError(errors.CodeHashGetFailed)
	}

	return &types.HashHMGetData{Values: values}, nil
}

// HGetAll gets all field-value pairs in a hash
func (s *RedisHashServiceImpl) HGetAll(ctx context.Context, req *types.HashHGetAllRequest) (*types.HashHGetAllData, error) {
	// Connect to Redis
	if err := s.redisDAO.Connect(req.RedisRequest); err != nil {
		return nil, errors.NewError(errors.CodeRedisConnectFailed)
	}
	defer s.redisDAO.Close()

	// Call DAO layer
	fields, err := s.redisDAO.HashHGetAll(ctx, req.Key)
	if err != nil {
		return nil, errors.NewError(errors.CodeHashGetFailed)
	}

	return &types.HashHGetAllData{Fields: fields}, nil
}

// HDel deletes fields from a hash
func (s *RedisHashServiceImpl) HDel(ctx context.Context, req *types.HashHDelRequest) (*types.HashHDelData, error) {
	// Connect to Redis
	if err := s.redisDAO.Connect(req.RedisRequest); err != nil {
		return nil, errors.NewError(errors.CodeRedisConnectFailed)
	}
	defer s.redisDAO.Close()

	// Call DAO layer
	deleted, err := s.redisDAO.HashHDel(ctx, req.Key, req.Fields)
	if err != nil {
		return nil, errors.NewError(errors.CodeHashDeleteFailed)
	}

	return &types.HashHDelData{Deleted: deleted}, nil
}

// HExists checks if a field exists in a hash
func (s *RedisHashServiceImpl) HExists(ctx context.Context, req *types.HashHExistsRequest) (*types.HashHExistsData, error) {
	// Connect to Redis
	if err := s.redisDAO.Connect(req.RedisRequest); err != nil {
		return nil, errors.NewError(errors.CodeRedisConnectFailed)
	}
	defer s.redisDAO.Close()

	// Call DAO layer
	exists, err := s.redisDAO.HashHExists(ctx, req.Key, req.Field)
	if err != nil {
		return nil, errors.NewError(errors.CodeHashGetFailed)
	}

	return &types.HashHExistsData{Exists: exists}, nil
}

// HLen gets the number of fields in a hash
func (s *RedisHashServiceImpl) HLen(ctx context.Context, req *types.HashHLenRequest) (*types.HashHLenData, error) {
	// Connect to Redis
	if err := s.redisDAO.Connect(req.RedisRequest); err != nil {
		return nil, errors.NewError(errors.CodeRedisConnectFailed)
	}
	defer s.redisDAO.Close()

	// Call DAO layer
	length, err := s.redisDAO.HashHLen(ctx, req.Key)
	if err != nil {
		return nil, errors.NewError(errors.CodeHashGetFailed)
	}

	return &types.HashHLenData{Length: length}, nil
}

// HKeys gets all field names in a hash
func (s *RedisHashServiceImpl) HKeys(ctx context.Context, req *types.HashHKeysRequest) (*types.HashHKeysData, error) {
	// Connect to Redis
	if err := s.redisDAO.Connect(req.RedisRequest); err != nil {
		return nil, errors.NewError(errors.CodeRedisConnectFailed)
	}
	defer s.redisDAO.Close()

	// Call DAO layer
	keys, err := s.redisDAO.HashHKeys(ctx, req.Key)
	if err != nil {
		return nil, errors.NewError(errors.CodeHashGetFailed)
	}

	return &types.HashHKeysData{Keys: keys}, nil
}

// HVals gets all field values in a hash
func (s *RedisHashServiceImpl) HVals(ctx context.Context, req *types.HashHValsRequest) (*types.HashHValsData, error) {
	// Connect to Redis
	if err := s.redisDAO.Connect(req.RedisRequest); err != nil {
		return nil, errors.NewError(errors.CodeRedisConnectFailed)
	}
	defer s.redisDAO.Close()

	// Call DAO layer
	values, err := s.redisDAO.HashHVals(ctx, req.Key)
	if err != nil {
		return nil, errors.NewError(errors.CodeHashGetFailed)
	}

	return &types.HashHValsData{Values: values}, nil
}

// HIncrBy increments the value of a field in a hash by an integer
func (s *RedisHashServiceImpl) HIncrBy(ctx context.Context, req *types.HashHIncrByRequest) (*types.HashHIncrByData, error) {
	// Connect to Redis
	if err := s.redisDAO.Connect(req.RedisRequest); err != nil {
		return nil, errors.NewError(errors.CodeRedisConnectFailed)
	}
	defer s.redisDAO.Close()

	// Call DAO layer
	value, err := s.redisDAO.HashHIncrBy(ctx, req.Key, req.Field, req.Increment)
	if err != nil {
		return nil, errors.NewError(errors.CodeHashIncrementFailed)
	}

	return &types.HashHIncrByData{Value: value}, nil
}