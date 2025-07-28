package dao

import (
	"context"
	"time"

	"github.com/ct-zh/go-redis-proxy/pkg/types"
)

// RedisDAO defines the data access interface for Redis operations
type RedisDAO interface {
	// Connection management
	Connect(config types.RedisRequest) error
	Close() error
	Ping(ctx context.Context) error

	// String operations
	StringGet(ctx context.Context, key string) (interface{}, error)
	StringSet(ctx context.Context, key string, value interface{}, ttl time.Duration) (string, error)
	StringDel(ctx context.Context, key string) (int64, error)
	StringExists(ctx context.Context, key string) (bool, error)
	StringIncr(ctx context.Context, key string) (int64, error)
	StringDecr(ctx context.Context, key string) (int64, error)
	StringExpire(ctx context.Context, key string, ttl time.Duration) (bool, error)

	// List operations
	ListLPush(ctx context.Context, key string, values []string) (int64, error)
	ListRPush(ctx context.Context, key string, values []string) (int64, error)
	ListLPop(ctx context.Context, key string) (interface{}, error)
	ListRPop(ctx context.Context, key string) (interface{}, error)
	ListLRem(ctx context.Context, key string, count int64, value string) (int64, error)
	ListLIndex(ctx context.Context, key string, index int64) (interface{}, error)
	ListLRange(ctx context.Context, key string, start, stop int64) ([]string, error)
	ListLLen(ctx context.Context, key string) (int64, error)
	ListLTrim(ctx context.Context, key string, start, stop int64) (string, error)
}

// RedisConnectionConfig holds the configuration for Redis connection
type RedisConnectionConfig struct {
	Addr     string
	Password string
	DB       int
}