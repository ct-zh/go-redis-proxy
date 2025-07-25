package router

import (
	"github.com/gin-gonic/gin"

	"github.com/ct-zh/go-redis-proxy/internal/handler"
)

func Setup(engine *gin.Engine) {
	// 注册ping路由
	engine.GET("/ping", handler.Ping)
	engine.GET("/health", handler.Ping) // 健康检查使用相同处理器

	// Redis string
	apiV1 := engine.Group("/api/v1")
	{
		redis := apiV1.Group("/redis")
		{
			redis.POST("/string/get", handler.RedisStringGet(nil))
		}
	}
}
