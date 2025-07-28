package service

import (
	"context"

	"github.com/ct-zh/go-redis-proxy/pkg/types"
)

// RedisSetService defines the interface for set operations
type RedisSetService interface {
	SAdd(ctx context.Context, req *types.RedisSAddRequest) (int64, error)
	SRem(ctx context.Context, req *types.RedisSRemRequest) (int64, error)
	SIsMember(ctx context.Context, req *types.RedisSIsMemberRequest) (bool, error)
	SMembers(ctx context.Context, req *types.RedisSMembersRequest) ([]string, error)
	SCard(ctx context.Context, req *types.RedisSCardRequest) (int64, error)
}
