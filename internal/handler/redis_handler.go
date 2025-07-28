package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	redis "github.com/go-redis/redis/v8"

	"github.com/ct-zh/go-redis-proxy/pkg/types"
)

// RedisClient defines the interface for a Redis client. It's used to allow mocking in tests.
type RedisClient interface {
	// String operations
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Del(ctx context.Context, keys ...string) *redis.IntCmd
	Exists(ctx context.Context, keys ...string) *redis.IntCmd
	Incr(ctx context.Context, key string) *redis.IntCmd
	Decr(ctx context.Context, key string) *redis.IntCmd
	Expire(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd
	
	// List operations
	LPush(ctx context.Context, key string, values ...interface{}) *redis.IntCmd
	RPush(ctx context.Context, key string, values ...interface{}) *redis.IntCmd
	LPop(ctx context.Context, key string) *redis.StringCmd
	RPop(ctx context.Context, key string) *redis.StringCmd
	LRem(ctx context.Context, key string, count int64, value interface{}) *redis.IntCmd
	LIndex(ctx context.Context, key string, index int64) *redis.StringCmd
	LRange(ctx context.Context, key string, start, stop int64) *redis.StringSliceCmd
	LLen(ctx context.Context, key string) *redis.IntCmd
	LTrim(ctx context.Context, key string, start, stop int64) *redis.StatusCmd
}

// createRedisClient 创建Redis客户端连接
func createRedisClient(client RedisClient, req types.RedisRequest) RedisClient {
	if client == nil {
		return redis.NewClient(&redis.Options{
			Addr:     req.Addr,
			Password: req.Password,
			DB:       req.DB,
		})
	}
	return client
}

// RedisStringGet godoc
// @Summary Redis字符串GET操作
// @Description 根据指定的key获取Redis中存储的字符串值
// @Tags Redis String Operations
// @Accept json
// @Produce json
// @Param request body types.StringGetRequest true "请求参数"
// @Success 200 {object} types.StringGetResponse "成功响应"
// @Failure 400 {object} types.ErrorResponse "请求参数错误"
// @Failure 500 {object} types.ErrorResponse "服务器内部错误"
// @Router /redis/string/get [post]
func RedisStringGet(client RedisClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.StringGetRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		rdb := createRedisClient(client, req.RedisRequest)

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

// RedisStringSet godoc
// @Summary Redis字符串SET操作
// @Description 设置指定key的字符串值，支持TTL过期时间
// @Tags Redis String Operations
// @Accept json
// @Produce json
// @Param request body types.StringSetRequest true "请求参数"
// @Success 200 {object} types.StringSetResponse "成功响应"
// @Failure 400 {object} types.ErrorResponse "请求参数错误"
// @Failure 500 {object} types.ErrorResponse "服务器内部错误"
// @Router /redis/string/set [post]
func RedisStringSet(client RedisClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.StringSetRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		rdb := createRedisClient(client, req.RedisRequest)

		var expiration time.Duration
		if req.TTL > 0 {
			expiration = time.Duration(req.TTL) * time.Second
		}

		err := rdb.Set(context.Background(), req.Key, req.Value, expiration).Err()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error setting redis key: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "success",
			"data": gin.H{
				"result": "OK",
			},
		})
	}
}

// RedisStringDel godoc
// @Summary Redis字符串DEL操作
// @Description 删除指定的key
// @Tags Redis String Operations
// @Accept json
// @Produce json
// @Param request body types.StringDelRequest true "请求参数"
// @Success 200 {object} types.StringDelResponse "成功响应"
// @Failure 400 {object} types.ErrorResponse "请求参数错误"
// @Failure 500 {object} types.ErrorResponse "服务器内部错误"
// @Router /redis/string/del [post]
func RedisStringDel(client RedisClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.StringDelRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		rdb := createRedisClient(client, req.RedisRequest)

		deleted, err := rdb.Del(context.Background(), req.Key).Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error deleting redis key: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "success",
			"data": gin.H{
				"deleted": int(deleted),
			},
		})
	}
}

// RedisStringExists godoc
// @Summary Redis字符串EXISTS操作
// @Description 检查指定key是否存在
// @Tags Redis String Operations
// @Accept json
// @Produce json
// @Param request body types.StringExistsRequest true "请求参数"
// @Success 200 {object} types.StringExistsResponse "成功响应"
// @Failure 400 {object} types.ErrorResponse "请求参数错误"
// @Failure 500 {object} types.ErrorResponse "服务器内部错误"
// @Router /redis/string/exists [post]
func RedisStringExists(client RedisClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.StringExistsRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		rdb := createRedisClient(client, req.RedisRequest)

		exists, err := rdb.Exists(context.Background(), req.Key).Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error checking redis key existence: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "success",
			"data": gin.H{
				"exists": exists > 0,
			},
		})
	}
}

