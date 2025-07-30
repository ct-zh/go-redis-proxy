package service

import (
	"context"
	"time"

	"github.com/ct-zh/go-redis-proxy/internal/connection"
	"github.com/ct-zh/go-redis-proxy/internal/dao"
	"github.com/ct-zh/go-redis-proxy/pkg/errors"
	"github.com/ct-zh/go-redis-proxy/pkg/types"
)

// DynamicRedisStringService 支持动态连接的Redis String服务
type DynamicRedisStringService struct {
	connectionService *connection.Service
	dynamicDAO        dao.DynamicRedisDAO
}

// NewDynamicRedisStringService 创建新的动态Redis String服务
func NewDynamicRedisStringService(connectionService *connection.Service, dynamicDAO dao.DynamicRedisDAO) *DynamicRedisStringService {
	return &DynamicRedisStringService{
		connectionService: connectionService,
		dynamicDAO:        dynamicDAO,
	}
}

// Get 获取字符串值
func (s *DynamicRedisStringService) Get(ctx context.Context, req *types.StringGetRequest) (*types.StringGetData, error) {
	// 获取Redis连接
	client, err := s.connectionService.GetConnection(ctx, &req.RedisRequest)
	if err != nil {
		return nil, errors.NewRedisConnectionError("Failed to get connection", err)
	}

	// 设置动态DAO的客户端
	dynamicDAO := s.dynamicDAO.WithClient(client)

	// 调用DAO层
	value, err := dynamicDAO.StringGet(ctx, req.Key)
	if err != nil {
		return nil, errors.NewRedisOperationError("Failed to get value", err)
	}

	return &types.StringGetData{
		Value: value,
	}, nil
}

// Set 设置字符串值
func (s *DynamicRedisStringService) Set(ctx context.Context, req *types.StringSetRequest) (*types.StringSetData, error) {
	// 获取Redis连接
	client, err := s.connectionService.GetConnection(ctx, &req.RedisRequest)
	if err != nil {
		return nil, errors.NewRedisConnectionError("Failed to get connection", err)
	}

	// 设置动态DAO的客户端
	dynamicDAO := s.dynamicDAO.WithClient(client)

	// 调用DAO层
	result, err := dynamicDAO.StringSet(ctx, req.Key, req.Value, time.Duration(req.TTL)*time.Second)
	if err != nil {
		return nil, errors.NewRedisOperationError("Failed to set value", err)
	}

	return &types.StringSetData{
		Result: result,
	}, nil
}

// Del 删除键
func (s *DynamicRedisStringService) Del(ctx context.Context, req *types.StringDelRequest) (*types.StringDelData, error) {
	// 获取Redis连接
	client, err := s.connectionService.GetConnection(ctx, &req.RedisRequest)
	if err != nil {
		return nil, errors.NewRedisConnectionError("Failed to get connection", err)
	}

	// 设置动态DAO的客户端
	dynamicDAO := s.dynamicDAO.WithClient(client)

	// 调用DAO层
	deleted, err := dynamicDAO.StringDel(ctx, req.Key)
	if err != nil {
		return nil, errors.NewRedisOperationError("Failed to delete key", err)
	}

	return &types.StringDelData{
		Deleted: deleted,
	}, nil
}

// Exists 检查键是否存在
func (s *DynamicRedisStringService) Exists(ctx context.Context, req *types.StringExistsRequest) (*types.StringExistsData, error) {
	// 获取Redis连接
	client, err := s.connectionService.GetConnection(ctx, &req.RedisRequest)
	if err != nil {
		return nil, errors.NewRedisConnectionError("Failed to get connection", err)
	}

	// 设置动态DAO的客户端
	dynamicDAO := s.dynamicDAO.WithClient(client)

	// 调用DAO层
	exists, err := dynamicDAO.StringExists(ctx, req.Key)
	if err != nil {
		return nil, errors.NewRedisOperationError("Failed to check existence", err)
	}

	return &types.StringExistsData{
		Exists: exists,
	}, nil
}

// Incr 递增键值
func (s *DynamicRedisStringService) Incr(ctx context.Context, req *types.StringIncrRequest) (*types.StringIncrData, error) {
	// 获取Redis连接
	client, err := s.connectionService.GetConnection(ctx, &req.RedisRequest)
	if err != nil {
		return nil, errors.NewRedisConnectionError("Failed to get connection", err)
	}

	// 设置动态DAO的客户端
	dynamicDAO := s.dynamicDAO.WithClient(client)

	// 调用DAO层
	value, err := dynamicDAO.StringIncr(ctx, req.Key)
	if err != nil {
		return nil, errors.NewRedisOperationError("Failed to increment", err)
	}

	return &types.StringIncrData{
		Value: value,
	}, nil
}

// Decr 递减键值
func (s *DynamicRedisStringService) Decr(ctx context.Context, req *types.StringDecrRequest) (*types.StringDecrData, error) {
	// 获取Redis连接
	client, err := s.connectionService.GetConnection(ctx, &req.RedisRequest)
	if err != nil {
		return nil, errors.NewRedisConnectionError("Failed to get connection", err)
	}

	// 设置动态DAO的客户端
	dynamicDAO := s.dynamicDAO.WithClient(client)

	// 调用DAO层
	value, err := dynamicDAO.StringDecr(ctx, req.Key)
	if err != nil {
		return nil, errors.NewRedisOperationError("Failed to decrement", err)
	}

	return &types.StringDecrData{
		Value: value,
	}, nil
}

// Expire 设置键过期时间
func (s *DynamicRedisStringService) Expire(ctx context.Context, req *types.StringExpireRequest) (*types.StringExpireData, error) {
	// 获取Redis连接
	client, err := s.connectionService.GetConnection(ctx, &req.RedisRequest)
	if err != nil {
		return nil, errors.NewRedisConnectionError("Failed to get connection", err)
	}

	// 设置动态DAO的客户端
	dynamicDAO := s.dynamicDAO.WithClient(client)

	// 调用DAO层
	success, err := dynamicDAO.StringExpire(ctx, req.Key, time.Duration(req.TTL)*time.Second)
	if err != nil {
		return nil, errors.NewRedisOperationError("Failed to set expiration", err)
	}

	return &types.StringExpireData{
		Success: success,
	}, nil
}