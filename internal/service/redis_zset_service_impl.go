package service

import (
	"context"

	"github.com/ct-zh/go-redis-proxy/internal/dao"
	"github.com/ct-zh/go-redis-proxy/pkg/errors"
	"github.com/ct-zh/go-redis-proxy/pkg/types"
)

// RedisZSetServiceImpl implements the RedisZSetService interface
type RedisZSetServiceImpl struct {
	redisDAO dao.RedisDAO
}

// NewRedisZSetService creates a new RedisZSetServiceImpl instance
func NewRedisZSetService(redisDAO dao.RedisDAO) RedisZSetService {
	return &RedisZSetServiceImpl{
		redisDAO: redisDAO,
	}
}

// ZAdd adds members with scores to a sorted set
func (s *RedisZSetServiceImpl) ZAdd(ctx context.Context, req *types.ZSetZAddRequest) (*types.ZSetZAddData, error) {
	// Connect to Redis
	if err := s.redisDAO.Connect(req.RedisRequest); err != nil {
		return nil, errors.NewError(errors.CodeRedisConnectFailed)
	}
	defer s.redisDAO.Close()

	// Call DAO layer
	added, err := s.redisDAO.ZSetZAdd(ctx, req.Key, req.Members)
	if err != nil {
		return nil, errors.NewError(errors.CodeZSetAddFailed)
	}

	return &types.ZSetZAddData{Added: added}, nil
}

// ZIncrBy increments the score of a member in a sorted set
func (s *RedisZSetServiceImpl) ZIncrBy(ctx context.Context, req *types.ZSetZIncrByRequest) (*types.ZSetZIncrByData, error) {
	// Connect to Redis
	if err := s.redisDAO.Connect(req.RedisRequest); err != nil {
		return nil, errors.NewError(errors.CodeRedisConnectFailed)
	}
	defer s.redisDAO.Close()

	// Call DAO layer
	score, err := s.redisDAO.ZSetZIncrBy(ctx, req.Key, req.Increment, req.Member)
	if err != nil {
		return nil, errors.NewError(errors.CodeZSetAddFailed)
	}

	return &types.ZSetZIncrByData{Score: score}, nil
}

// ZScore gets the score of a member in a sorted set
func (s *RedisZSetServiceImpl) ZScore(ctx context.Context, req *types.ZSetZScoreRequest) (*types.ZSetZScoreData, error) {
	// Connect to Redis
	if err := s.redisDAO.Connect(req.RedisRequest); err != nil {
		return nil, errors.NewError(errors.CodeRedisConnectFailed)
	}
	defer s.redisDAO.Close()

	// Call DAO layer
	score, err := s.redisDAO.ZSetZScore(ctx, req.Key, req.Member)
	if err != nil {
		return nil, errors.NewError(errors.CodeZSetMemberNotFound)
	}

	return &types.ZSetZScoreData{Score: score}, nil
}

// ZCard gets the number of members in a sorted set
func (s *RedisZSetServiceImpl) ZCard(ctx context.Context, req *types.ZSetZCardRequest) (*types.ZSetZCardData, error) {
	// Connect to Redis
	if err := s.redisDAO.Connect(req.RedisRequest); err != nil {
		return nil, errors.NewError(errors.CodeRedisConnectFailed)
	}
	defer s.redisDAO.Close()

	// Call DAO layer
	count, err := s.redisDAO.ZSetZCard(ctx, req.Key)
	if err != nil {
		return nil, errors.NewError(errors.CodeZSetRankNotFound)
	}

	return &types.ZSetZCardData{Count: count}, nil
}

// ZCount counts members in a sorted set within a score range
func (s *RedisZSetServiceImpl) ZCount(ctx context.Context, req *types.ZSetZCountRequest) (*types.ZSetZCountData, error) {
	// Connect to Redis
	if err := s.redisDAO.Connect(req.RedisRequest); err != nil {
		return nil, errors.NewError(errors.CodeRedisConnectFailed)
	}
	defer s.redisDAO.Close()

	// Call DAO layer
	count, err := s.redisDAO.ZSetZCount(ctx, req.Key, req.Min, req.Max)
	if err != nil {
		return nil, errors.NewError(errors.CodeZSetRankNotFound)
	}

	return &types.ZSetZCountData{Count: count}, nil
}

// ZRank gets the rank of a member in a sorted set (ascending order)
func (s *RedisZSetServiceImpl) ZRank(ctx context.Context, req *types.ZSetZRankRequest) (*types.ZSetZRankData, error) {
	// Connect to Redis
	if err := s.redisDAO.Connect(req.RedisRequest); err != nil {
		return nil, errors.NewError(errors.CodeRedisConnectFailed)
	}
	defer s.redisDAO.Close()

	// Call DAO layer
	rank, err := s.redisDAO.ZSetZRank(ctx, req.Key, req.Member)
	if err != nil {
		return nil, errors.NewError(errors.CodeZSetRankNotFound)
	}

	return &types.ZSetZRankData{Rank: rank}, nil
}

// ZRevRank gets the rank of a member in a sorted set (descending order)
func (s *RedisZSetServiceImpl) ZRevRank(ctx context.Context, req *types.ZSetZRevRankRequest) (*types.ZSetZRevRankData, error) {
	// Connect to Redis
	if err := s.redisDAO.Connect(req.RedisRequest); err != nil {
		return nil, errors.NewError(errors.CodeRedisConnectFailed)
	}
	defer s.redisDAO.Close()

	// Call DAO layer
	rank, err := s.redisDAO.ZSetZRevRank(ctx, req.Key, req.Member)
	if err != nil {
		return nil, errors.NewError(errors.CodeZSetRankNotFound)
	}

	return &types.ZSetZRevRankData{Rank: rank}, nil
}