// RedisStringIncr godoc
// @Summary Redis字符串INCR操作
// @Description 将指定key的值增加1
// @Tags Redis String Operations
// @Accept json
// @Produce json
// @Param request body types.StringIncrRequest true "请求参数"
// @Success 200 {object} types.StringIncrResponse "成功响应"
// @Failure 400 {object} types.ErrorResponse "请求参数错误"
// @Failure 500 {object} types.ErrorResponse "服务器内部错误"
// @Router /redis/string/incr [post]
func RedisStringIncr(client RedisClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.StringIncrRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		rdb := createRedisClient(client, req.RedisRequest)

		value, err := rdb.Incr(context.Background(), req.Key).Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error incrementing redis key: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "success",
			"data": gin.H{
				"value": value,
			},
		})
	}
}

// RedisStringDecr godoc
// @Summary Redis字符串DECR操作
// @Description 将指定key的值减少1
// @Tags Redis String Operations
// @Accept json
// @Produce json
// @Param request body types.StringDecrRequest true "请求参数"
// @Success 200 {object} types.StringDecrResponse "成功响应"
// @Failure 400 {object} types.ErrorResponse "请求参数错误"
// @Failure 500 {object} types.ErrorResponse "服务器内部错误"
// @Router /redis/string/decr [post]
func RedisStringDecr(client RedisClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.StringDecrRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		rdb := createRedisClient(client, req.RedisRequest)

		value, err := rdb.Decr(context.Background(), req.Key).Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error decrementing redis key: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "success",
			"data": gin.H{
				"value": value,
			},
		})
	}
}

// RedisStringExpire godoc
// @Summary Redis字符串EXPIRE操作
// @Description 为指定key设置过期时间
// @Tags Redis String Operations
// @Accept json
// @Produce json
// @Param request body types.StringExpireRequest true "请求参数"
// @Success 200 {object} types.StringExpireResponse "成功响应"
// @Failure 400 {object} types.ErrorResponse "请求参数错误"
// @Failure 500 {object} types.ErrorResponse "服务器内部错误"
// @Router /redis/string/expire [post]
func RedisStringExpire(client RedisClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.StringExpireRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		rdb := createRedisClient(client, req.RedisRequest)

		success, err := rdb.Expire(context.Background(), req.Key, time.Duration(req.TTL)*time.Second).Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error setting expiration for redis key: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "success",
			"data": gin.H{
				"success": success,
			},
		})
	}
}

// RedisListLPush godoc
// @Summary Redis列表LPUSH操作
// @Description 从列表左侧推入一个或多个值
// @Tags Redis List Operations
// @Accept json
// @Produce json
// @Param request body types.ListLPushRequest true "请求参数"
// @Success 200 {object} types.ListLPushResponse "成功响应"
// @Failure 400 {object} types.ErrorResponse "请求参数错误"
// @Failure 500 {object} types.ErrorResponse "服务器内部错误"
// @Router /redis/list/lpush [post]
func RedisListLPush(client RedisClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.ListLPushRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		rdb := createRedisClient(client, req.RedisRequest)

		// 转换string slice为interface slice
		values := make([]interface{}, len(req.Values))
		for i, v := range req.Values {
			values[i] = v
		}

		length, err := rdb.LPush(context.Background(), req.Key, values...).Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error pushing to redis list: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "success",
			"data": gin.H{
				"length": length,
			},
		})
	}
}

// RedisListRPush godoc
// @Summary Redis列表RPUSH操作
// @Description 从列表右侧推入一个或多个值
// @Tags Redis List Operations
// @Accept json
// @Produce json
// @Param request body types.ListRPushRequest true "请求参数"
// @Success 200 {object} types.ListRPushResponse "成功响应"
// @Failure 400 {object} types.ErrorResponse "请求参数错误"
// @Failure 500 {object} types.ErrorResponse "服务器内部错误"
// @Router /redis/list/rpush [post]
func RedisListRPush(client RedisClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.ListRPushRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		rdb := createRedisClient(client, req.RedisRequest)

		// 转换string slice为interface slice
		values := make([]interface{}, len(req.Values))
		for i, v := range req.Values {
			values[i] = v
		}

		length, err := rdb.RPush(context.Background(), req.Key, values...).Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error pushing to redis list: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "success",
			"data": gin.H{
				"length": length,
			},
		})
	}
}

// RedisListLPop godoc
// @Summary Redis列表LPOP操作
// @Description 从列表左侧弹出一个值
// @Tags Redis List Operations
// @Accept json
// @Produce json
// @Param request body types.ListLPopRequest true "请求参数"
// @Success 200 {object} types.ListLPopResponse "成功响应"
// @Failure 400 {object} types.ErrorResponse "请求参数错误"
// @Failure 500 {object} types.ErrorResponse "服务器内部错误"
// @Router /redis/list/lpop [post]
func RedisListLPop(client RedisClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.ListLPopRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		rdb := createRedisClient(client, req.RedisRequest)

		value, err := rdb.LPop(context.Background(), req.Key).Result()
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
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error popping from redis list: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "success",
			"data": gin.H{
				"value": value,
			},
		})
	}
}

