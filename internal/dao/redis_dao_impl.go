package dao

import (
	"context"
	"time"

	redis "github.com/go-redis/redis/v8"

	"github.com/ct-zh/go-redis-proxy/pkg/types"
)

// RedisDAOImpl implements the RedisDAO interface using go-redis client
type RedisDAOImpl struct {
	client *redis.Client
}

// NewRedisDAO creates a new instance of RedisDAOImpl
func NewRedisDAO() *RedisDAOImpl {
	return &RedisDAOImpl{}
}

// Connect establishes a connection to Redis
func (r *RedisDAOImpl) Connect(config types.RedisRequest) error {
	r.client = redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
	})
	
	// Test the connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	return r.client.Ping(ctx).Err()
}

// Close closes the Redis connection
func (r *RedisDAOImpl) Close() error {
	if r.client != nil {
		return r.client.Close()
	}
	return nil
}

// Ping tests the Redis connection
func (r *RedisDAOImpl) Ping(ctx context.Context) error {
	return r.client.Ping(ctx).Err()
}

// String operations

// StringGet retrieves a string value from Redis
func (r *RedisDAOImpl) StringGet(ctx context.Context, key string) (interface{}, error) {
	result := r.client.Get(ctx, key)
	if result.Err() == redis.Nil {
		return nil, nil // Key does not exist
	}
	if result.Err() != nil {
		return nil, result.Err()
	}
	return result.Val(), nil
}

// StringSet sets a string value in Redis with optional TTL
func (r *RedisDAOImpl) StringSet(ctx context.Context, key string, value interface{}, ttl time.Duration) (string, error) {
	result := r.client.Set(ctx, key, value, ttl)
	return result.Val(), result.Err()
}

// StringDel deletes a key from Redis
func (r *RedisDAOImpl) StringDel(ctx context.Context, key string) (int64, error) {
	result := r.client.Del(ctx, key)
	return result.Val(), result.Err()
}

// StringExists checks if a key exists in Redis
func (r *RedisDAOImpl) StringExists(ctx context.Context, key string) (bool, error) {
	result := r.client.Exists(ctx, key)
	if result.Err() != nil {
		return false, result.Err()
	}
	return result.Val() > 0, nil
}

// StringIncr increments the integer value of a key by 1
func (r *RedisDAOImpl) StringIncr(ctx context.Context, key string) (int64, error) {
	result := r.client.Incr(ctx, key)
	return result.Val(), result.Err()
}

// StringDecr decrements the integer value of a key by 1
func (r *RedisDAOImpl) StringDecr(ctx context.Context, key string) (int64, error) {
	result := r.client.Decr(ctx, key)
	return result.Val(), result.Err()
}

// StringExpire sets TTL for a key
func (r *RedisDAOImpl) StringExpire(ctx context.Context, key string, ttl time.Duration) (bool, error) {
	result := r.client.Expire(ctx, key, ttl)
	return result.Val(), result.Err()
}

// List operations

// ListLPush pushes values to the left of a list
func (r *RedisDAOImpl) ListLPush(ctx context.Context, key string, values []string) (int64, error) {
	if len(values) == 0 {
		return 0, nil
	}
	
	// Convert []string to []interface{}
	interfaces := make([]interface{}, len(values))
	for i, v := range values {
		interfaces[i] = v
	}
	
	result := r.client.LPush(ctx, key, interfaces...)
	return result.Val(), result.Err()
}

// ListRPush pushes values to the right of a list
func (r *RedisDAOImpl) ListRPush(ctx context.Context, key string, values []string) (int64, error) {
	if len(values) == 0 {
		return 0, nil
	}
	
	// Convert []string to []interface{}
	interfaces := make([]interface{}, len(values))
	for i, v := range values {
		interfaces[i] = v
	}
	
	result := r.client.RPush(ctx, key, interfaces...)
	return result.Val(), result.Err()
}

// ListLPop pops a value from the left of a list
func (r *RedisDAOImpl) ListLPop(ctx context.Context, key string) (interface{}, error) {
	result := r.client.LPop(ctx, key)
	if result.Err() == redis.Nil {
		return nil, nil // List is empty or key does not exist
	}
	if result.Err() != nil {
		return nil, result.Err()
	}
	return result.Val(), nil
}

// ListRPop pops a value from the right of a list
func (r *RedisDAOImpl) ListRPop(ctx context.Context, key string) (interface{}, error) {
	result := r.client.RPop(ctx, key)
	if result.Err() == redis.Nil {
		return nil, nil // List is empty or key does not exist
	}
	if result.Err() != nil {
		return nil, result.Err()
	}
	return result.Val(), nil
}

// ListLRem removes elements from a list
func (r *RedisDAOImpl) ListLRem(ctx context.Context, key string, count int64, value string) (int64, error) {
	result := r.client.LRem(ctx, key, count, value)
	return result.Val(), result.Err()
}

// ListLIndex gets an element from a list by index
func (r *RedisDAOImpl) ListLIndex(ctx context.Context, key string, index int64) (interface{}, error) {
	result := r.client.LIndex(ctx, key, index)
	if result.Err() == redis.Nil {
		return nil, nil // Index out of range or key does not exist
	}
	if result.Err() != nil {
		return nil, result.Err()
	}
	return result.Val(), nil
}

// ListLRange gets a range of elements from a list
func (r *RedisDAOImpl) ListLRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	result := r.client.LRange(ctx, key, start, stop)
	return result.Val(), result.Err()
}

// ListLLen gets the length of a list
func (r *RedisDAOImpl) ListLLen(ctx context.Context, key string) (int64, error) {
	result := r.client.LLen(ctx, key)
	return result.Val(), result.Err()
}

// ListLTrim trims a list to a specified range
func (r *RedisDAOImpl) ListLTrim(ctx context.Context, key string, start, stop int64) (string, error) {
	result := r.client.LTrim(ctx, key, start, stop)
	return result.Val(), result.Err()
}