package service

import (
	"context"

	"github.com/ct-zh/go-redis-proxy/pkg/types"
)

// RedisStringService defines the business logic interface for Redis string operations
// 返回业务数据和错误，Handler层负责包装响应格式
type RedisStringService interface {
	Get(ctx context.Context, req *types.StringGetRequest) (*types.StringGetData, error)
	Set(ctx context.Context, req *types.StringSetRequest) (*types.StringSetData, error)
	Del(ctx context.Context, req *types.StringDelRequest) (*types.StringDelData, error)
	Exists(ctx context.Context, req *types.StringExistsRequest) (*types.StringExistsData, error)
	Incr(ctx context.Context, req *types.StringIncrRequest) (*types.StringIncrData, error)
	Decr(ctx context.Context, req *types.StringDecrRequest) (*types.StringDecrData, error)
	Expire(ctx context.Context, req *types.StringExpireRequest) (*types.StringExpireData, error)
}

// RedisListService defines the business logic interface for Redis list operations
// 返回业务数据和错误，Handler层负责包装响应格式
type RedisListService interface {
	LPush(ctx context.Context, req *types.ListLPushRequest) (*types.ListLPushData, error)
	RPush(ctx context.Context, req *types.ListRPushRequest) (*types.ListRPushData, error)
	LPop(ctx context.Context, req *types.ListLPopRequest) (*types.ListLPopData, error)
	RPop(ctx context.Context, req *types.ListRPopRequest) (*types.ListRPopData, error)
	LRem(ctx context.Context, req *types.ListLRemRequest) (*types.ListLRemData, error)
	LIndex(ctx context.Context, req *types.ListLIndexRequest) (*types.ListLIndexData, error)
	LRange(ctx context.Context, req *types.ListLRangeRequest) (*types.ListLRangeData, error)
	LLen(ctx context.Context, req *types.ListLLenRequest) (*types.ListLLenData, error)
	LTrim(ctx context.Context, req *types.ListLTrimRequest) (*types.ListLTrimData, error)
}

// RedisZSetService defines the business logic interface for Redis sorted set operations
// 返回业务数据和错误，Handler层负责包装响应格式
type RedisZSetService interface {
	ZAdd(ctx context.Context, req *types.ZSetZAddRequest) (*types.ZSetZAddData, error)
	ZIncrBy(ctx context.Context, req *types.ZSetZIncrByRequest) (*types.ZSetZIncrByData, error)
	ZScore(ctx context.Context, req *types.ZSetZScoreRequest) (*types.ZSetZScoreData, error)
	ZCard(ctx context.Context, req *types.ZSetZCardRequest) (*types.ZSetZCardData, error)
	ZCount(ctx context.Context, req *types.ZSetZCountRequest) (*types.ZSetZCountData, error)
	ZRank(ctx context.Context, req *types.ZSetZRankRequest) (*types.ZSetZRankData, error)
	ZRevRank(ctx context.Context, req *types.ZSetZRevRankRequest) (*types.ZSetZRevRankData, error)
	ZRange(ctx context.Context, req *types.ZSetZRangeRequest) (*types.ZSetZRangeData, error)
	ZRevRange(ctx context.Context, req *types.ZSetZRevRangeRequest) (*types.ZSetZRevRangeData, error)
	ZRangeByScore(ctx context.Context, req *types.ZSetZRangeByScoreRequest) (*types.ZSetZRangeByScoreData, error)
	ZRevRangeByScore(ctx context.Context, req *types.ZSetZRevRangeByScoreRequest) (*types.ZSetZRevRangeByScoreData, error)
	ZRem(ctx context.Context, req *types.ZSetZRemRequest) (*types.ZSetZRemData, error)
	ZRemRangeByRank(ctx context.Context, req *types.ZSetZRemRangeByRankRequest) (*types.ZSetZRemRangeByRankData, error)
	ZRemRangeByScore(ctx context.Context, req *types.ZSetZRemRangeByScoreRequest) (*types.ZSetZRemRangeByScoreData, error)
}