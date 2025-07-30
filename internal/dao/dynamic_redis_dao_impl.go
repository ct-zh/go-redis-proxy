package dao

import (
	"context"
	"fmt"
	"time"

	redis "github.com/go-redis/redis/v8"
)

// DynamicRedisDAOImpl 支持动态连接配置的Redis DAO实现
type DynamicRedisDAOImpl struct {
	client *redis.Client
}

// NewDynamicRedisDAO 创建新的动态Redis DAO实例
func NewDynamicRedisDAO() *DynamicRedisDAOImpl {
	return &DynamicRedisDAOImpl{}
}

// WithClient 设置Redis客户端，返回新的DAO实例
func (r *DynamicRedisDAOImpl) WithClient(client *redis.Client) DynamicRedisDAO {
	return &DynamicRedisDAOImpl{
		client: client,
	}
}

// String operations

// StringGet retrieves a string value from Redis
func (r *DynamicRedisDAOImpl) StringGet(ctx context.Context, key string) (interface{}, error) {
	result := r.client.Get(ctx, key)
	if result.Err() == redis.Nil {
		return nil, nil // Key does not exist
	}
	return result.Val(), result.Err()
}

// StringSet sets a string value in Redis
func (r *DynamicRedisDAOImpl) StringSet(ctx context.Context, key string, value interface{}, ttl time.Duration) (string, error) {
	result := r.client.Set(ctx, key, value, ttl)
	return result.Val(), result.Err()
}

// StringDel deletes a key from Redis
func (r *DynamicRedisDAOImpl) StringDel(ctx context.Context, key string) (int64, error) {
	result := r.client.Del(ctx, key)
	return result.Val(), result.Err()
}

// StringExists checks if a key exists in Redis
func (r *DynamicRedisDAOImpl) StringExists(ctx context.Context, key string) (bool, error) {
	result := r.client.Exists(ctx, key)
	return result.Val() > 0, result.Err()
}

// StringIncr increments the value of a key
func (r *DynamicRedisDAOImpl) StringIncr(ctx context.Context, key string) (int64, error) {
	result := r.client.Incr(ctx, key)
	return result.Val(), result.Err()
}

// StringDecr decrements the value of a key
func (r *DynamicRedisDAOImpl) StringDecr(ctx context.Context, key string) (int64, error) {
	result := r.client.Decr(ctx, key)
	return result.Val(), result.Err()
}

// StringExpire sets the expiration time for a key
func (r *DynamicRedisDAOImpl) StringExpire(ctx context.Context, key string, ttl time.Duration) (bool, error) {
	result := r.client.Expire(ctx, key, ttl)
	return result.Val(), result.Err()
}

// List operations

// ListLPush pushes values to the left of a list
func (r *DynamicRedisDAOImpl) ListLPush(ctx context.Context, key string, values []string) (int64, error) {
	interfaces := make([]interface{}, len(values))
	for i, v := range values {
		interfaces[i] = v
	}
	result := r.client.LPush(ctx, key, interfaces...)
	return result.Val(), result.Err()
}

// ListRPush pushes values to the right of a list
func (r *DynamicRedisDAOImpl) ListRPush(ctx context.Context, key string, values []string) (int64, error) {
	interfaces := make([]interface{}, len(values))
	for i, v := range values {
		interfaces[i] = v
	}
	result := r.client.RPush(ctx, key, interfaces...)
	return result.Val(), result.Err()
}

// ListLPop pops a value from the left of a list
func (r *DynamicRedisDAOImpl) ListLPop(ctx context.Context, key string) (interface{}, error) {
	result := r.client.LPop(ctx, key)
	if result.Err() == redis.Nil {
		return nil, nil
	}
	return result.Val(), result.Err()
}

// ListRPop pops a value from the right of a list
func (r *DynamicRedisDAOImpl) ListRPop(ctx context.Context, key string) (interface{}, error) {
	result := r.client.RPop(ctx, key)
	if result.Err() == redis.Nil {
		return nil, nil
	}
	return result.Val(), result.Err()
}

// ListLRem removes elements from a list
func (r *DynamicRedisDAOImpl) ListLRem(ctx context.Context, key string, count int64, value string) (int64, error) {
	result := r.client.LRem(ctx, key, count, value)
	return result.Val(), result.Err()
}

// ListLIndex gets an element from a list by index
func (r *DynamicRedisDAOImpl) ListLIndex(ctx context.Context, key string, index int64) (interface{}, error) {
	result := r.client.LIndex(ctx, key, index)
	if result.Err() == redis.Nil {
		return nil, nil
	}
	return result.Val(), result.Err()
}

// ListLRange gets a range of elements from a list
func (r *DynamicRedisDAOImpl) ListLRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	result := r.client.LRange(ctx, key, start, stop)
	return result.Val(), result.Err()
}

// ListLLen gets the length of a list
func (r *DynamicRedisDAOImpl) ListLLen(ctx context.Context, key string) (int64, error) {
	result := r.client.LLen(ctx, key)
	return result.Val(), result.Err()
}

