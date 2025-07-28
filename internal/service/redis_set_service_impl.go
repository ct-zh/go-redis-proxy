package service

import (
	"context"

	"github.com/ct-zh/go-redis-proxy/internal/dao"
	"github.com/ct-zh/go-redis-proxy/pkg/types"
)

// RedisSetServiceImpl implements the RedisSetService interface
type RedisSetServiceImpl struct {
	dao dao.RedisDAO
}

// NewRedisSetService creates a new instance of RedisSetServiceImpl
func NewRedisSetService(dao dao.RedisDAO) *RedisSetServiceImpl {
	return &RedisSetServiceImpl{dao: dao}
}

// SAdd adds members to a set
func (s *RedisSetServiceImpl) SAdd(ctx context.Context, req *types.RedisSAddRequest) (int64, error) {
	return s.dao.SetSAdd(ctx, req.Key, req.Members)
}

// SRem removes members from a set
func (s *RedisSetServiceImpl) SRem(ctx context.Context, req *types.RedisSRemRequest) (int64, error) {
	return s.dao.SetSRem(ctx, req.Key, req.Members)
}

// SIsMember checks if a member exists in a set
func (s *RedisSetServiceImpl) SIsMember(ctx context.Context, req *types.RedisSIsMemberRequest) (bool, error) {
	return s.dao.SetSIsMember(ctx, req.Key, req.Member)
}

// SMembers returns all members of a set
func (s *RedisSetServiceImpl) SMembers(ctx context.Context, req *types.RedisSMembersRequest) ([]string, error) {
	return s.dao.SetSMembers(ctx, req.Key)
}

// SCard returns the number of members in a set
func (s *RedisSetServiceImpl) SCard(ctx context.Context, req *types.RedisSCardRequest) (int64, error) {
	return s.dao.SetSCard(ctx, req.Key)
}
