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
				listGroup.POST("/lpush", container.RedisHandler.RedisListLPush)
				listGroup.POST("/rpush", container.RedisHandler.RedisListRPush)
				listGroup.POST("/lpop", container.RedisHandler.RedisListLPop)
				listGroup.POST("/rpop", container.RedisHandler.RedisListRPop)
				listGroup.POST("/lrem", container.RedisHandler.RedisListLRem)
				listGroup.POST("/lindex", container.RedisHandler.RedisListLIndex)
				listGroup.POST("/lrange", container.RedisHandler.RedisListLRange)
				listGroup.POST("/llen", container.RedisHandler.RedisListLLen)
				listGroup.POST("/ltrim", container.RedisHandler.RedisListLTrim)
			}
		}
	}
}