package router

import (
	"github.com/gin-gonic/gin"

	"github.com/ct-zh/go-redis-proxy/internal/handler"
)

func Setup(engine *gin.Engine) {
	// 注册ping路由
	engine.GET("/ping", handler.Ping)
	engine.GET("/health", handler.Ping) // 健康检查使用相同处理器

	// Redis operations
	apiV1 := engine.Group("/api/v1")
	{
		redis := apiV1.Group("/redis")
		{
			// String operations
			stringGroup := redis.Group("/string")
			{
				stringGroup.POST("/get", handler.RedisStringGet(nil))
				stringGroup.POST("/set", handler.RedisStringSet(nil))
				stringGroup.POST("/del", handler.RedisStringDel(nil))
				stringGroup.POST("/exists", handler.RedisStringExists(nil))
				stringGroup.POST("/incr", handler.RedisStringIncr(nil))
				stringGroup.POST("/decr", handler.RedisStringDecr(nil))
				stringGroup.POST("/expire", handler.RedisStringExpire(nil))
			}

			// List operations
			listGroup := redis.Group("/list")
			{
				listGroup.POST("/lpush", handler.RedisListLPush(nil))
				listGroup.POST("/rpush", handler.RedisListRPush(nil))
				listGroup.POST("/lpop", handler.RedisListLPop(nil))
				listGroup.POST("/rpop", handler.RedisListRPop(nil))
				listGroup.POST("/lrem", handler.RedisListLRem(nil))
				listGroup.POST("/lindex", handler.RedisListLIndex(nil))
				listGroup.POST("/lrange", handler.RedisListLRange(nil))
				listGroup.POST("/llen", handler.RedisListLLen(nil))
				listGroup.POST("/ltrim", handler.RedisListLTrim(nil))
			}
		}
	}
}