// RedisListRPop godoc
// @Summary Redis列表RPOP操作
// @Description 从列表右侧弹出一个值
// @Tags Redis List Operations
// @Accept json
// @Produce json
// @Param request body types.ListRPopRequest true "请求参数"
// @Success 200 {object} types.ListRPopResponse "成功响应"
// @Failure 400 {object} types.ErrorResponse "请求参数错误"
// @Failure 500 {object} types.ErrorResponse "服务器内部错误"
// @Router /redis/list/rpop [post]
func RedisListRPop(client RedisClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.ListRPopRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		rdb := createRedisClient(client, req.RedisRequest)

		value, err := rdb.RPop(context.Background(), req.Key).Result()
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
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error popping from redis list: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "success",
			"data": gin.H{
				"value": value,
			},
		})
	}
}

// RedisListLRem godoc
// @Summary Redis列表LREM操作
// @Description 从列表中删除指定数量的元素
// @Tags Redis List Operations
// @Accept json
// @Produce json
// @Param request body types.ListLRemRequest true "请求参数"
// @Success 200 {object} types.ListLRemResponse "成功响应"
// @Failure 400 {object} types.ErrorResponse "请求参数错误"
// @Failure 500 {object} types.ErrorResponse "服务器内部错误"
// @Router /redis/list/lrem [post]
func RedisListLRem(client RedisClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.ListLRemRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		rdb := createRedisClient(client, req.RedisRequest)

		removed, err := rdb.LRem(context.Background(), req.Key, req.Count, req.Value).Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error removing from redis list: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "success",
			"data": gin.H{
				"removed": removed,
			},
		})
	}
}

// RedisListLIndex godoc
// @Summary Redis列表LINDEX操作
// @Description 获取列表指定索引位置的元素
// @Tags Redis List Operations
// @Accept json
// @Produce json
// @Param request body types.ListLIndexRequest true "请求参数"
// @Success 200 {object} types.ListLIndexResponse "成功响应"
// @Failure 400 {object} types.ErrorResponse "请求参数错误"
// @Failure 500 {object} types.ErrorResponse "服务器内部错误"
// @Router /redis/list/lindex [post]
func RedisListLIndex(client RedisClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.ListLIndexRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		rdb := createRedisClient(client, req.RedisRequest)

		value, err := rdb.LIndex(context.Background(), req.Key, req.Index).Result()
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
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error getting list index: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "success",
			"data": gin.H{
				"value": value,
			},
		})
	}
}

// RedisListLRange godoc
// @Summary Redis列表LRANGE操作
// @Description 获取列表指定范围内的元素
// @Tags Redis List Operations
// @Accept json
// @Produce json
// @Param request body types.ListLRangeRequest true "请求参数"
// @Success 200 {object} types.ListLRangeResponse "成功响应"
// @Failure 400 {object} types.ErrorResponse "请求参数错误"
// @Failure 500 {object} types.ErrorResponse "服务器内部错误"
// @Router /redis/list/lrange [post]
func RedisListLRange(client RedisClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.ListLRangeRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		rdb := createRedisClient(client, req.RedisRequest)

		values, err := rdb.LRange(context.Background(), req.Key, req.Start, req.Stop).Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error getting list range: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "success",
			"data": gin.H{
				"values": values,
			},
		})
	}
}

// RedisListLLen godoc
// @Summary Redis列表LLEN操作
// @Description 获取列表的长度
// @Tags Redis List Operations
// @Accept json
// @Produce json
// @Param request body types.ListLLenRequest true "请求参数"
// @Success 200 {object} types.ListLLenResponse "成功响应"
// @Failure 400 {object} types.ErrorResponse "请求参数错误"
// @Failure 500 {object} types.ErrorResponse "服务器内部错误"
// @Router /redis/list/llen [post]
func RedisListLLen(client RedisClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.ListLLenRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		rdb := createRedisClient(client, req.RedisRequest)

		length, err := rdb.LLen(context.Background(), req.Key).Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error getting list length: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "success",
			"data": gin.H{
				"length": length,
			},
		})
	}
}

// RedisListLTrim godoc
// @Summary Redis列表LTRIM操作
// @Description 修剪列表，只保留指定范围内的元素
// @Tags Redis List Operations
// @Accept json
// @Produce json
// @Param request body types.ListLTrimRequest true "请求参数"
// @Success 200 {object} types.ListLTrimResponse "成功响应"
// @Failure 400 {object} types.ErrorResponse "请求参数错误"
// @Failure 500 {object} types.ErrorResponse "服务器内部错误"
// @Router /redis/list/ltrim [post]
func RedisListLTrim(client RedisClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.ListLTrimRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		rdb := createRedisClient(client, req.RedisRequest)

		err := rdb.LTrim(context.Background(), req.Key, req.Start, req.Stop).Err()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error trimming list: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "success",
			"data": gin.H{
				"result": "OK",
			},
		})
	}
}
