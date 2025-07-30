package dao

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/ct-zh/go-redis-proxy/internal/connection"
	"github.com/ct-zh/go-redis-proxy/pkg/types"
)

// RedisDAOAdapter 适配器，将新的连接管理系统适配到现有的RedisDAO接口
type RedisDAOAdapter struct {
	connectionService connection.Service
	dynamicDAO        DynamicRedisDAO
}

// NewRedisDAOAdapter 创建新的Redis DAO适配器
func NewRedisDAOAdapter(connectionService connection.Service, dynamicDAO DynamicRedisDAO) *RedisDAOAdapter {
	return &RedisDAOAdapter{
		connectionService: connectionService,
		dynamicDAO:        dynamicDAO,
	}
}

// Connect 连接到Redis（适配现有接口）
func (r *RedisDAOAdapter) Connect(ctx context.Context, req *types.RedisRequest) error {
	// 新的连接管理系统不需要显式连接，连接在使用时自动获取
	// 这里只是验证连接配置是否有效
	_, err := r.connectionService.GetConnection(ctx, req)
	return err
}

// Close 关闭连接（适配现有接口）
func (r *RedisDAOAdapter) Close() error {
	// 新的连接管理系统自动管理连接生命周期
	// 这里可以选择性地关闭连接服务
	return r.connectionService.Close()
}

// Ping 测试连接
func (r *RedisDAOAdapter) Ping(ctx context.Context, req *types.RedisRequest) error {
	client, err := r.connectionService.GetConnection(ctx, req)
	if err != nil {
		return err
	}
	
	return client.Ping(ctx).Err()
}

// 以下方法通过动态DAO实现，需要在调用时传入RedisRequest来获取连接

// StringGet 获取字符串值
func (r *RedisDAOAdapter) StringGet(ctx context.Context, req *types.RedisRequest, key string) (interface{}, error) {
	client, err := r.connectionService.GetConnection(ctx, req)
	if err != nil {
		return nil, err
	}
	
	dynamicDAO := r.dynamicDAO.WithClient(client)
	return dynamicDAO.StringGet(ctx, key)
}

// StringSet 设置字符串值
func (r *RedisDAOAdapter) StringSet(ctx context.Context, req *types.RedisRequest, key, value string, expiration int) (string, error) {
	client, err := r.connectionService.GetConnection(ctx, req)
	if err != nil {
		return "", err
	}
	
	dynamicDAO := r.dynamicDAO.WithClient(client)
	return dynamicDAO.StringSet(ctx, key, value, 0) // expiration处理可以后续优化
}

// StringDel 删除键
func (r *RedisDAOAdapter) StringDel(ctx context.Context, req *types.RedisRequest, key string) (int64, error) {
	client, err := r.connectionService.GetConnection(ctx, req)
	if err != nil {
		return 0, err
	}
	
	dynamicDAO := r.dynamicDAO.WithClient(client)
	return dynamicDAO.StringDel(ctx, key)
}

// StringExists 检查键是否存在
func (r *RedisDAOAdapter) StringExists(ctx context.Context, req *types.RedisRequest, key string) (bool, error) {
	client, err := r.connectionService.GetConnection(ctx, req)
	if err != nil {
		return false, err
	}
	
	dynamicDAO := r.dynamicDAO.WithClient(client)
	return dynamicDAO.StringExists(ctx, key)
}

// StringIncr 递增键值
func (r *RedisDAOAdapter) StringIncr(ctx context.Context, req *types.RedisRequest, key string) (int64, error) {
	client, err := r.connectionService.GetConnection(ctx, req)
	if err != nil {
		return 0, err
	}
	
	dynamicDAO := r.dynamicDAO.WithClient(client)
	return dynamicDAO.StringIncr(ctx, key)
}

// StringDecr 递减键值
func (r *RedisDAOAdapter) StringDecr(ctx context.Context, req *types.RedisRequest, key string) (int64, error) {
	client, err := r.connectionService.GetConnection(ctx, req)
	if err != nil {
		return 0, err
	}
	
	dynamicDAO := r.dynamicDAO.WithClient(client)
	return dynamicDAO.StringDecr(ctx, key)
}

// 注意：以下方法需要根据现有RedisDAO接口的具体定义来实现
// 这里只是示例，实际实现需要查看完整的接口定义

// GetClient 获取Redis客户端（如果现有接口需要）
func (r *RedisDAOAdapter) GetClient(ctx context.Context, req *types.RedisRequest) (*redis.Client, error) {
	return r.connectionService.GetConnection(ctx, req)
}