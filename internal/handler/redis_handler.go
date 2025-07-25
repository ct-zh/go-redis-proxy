package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	redis "github.com/go-redis/redis/v8"

	"github.com/ct-zh/go-redis-proxy/pkg/types"
)

// RedisClient defines the interface for a Redis client. It's used to allow mocking in tests.
type RedisClient interface {
	Get(ctx context.Context, key string) *redis.StringCmd
}

func RedisStringGet(client RedisClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.StringGetRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// If the client is nil, create a new one. This allows the handler to be used in production
		// without having to pass a client to it.
		var rdb RedisClient
		if client == nil {
			rdb = redis.NewClient(&redis.Options{
				Addr:     req.Addr,
				Password: req.Password,
				DB:       req.DB,
			})
		} else {
			rdb = client
		}

		val, err := rdb.Get(context.Background(), req.Key).Result()
		if err != nil {
			if err == redis.Nil {
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "success",
					"data": gin.H{
						"value": nil,
					},
				})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error connecting to redis: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "success",
			"data": gin.H{
				"value": val,
			},
		})
	}
}
