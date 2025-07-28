package dao

import (
	"context"
	"fmt"
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

// Set operations

// SetSAdd adds members to a set
func (r *RedisDAOImpl) SetSAdd(ctx context.Context, key string, members []string) (int64, error) {
	// convert []string to []interface{}
	interfaces := make([]interface{}, len(members))
	for i, v := range members {
		interfaces[i] = v
	}
	result := r.client.SAdd(ctx, key, interfaces...)
	return result.Val(), result.Err()
}

// SetSRem removes members from a set
func (r *RedisDAOImpl) SetSRem(ctx context.Context, key string, members []string) (int64, error) {
	// convert []string to []interface{}
	interfaces := make([]interface{}, len(members))
	for i, v := range members {
		interfaces[i] = v
	}
	result := r.client.SRem(ctx, key, interfaces...)
	return result.Val(), result.Err()
}

// SetSIsMember checks if a member exists in a set
func (r *RedisDAOImpl) SetSIsMember(ctx context.Context, key string, member string) (bool, error) {
	result := r.client.SIsMember(ctx, key, member)
	return result.Val(), result.Err()
}

// SetSMembers returns all members of a set
func (r *RedisDAOImpl) SetSMembers(ctx context.Context, key string) ([]string, error) {
	result := r.client.SMembers(ctx, key)
	return result.Val(), result.Err()
}

// SetSCard returns the number of members in a set
func (r *RedisDAOImpl) SetSCard(ctx context.Context, key string) (int64, error) {
	result := r.client.SCard(ctx, key)
	return result.Val(), result.Err()
}

// ZSet operations

// ZSetZAdd adds members with scores to a sorted set
func (r *RedisDAOImpl) ZSetZAdd(ctx context.Context, key string, members map[string]float64) (int64, error) {
	// convert map to []*redis.Z
	zs := make([]*redis.Z, 0, len(members))
	for member, score := range members {
		zs = append(zs, &redis.Z{Score: score, Member: member})
	}
	result := r.client.ZAdd(ctx, key, zs...)
	return result.Val(), result.Err()
}

// ZSetZIncrBy increments the score of a member in a sorted set
func (r *RedisDAOImpl) ZSetZIncrBy(ctx context.Context, key string, increment float64, member string) (float64, error) {
	result := r.client.ZIncrBy(ctx, key, increment, member)
	return result.Val(), result.Err()
}

// ZSetZScore gets the score of a member in a sorted set
func (r *RedisDAOImpl) ZSetZScore(ctx context.Context, key string, member string) (interface{}, error) {
	result := r.client.ZScore(ctx, key, member)
	if result.Err() == redis.Nil {
		return nil, nil
	}
	if result.Err() != nil {
		return nil, result.Err()
	}
	return result.Val(), nil
}

// ZSetZCard gets the number of members in a sorted set
func (r *RedisDAOImpl) ZSetZCard(ctx context.Context, key string) (int64, error) {
	result := r.client.ZCard(ctx, key)
	return result.Val(), result.Err()
}

// ZSetZCount counts members in a sorted set within a score range
func (r *RedisDAOImpl) ZSetZCount(ctx context.Context, key string, min, max float64) (int64, error) {
	result := r.client.ZCount(ctx, key, fmt.Sprintf("%f", min), fmt.Sprintf("%f", max))
	return result.Val(), result.Err()
}

// ZSetZRank gets the rank of a member in a sorted set (ascending order)
func (r *RedisDAOImpl) ZSetZRank(ctx context.Context, key string, member string) (interface{}, error) {
	result := r.client.ZRank(ctx, key, member)
	if result.Err() == redis.Nil {
		return nil, nil
	}
	if result.Err() != nil {
		return nil, result.Err()
	}
	return result.Val(), nil
}

// ZSetZRevRank gets the rank of a member in a sorted set (descending order)
func (r *RedisDAOImpl) ZSetZRevRank(ctx context.Context, key string, member string) (interface{}, error) {
	result := r.client.ZRevRank(ctx, key, member)
	if result.Err() == redis.Nil {
		return nil, nil
	}
	if result.Err() != nil {
		return nil, result.Err()
	}
	return result.Val(), nil
}

// ZSetZRange gets members from a sorted set by rank range (ascending order)
func (r *RedisDAOImpl) ZSetZRange(ctx context.Context, key string, start, stop int64, withScores bool) ([]interface{}, error) {
	var result []interface{}
	var err error
	
	if withScores {
		zResult := r.client.ZRangeWithScores(ctx, key, start, stop)
		if zResult.Err() != nil {
			return nil, zResult.Err()
		}
		for _, z := range zResult.Val() {
			result = append(result, z.Member, z.Score)
		}
	} else {
		strResult := r.client.ZRange(ctx, key, start, stop)
		if strResult.Err() != nil {
			return nil, strResult.Err()
		}
		for _, member := range strResult.Val() {
			result = append(result, member)
		}
	}
	
	return result, err
}

// ZSetZRevRange gets members from a sorted set by rank range (descending order)
func (r *RedisDAOImpl) ZSetZRevRange(ctx context.Context, key string, start, stop int64, withScores bool) ([]interface{}, error) {
	var result []interface{}
	var err error
	
	if withScores {
		zResult := r.client.ZRevRangeWithScores(ctx, key, start, stop)
		if zResult.Err() != nil {
			return nil, zResult.Err()
		}
		for _, z := range zResult.Val() {
			result = append(result, z.Member, z.Score)
		}
	} else {
		strResult := r.client.ZRevRange(ctx, key, start, stop)
		if strResult.Err() != nil {
			return nil, strResult.Err()
		}
		for _, member := range strResult.Val() {
			result = append(result, member)
		}
	}
	
	return result, err
}

// ZSetZRangeByScore gets members from a sorted set by score range (ascending order)
func (r *RedisDAOImpl) ZSetZRangeByScore(ctx context.Context, key string, min, max string, withScores bool, offset, count int64) ([]interface{}, error) {
	var result []interface{}
	var err error
	
	opt := &redis.ZRangeBy{
		Min: min,
		Max: max,
	}
	
	if count > 0 {
		opt.Offset = offset
		opt.Count = count
	}
	
	if withScores {
		zResult := r.client.ZRangeByScoreWithScores(ctx, key, opt)
		if zResult.Err() != nil {
			return nil, zResult.Err()
		}
		for _, z := range zResult.Val() {
			result = append(result, z.Member, z.Score)
		}
	} else {
		strResult := r.client.ZRangeByScore(ctx, key, opt)
		if strResult.Err() != nil {
			return nil, strResult.Err()
		}
		for _, member := range strResult.Val() {
			result = append(result, member)
		}
	}
	
	return result, err
}

// ZSetZRevRangeByScore gets members from a sorted set by score range (descending order)
func (r *RedisDAOImpl) ZSetZRevRangeByScore(ctx context.Context, key string, max, min string, withScores bool, offset, count int64) ([]interface{}, error) {
	var result []interface{}
	var err error
	
	opt := &redis.ZRangeBy{
		Min: min,
		Max: max,
	}
	
	if count > 0 {
		opt.Offset = offset
		opt.Count = count
	}
	
	if withScores {
		zResult := r.client.ZRevRangeByScoreWithScores(ctx, key, opt)
		if zResult.Err() != nil {
			return nil, zResult.Err()
		}
		for _, z := range zResult.Val() {
			result = append(result, z.Member, z.Score)
		}
	} else {
		strResult := r.client.ZRevRangeByScore(ctx, key, opt)
		if strResult.Err() != nil {
			return nil, strResult.Err()
		}
		for _, member := range strResult.Val() {
			result = append(result, member)
		}
	}
	
	return result, err
}

// ZSetZRem removes members from a sorted set
func (r *RedisDAOImpl) ZSetZRem(ctx context.Context, key string, members []string) (int64, error) {
	// convert []string to []interface{}
	interfaces := make([]interface{}, len(members))
	for i, v := range members {
		interfaces[i] = v
	}
	result := r.client.ZRem(ctx, key, interfaces...)
	return result.Val(), result.Err()
}

// ZSetZRemRangeByRank removes members from a sorted set by rank range
func (r *RedisDAOImpl) ZSetZRemRangeByRank(ctx context.Context, key string, start, stop int64) (int64, error) {
	result := r.client.ZRemRangeByRank(ctx, key, start, stop)
	return result.Val(), result.Err()
}

// ZSetZRemRangeByScore removes members from a sorted set by score range
func (r *RedisDAOImpl) ZSetZRemRangeByScore(ctx context.Context, key string, min, max string) (int64, error) {
	result := r.client.ZRemRangeByScore(ctx, key, min, max)
	return result.Val(), result.Err()
}