// ListLTrim trims a list to the specified range
func (r *DynamicRedisDAOImpl) ListLTrim(ctx context.Context, key string, start, stop int64) (string, error) {
	result := r.client.LTrim(ctx, key, start, stop)
	return result.Val(), result.Err()
}

// Set operations

// SetSAdd adds members to a set
func (r *DynamicRedisDAOImpl) SetSAdd(ctx context.Context, key string, members []string) (int64, error) {
	interfaces := make([]interface{}, len(members))
	for i, v := range members {
		interfaces[i] = v
	}
	result := r.client.SAdd(ctx, key, interfaces...)
	return result.Val(), result.Err()
}

// SetSRem removes members from a set
func (r *DynamicRedisDAOImpl) SetSRem(ctx context.Context, key string, members []string) (int64, error) {
	interfaces := make([]interface{}, len(members))
	for i, v := range members {
		interfaces[i] = v
	}
	result := r.client.SRem(ctx, key, interfaces...)
	return result.Val(), result.Err()
}

// SetSIsMember checks if a member exists in a set
func (r *DynamicRedisDAOImpl) SetSIsMember(ctx context.Context, key string, member string) (bool, error) {
	result := r.client.SIsMember(ctx, key, member)
	return result.Val(), result.Err()
}

// SetSMembers gets all members of a set
func (r *DynamicRedisDAOImpl) SetSMembers(ctx context.Context, key string) ([]string, error) {
	result := r.client.SMembers(ctx, key)
	return result.Val(), result.Err()
}

// SetSCard gets the cardinality of a set
func (r *DynamicRedisDAOImpl) SetSCard(ctx context.Context, key string) (int64, error) {
	result := r.client.SCard(ctx, key)
	return result.Val(), result.Err()
}

// ZSet operations

// ZSetZAdd adds members with scores to a sorted set
func (r *DynamicRedisDAOImpl) ZSetZAdd(ctx context.Context, key string, members map[string]float64) (int64, error) {
	var z []*redis.Z
	for member, score := range members {
		z = append(z, &redis.Z{Score: score, Member: member})
	}
	result := r.client.ZAdd(ctx, key, z...)
	return result.Val(), result.Err()
}

// ZSetZIncrBy increments the score of a member in a sorted set
func (r *DynamicRedisDAOImpl) ZSetZIncrBy(ctx context.Context, key string, increment float64, member string) (float64, error) {
	result := r.client.ZIncrBy(ctx, key, increment, member)
	return result.Val(), result.Err()
}

// ZSetZScore gets the score of a member in a sorted set
func (r *DynamicRedisDAOImpl) ZSetZScore(ctx context.Context, key string, member string) (interface{}, error) {
	result := r.client.ZScore(ctx, key, member)
	if result.Err() == redis.Nil {
		return nil, nil
	}
	return result.Val(), result.Err()
}

// ZSetZCard gets the cardinality of a sorted set
func (r *DynamicRedisDAOImpl) ZSetZCard(ctx context.Context, key string) (int64, error) {
	result := r.client.ZCard(ctx, key)
	return result.Val(), result.Err()
}

// ZSetZCount counts members in a sorted set within a score range
func (r *DynamicRedisDAOImpl) ZSetZCount(ctx context.Context, key string, min, max float64) (int64, error) {
	result := r.client.ZCount(ctx, key, fmt.Sprintf("%f", min), fmt.Sprintf("%f", max))
	return result.Val(), result.Err()
}

// ZSetZRank gets the rank of a member in a sorted set
func (r *DynamicRedisDAOImpl) ZSetZRank(ctx context.Context, key string, member string) (interface{}, error) {
	result := r.client.ZRank(ctx, key, member)
	if result.Err() == redis.Nil {
		return nil, nil
	}
	return result.Val(), result.Err()
}

// ZSetZRevRank gets the reverse rank of a member in a sorted set
func (r *DynamicRedisDAOImpl) ZSetZRevRank(ctx context.Context, key string, member string) (interface{}, error) {
	result := r.client.ZRevRank(ctx, key, member)
	if result.Err() == redis.Nil {
		return nil, nil
	}
	return result.Val(), result.Err()
}

// ZSetZRange gets members in a sorted set by rank range
func (r *DynamicRedisDAOImpl) ZSetZRange(ctx context.Context, key string, start, stop int64, withScores bool) ([]interface{}, error) {
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
		for _, s := range strResult.Val() {
			result = append(result, s)
		}
	}
	
	return result, err
}

// ZSetZRevRange gets members in a sorted set by reverse rank range
func (r *DynamicRedisDAOImpl) ZSetZRevRange(ctx context.Context, key string, start, stop int64, withScores bool) ([]interface{}, error) {
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
		for _, s := range strResult.Val() {
			result = append(result, s)
		}
	}
	
	return result, err
}

