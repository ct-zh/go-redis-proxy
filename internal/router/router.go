package router

import (
	"github.com/gin-gonic/gin"

	"github.com/ct-zh/go-redis-proxy/internal/container"
	"github.com/ct-zh/go-redis-proxy/internal/handler"
)

// SetupWithContainer sets up routes using dependency injection container
func SetupWithContainer(engine *gin.Engine, container *container.Container) {
	// Health check endpoint
	engine.GET("/ping", handler.Ping)

	// API v1 group
	api := engine.Group("/api/v1")
	{
		// Redis operations
		redis := api.Group("/redis")
		{
			// String operations
			stringGroup := redis.Group("/string")
			{
				stringGroup.POST("/get", container.RedisHandler.RedisStringGet)
				stringGroup.POST("/set", container.RedisHandler.RedisStringSet)
				stringGroup.POST("/del", container.RedisHandler.RedisStringDel)
				stringGroup.POST("/exists", container.RedisHandler.RedisStringExists)
				stringGroup.POST("/incr", container.RedisHandler.RedisStringIncr)
				stringGroup.POST("/decr", container.RedisHandler.RedisStringDecr)
				stringGroup.POST("/expire", container.RedisHandler.RedisStringExpire)
			}

			// List operations
			listGroup := redis.Group("/list")
			{
				listGroup.POST("/lpush", container.RedisListHandler.RedisListLPush)
				listGroup.POST("/rpush", container.RedisListHandler.RedisListRPush)
				listGroup.POST("/lpop", container.RedisListHandler.RedisListLPop)
				listGroup.POST("/rpop", container.RedisListHandler.RedisListRPop)
				listGroup.POST("/lrem", container.RedisListHandler.RedisListLRem)
				listGroup.POST("/lindex", container.RedisListHandler.RedisListLIndex)
				listGroup.POST("/lrange", container.RedisListHandler.RedisListLRange)
				listGroup.POST("/llen", container.RedisListHandler.RedisListLLen)
				listGroup.POST("/ltrim", container.RedisListHandler.RedisListLTrim)
			}

			// Set operations
			setGroup := redis.Group("/set")
			{
				setGroup.POST("/sadd", container.RedisSetHandler.SAdd)
				setGroup.POST("/srem", container.RedisSetHandler.SRem)
				setGroup.POST("/sismember", container.RedisSetHandler.SIsMember)
				setGroup.POST("/smembers", container.RedisSetHandler.SMembers)
				setGroup.POST("/scard", container.RedisSetHandler.SCard)
			}

			// ZSet operations
			zsetGroup := redis.Group("/zset")
			{
				zsetGroup.POST("/zadd", container.RedisZSetHandler.RedisZSetZAdd)
				zsetGroup.POST("/zincrby", container.RedisZSetHandler.RedisZSetZIncrBy)
				zsetGroup.POST("/zscore", container.RedisZSetHandler.RedisZSetZScore)
				zsetGroup.POST("/zcard", container.RedisZSetHandler.RedisZSetZCard)
				zsetGroup.POST("/zcount", container.RedisZSetHandler.RedisZSetZCount)
				zsetGroup.POST("/zrank", container.RedisZSetHandler.RedisZSetZRank)
				zsetGroup.POST("/zrevrank", container.RedisZSetHandler.RedisZSetZRevRank)
				zsetGroup.POST("/zrange", container.RedisZSetHandler.RedisZSetZRange)
				zsetGroup.POST("/zrevrange", container.RedisZSetHandler.RedisZSetZRevRange)
				zsetGroup.POST("/zrangebyscore", container.RedisZSetHandler.RedisZSetZRangeByScore)
				zsetGroup.POST("/zrevrangebyscore", container.RedisZSetHandler.RedisZSetZRevRangeByScore)
				zsetGroup.POST("/zrem", container.RedisZSetHandler.RedisZSetZRem)
				zsetGroup.POST("/zremrangebyrank", container.RedisZSetHandler.RedisZSetZRemRangeByRank)
				zsetGroup.POST("/zremrangebyscore", container.RedisZSetHandler.RedisZSetZRemRangeByScore)
			}
		}
	}
}