// ZRange gets members from a sorted set by rank range (ascending order)
func (s *RedisZSetServiceImpl) ZRange(ctx context.Context, req *types.ZSetZRangeRequest) (*types.ZSetZRangeData, error) {
	// Connect to Redis
	if err := s.redisDAO.Connect(req.RedisRequest); err != nil {
		return nil, errors.NewError(errors.CodeRedisConnectFailed)
	}
	defer s.redisDAO.Close()

	// Call DAO layer
	members, err := s.redisDAO.ZSetZRange(ctx, req.Key, req.Start, req.Stop, req.WithScores)
	if err != nil {
		return nil, errors.NewError(errors.CodeZSetRankNotFound)
	}

	return &types.ZSetZRangeData{Members: members}, nil
}

// ZRevRange gets members from a sorted set by rank range (descending order)
func (s *RedisZSetServiceImpl) ZRevRange(ctx context.Context, req *types.ZSetZRevRangeRequest) (*types.ZSetZRevRangeData, error) {
	// Connect to Redis
	if err := s.redisDAO.Connect(req.RedisRequest); err != nil {
		return nil, errors.NewError(errors.CodeRedisConnectFailed)
	}
	defer s.redisDAO.Close()

	// Call DAO layer
	members, err := s.redisDAO.ZSetZRevRange(ctx, req.Key, req.Start, req.Stop, req.WithScores)
	if err != nil {
		return nil, errors.NewError(errors.CodeZSetRankNotFound)
	}

	return &types.ZSetZRevRangeData{Members: members}, nil
}

// ZRangeByScore gets members from a sorted set by score range (ascending order)
func (s *RedisZSetServiceImpl) ZRangeByScore(ctx context.Context, req *types.ZSetZRangeByScoreRequest) (*types.ZSetZRangeByScoreData, error) {
	// Connect to Redis
	if err := s.redisDAO.Connect(req.RedisRequest); err != nil {
		return nil, errors.NewError(errors.CodeRedisConnectFailed)
	}
	defer s.redisDAO.Close()

	// Call DAO layer
	members, err := s.redisDAO.ZSetZRangeByScore(ctx, req.Key, req.Min, req.Max, req.WithScores, req.Offset, req.Count)
	if err != nil {
		return nil, errors.NewError(errors.CodeZSetRankNotFound)
	}

	return &types.ZSetZRangeByScoreData{Members: members}, nil
}

// ZRevRangeByScore gets members from a sorted set by score range (descending order)
func (s *RedisZSetServiceImpl) ZRevRangeByScore(ctx context.Context, req *types.ZSetZRevRangeByScoreRequest) (*types.ZSetZRevRangeByScoreData, error) {
	// Connect to Redis
	if err := s.redisDAO.Connect(req.RedisRequest); err != nil {
		return nil, errors.NewError(errors.CodeRedisConnectFailed)
	}
	defer s.redisDAO.Close()

	// Call DAO layer
	members, err := s.redisDAO.ZSetZRevRangeByScore(ctx, req.Key, req.Max, req.Min, req.WithScores, req.Offset, req.Count)
	if err != nil {
		return nil, errors.NewError(errors.CodeZSetRankNotFound)
	}

	return &types.ZSetZRevRangeByScoreData{Members: members}, nil
}

// ZRem removes members from a sorted set
func (s *RedisZSetServiceImpl) ZRem(ctx context.Context, req *types.ZSetZRemRequest) (*types.ZSetZRemData, error) {
	// Connect to Redis
	if err := s.redisDAO.Connect(req.RedisRequest); err != nil {
		return nil, errors.NewError(errors.CodeRedisConnectFailed)
	}
	defer s.redisDAO.Close()

	// Call DAO layer
	removed, err := s.redisDAO.ZSetZRem(ctx, req.Key, req.Members)
	if err != nil {
		return nil, errors.NewError(errors.CodeZSetRemoveFailed)
	}

	return &types.ZSetZRemData{Removed: removed}, nil
}

// ZRemRangeByRank removes members from a sorted set by rank range
func (s *RedisZSetServiceImpl) ZRemRangeByRank(ctx context.Context, req *types.ZSetZRemRangeByRankRequest) (*types.ZSetZRemRangeByRankData, error) {
	// Connect to Redis
	if err := s.redisDAO.Connect(req.RedisRequest); err != nil {
		return nil, errors.NewError(errors.CodeRedisConnectFailed)
	}
	defer s.redisDAO.Close()

	// Call DAO layer
	removed, err := s.redisDAO.ZSetZRemRangeByRank(ctx, req.Key, req.Start, req.Stop)
	if err != nil {
		return nil, errors.NewError(errors.CodeZSetRemoveFailed)
	}

	return &types.ZSetZRemRangeByRankData{Removed: removed}, nil
}

// ZRemRangeByScore removes members from a sorted set by score range
func (s *RedisZSetServiceImpl) ZRemRangeByScore(ctx context.Context, req *types.ZSetZRemRangeByScoreRequest) (*types.ZSetZRemRangeByScoreData, error) {
	// Connect to Redis
	if err := s.redisDAO.Connect(req.RedisRequest); err != nil {
		return nil, errors.NewError(errors.CodeRedisConnectFailed)
	}
	defer s.redisDAO.Close()

	// Call DAO layer
	removed, err := s.redisDAO.ZSetZRemRangeByScore(ctx, req.Key, req.Min, req.Max)
	if err != nil {
		return nil, errors.NewError(errors.CodeZSetRemoveFailed)
	}

	return &types.ZSetZRemRangeByScoreData{Removed: removed}, nil
}