// ZSetZRangeByScore gets members in a sorted set by score range
func (r *DynamicRedisDAOImpl) ZSetZRangeByScore(ctx context.Context, key string, min, max string, withScores bool, offset, count int64) ([]interface{}, error) {
	opt := &redis.ZRangeBy{
		Min:    min,
		Max:    max,
		Offset: offset,
		Count:  count,
	}
	
	var result []interface{}
	var err error
	
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
		for _, s := range strResult.Val() {
			result = append(result, s)
		}
	}
	
	return result, err
}

// ZSetZRevRangeByScore gets members in a sorted set by reverse score range
func (r *DynamicRedisDAOImpl) ZSetZRevRangeByScore(ctx context.Context, key string, max, min string, withScores bool, offset, count int64) ([]interface{}, error) {
	opt := &redis.ZRangeBy{
		Min:    min,
		Max:    max,
		Offset: offset,
		Count:  count,
	}
	
	var result []interface{}
	var err error
	
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
		for _, s := range strResult.Val() {
			result = append(result, s)
		}
	}
	
	return result, err
}

// ZSetZRem removes members from a sorted set
func (r *DynamicRedisDAOImpl) ZSetZRem(ctx context.Context, key string, members []string) (int64, error) {
	interfaces := make([]interface{}, len(members))
	for i, v := range members {
		interfaces[i] = v
	}
	result := r.client.ZRem(ctx, key, interfaces...)
	return result.Val(), result.Err()
}

// ZSetZRemRangeByRank removes members from a sorted set by rank range
func (r *DynamicRedisDAOImpl) ZSetZRemRangeByRank(ctx context.Context, key string, start, stop int64) (int64, error) {
	result := r.client.ZRemRangeByRank(ctx, key, start, stop)
	return result.Val(), result.Err()
}

// ZSetZRemRangeByScore removes members from a sorted set by score range
func (r *DynamicRedisDAOImpl) ZSetZRemRangeByScore(ctx context.Context, key string, min, max string) (int64, error) {
	result := r.client.ZRemRangeByScore(ctx, key, min, max)
	return result.Val(), result.Err()
}

// Hash operations

// HashHSet sets field-value pairs in a hash
func (r *DynamicRedisDAOImpl) HashHSet(ctx context.Context, key string, fields map[string]string) (int64, error) {
	interfaces := make([]interface{}, 0, len(fields)*2)
	for field, value := range fields {
		interfaces = append(interfaces, field, value)
	}
	result := r.client.HSet(ctx, key, interfaces...)
	return result.Val(), result.Err()
}

// HashHGet gets a field value from a hash
func (r *DynamicRedisDAOImpl) HashHGet(ctx context.Context, key string, field string) (interface{}, error) {
	result := r.client.HGet(ctx, key, field)
	if result.Err() == redis.Nil {
		return nil, nil
	}
	return result.Val(), result.Err()
}

// HashHMGet gets multiple field values from a hash
func (r *DynamicRedisDAOImpl) HashHMGet(ctx context.Context, key string, fields []string) ([]interface{}, error) {
	result := r.client.HMGet(ctx, key, fields...)
	return result.Val(), result.Err()
}

// HashHGetAll gets all field-value pairs from a hash
func (r *DynamicRedisDAOImpl) HashHGetAll(ctx context.Context, key string) (map[string]string, error) {
	result := r.client.HGetAll(ctx, key)
	return result.Val(), result.Err()
}

// HashHDel deletes fields from a hash
func (r *DynamicRedisDAOImpl) HashHDel(ctx context.Context, key string, fields []string) (int64, error) {
	result := r.client.HDel(ctx, key, fields...)
	return result.Val(), result.Err()
}

// HashHExists checks if a field exists in a hash
func (r *DynamicRedisDAOImpl) HashHExists(ctx context.Context, key string, field string) (bool, error) {
	result := r.client.HExists(ctx, key, field)
	return result.Val(), result.Err()
}

// HashHLen gets the number of fields in a hash
func (r *DynamicRedisDAOImpl) HashHLen(ctx context.Context, key string) (int64, error) {
	result := r.client.HLen(ctx, key)
	return result.Val(), result.Err()
}

// HashHKeys gets all field names in a hash
func (r *DynamicRedisDAOImpl) HashHKeys(ctx context.Context, key string) ([]string, error) {
	result := r.client.HKeys(ctx, key)
	return result.Val(), result.Err()
}

// HashHVals gets all field values in a hash
func (r *DynamicRedisDAOImpl) HashHVals(ctx context.Context, key string) ([]string, error) {
	result := r.client.HVals(ctx, key)
	return result.Val(), result.Err()
}

// HashHIncrBy increments a field value in a hash
func (r *DynamicRedisDAOImpl) HashHIncrBy(ctx context.Context, key string, field string, increment int64) (int64, error) {
	result := r.client.HIncrBy(ctx, key, field, increment)
	return result.Val(), result.Err()
}