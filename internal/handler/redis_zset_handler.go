package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/ct-zh/go-redis-proxy/internal/service"
	"github.com/ct-zh/go-redis-proxy/pkg/types"
	"github.com/ct-zh/go-redis-proxy/pkg/response"
)

// RedisZSetHandler handles HTTP requests for Redis sorted set operations
type RedisZSetHandler struct {
	zsetService service.RedisZSetService
}

// NewRedisZSetHandler creates a new RedisZSetHandler instance
func NewRedisZSetHandler(zsetService service.RedisZSetService) *RedisZSetHandler {
	return &RedisZSetHandler{
		zsetService: zsetService,
	}
}

// RedisZSetZAdd godoc
// @Summary Redis有序集合ZADD操作
// @Description 将一个或多个成员元素及其分数值加入到有序集合中
// @Tags Redis ZSet Operations
// @Accept json
// @Produce json
// @Param request body types.ZSetZAddRequest true "请求参数"
// @Success 200 {object} response.BaseResponse "成功响应"
// @Failure 400 {object} response.BaseResponse "请求参数错误"
// @Failure 500 {object} response.BaseResponse "服务器内部错误"
// @Router /redis/zset/zadd [post]
func (h *RedisZSetHandler) RedisZSetZAdd(c *gin.Context) {
	var req types.ZSetZAddRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request parameters", err)
		return
	}

	// Validate required fields
	if req.Key == "" || len(req.Members) == 0 {
		response.BadRequest(c, "Key and members are required", nil)
		return
	}

	// Call service layer
	data, err := h.zsetService.ZAdd(c.Request.Context(), &req)
	response.JSON(c, data, err)
}

// RedisZSetZIncrBy godoc
// @Summary Redis有序集合ZINCRBY操作
// @Description 对有序集合中指定成员的分数加上增量increment
// @Tags Redis ZSet Operations
// @Accept json
// @Produce json
// @Param request body types.ZSetZIncrByRequest true "请求参数"
// @Success 200 {object} response.BaseResponse "成功响应"
// @Failure 400 {object} response.BaseResponse "请求参数错误"
// @Failure 500 {object} response.BaseResponse "服务器内部错误"
// @Router /redis/zset/zincrby [post]
func (h *RedisZSetHandler) RedisZSetZIncrBy(c *gin.Context) {
	var req types.ZSetZIncrByRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request parameters", err)
		return
	}

	// Validate required fields
	if req.Key == "" || req.Member == "" {
		response.BadRequest(c, "Key and member are required", nil)
		return
	}

	// Call service layer
	data, err := h.zsetService.ZIncrBy(c.Request.Context(), &req)
	response.JSON(c, data, err)
}

// RedisZSetZScore godoc
// @Summary Redis有序集合ZSCORE操作
// @Description 返回有序集合中指定成员的分数值
// @Tags Redis ZSet Operations
// @Accept json
// @Produce json
// @Param request body types.ZSetZScoreRequest true "请求参数"
// @Success 200 {object} response.BaseResponse "成功响应"
// @Failure 400 {object} response.BaseResponse "请求参数错误"
// @Failure 500 {object} response.BaseResponse "服务器内部错误"
// @Router /redis/zset/zscore [post]
func (h *RedisZSetHandler) RedisZSetZScore(c *gin.Context) {
	var req types.ZSetZScoreRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request parameters", err)
		return
	}

	// Validate required fields
	if req.Key == "" || req.Member == "" {
		response.BadRequest(c, "Key and member are required", nil)
		return
	}

	// Call service layer
	data, err := h.zsetService.ZScore(c.Request.Context(), &req)
	response.JSON(c, data, err)
}

// RedisZSetZCard godoc
// @Summary Redis有序集合ZCARD操作
// @Description 获取有序集合的成员数
// @Tags Redis ZSet Operations
// @Accept json
// @Produce json
// @Param request body types.ZSetZCardRequest true "请求参数"
// @Success 200 {object} response.BaseResponse "成功响应"
// @Failure 400 {object} response.BaseResponse "请求参数错误"
// @Failure 500 {object} response.BaseResponse "服务器内部错误"
// @Router /redis/zset/zcard [post]
func (h *RedisZSetHandler) RedisZSetZCard(c *gin.Context) {
	var req types.ZSetZCardRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request parameters", err)
		return
	}

	// Validate required fields
	if req.Key == "" {
		response.BadRequest(c, "Key is required", nil)
		return
	}

	// Call service layer
	data, err := h.zsetService.ZCard(c.Request.Context(), &req)
	response.JSON(c, data, err)
}

// RedisZSetZCount godoc
// @Summary Redis有序集合ZCOUNT操作
// @Description 计算在有序集合中指定区间分数的成员数
// @Tags Redis ZSet Operations
// @Accept json
// @Produce json
// @Param request body types.ZSetZCountRequest true "请求参数"
// @Success 200 {object} response.BaseResponse "成功响应"
// @Failure 400 {object} response.BaseResponse "请求参数错误"
// @Failure 500 {object} response.BaseResponse "服务器内部错误"
// @Router /redis/zset/zcount [post]
func (h *RedisZSetHandler) RedisZSetZCount(c *gin.Context) {
	var req types.ZSetZCountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request parameters", err)
		return
	}

	// Validate required fields
	if req.Key == "" {
		response.BadRequest(c, "Key is required", nil)
		return
	}

	// Call service layer
	data, err := h.zsetService.ZCount(c.Request.Context(), &req)
	response.JSON(c, data, err)
}

// RedisZSetZRank godoc
// @Summary Redis有序集合ZRANK操作
// @Description 返回有序集合中指定成员的排名，从0开始（分数从小到大）
// @Tags Redis ZSet Operations
// @Accept json
// @Produce json
// @Param request body types.ZSetZRankRequest true "请求参数"
// @Success 200 {object} response.BaseResponse "成功响应"
// @Failure 400 {object} response.BaseResponse "请求参数错误"
// @Failure 500 {object} response.BaseResponse "服务器内部错误"
// @Router /redis/zset/zrank [post]
func (h *RedisZSetHandler) RedisZSetZRank(c *gin.Context) {
	var req types.ZSetZRankRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request parameters", err)
		return
	}

	// Validate required fields
	if req.Key == "" || req.Member == "" {
		response.BadRequest(c, "Key and member are required", nil)
		return
	}

	// Call service layer
	data, err := h.zsetService.ZRank(c.Request.Context(), &req)
	response.JSON(c, data, err)
}

// RedisZSetZRevRank godoc
// @Summary Redis有序集合ZREVRANK操作
// @Description 返回有序集合中指定成员的排名，从0开始（分数从大到小）
// @Tags Redis ZSet Operations
// @Accept json
// @Produce json
// @Param request body types.ZSetZRevRankRequest true "请求参数"
// @Success 200 {object} response.BaseResponse "成功响应"
// @Failure 400 {object} response.BaseResponse "请求参数错误"
// @Failure 500 {object} response.BaseResponse "服务器内部错误"
// @Router /redis/zset/zrevrank [post]
func (h *RedisZSetHandler) RedisZSetZRevRank(c *gin.Context) {
	var req types.ZSetZRevRankRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request parameters", err)
		return
	}

	// Validate required fields
	if req.Key == "" || req.Member == "" {
		response.BadRequest(c, "Key and member are required", nil)
		return
	}

	// Call service layer
	data, err := h.zsetService.ZRevRank(c.Request.Context(), &req)
	response.JSON(c, data, err)
}

// RedisZSetZRange godoc
// @Summary Redis有序集合ZRANGE操作
// @Description 通过索引区间返回有序集合成指定区间内的成员（分数从小到大）
// @Tags Redis ZSet Operations
// @Accept json
// @Produce json
// @Param request body types.ZSetZRangeRequest true "请求参数"
// @Success 200 {object} response.BaseResponse "成功响应"
// @Failure 400 {object} response.BaseResponse "请求参数错误"
// @Failure 500 {object} response.BaseResponse "服务器内部错误"
// @Router /redis/zset/zrange [post]
func (h *RedisZSetHandler) RedisZSetZRange(c *gin.Context) {
	var req types.ZSetZRangeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request parameters", err)
		return
	}

	// Validate required fields
	if req.Key == "" {
		response.BadRequest(c, "Key is required", nil)
		return
	}

	// Call service layer
	data, err := h.zsetService.ZRange(c.Request.Context(), &req)
	response.JSON(c, data, err)
}

// RedisZSetZRevRange godoc
// @Summary Redis有序集合ZREVRANGE操作
// @Description 通过索引区间返回有序集合成指定区间内的成员（分数从大到小）
// @Tags Redis ZSet Operations
// @Accept json
// @Produce json
// @Param request body types.ZSetZRevRangeRequest true "请求参数"
// @Success 200 {object} response.BaseResponse "成功响应"
// @Failure 400 {object} response.BaseResponse "请求参数错误"
// @Failure 500 {object} response.BaseResponse "服务器内部错误"
// @Router /redis/zset/zrevrange [post]
func (h *RedisZSetHandler) RedisZSetZRevRange(c *gin.Context) {
	var req types.ZSetZRevRangeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request parameters", err)
		return
	}

	// Validate required fields
	if req.Key == "" {
		response.BadRequest(c, "Key is required", nil)
		return
	}

	// Call service layer
	data, err := h.zsetService.ZRevRange(c.Request.Context(), &req)
	response.JSON(c, data, err)
}

// RedisZSetZRangeByScore godoc
// @Summary Redis有序集合ZRANGEBYSCORE操作
// @Description 通过分数返回有序集合指定区间内的成员（分数从小到大）
// @Tags Redis ZSet Operations
// @Accept json
// @Produce json
// @Param request body types.ZSetZRangeByScoreRequest true "请求参数"
// @Success 200 {object} response.BaseResponse "成功响应"
// @Failure 400 {object} response.BaseResponse "请求参数错误"
// @Failure 500 {object} response.BaseResponse "服务器内部错误"
// @Router /redis/zset/zrangebyscore [post]
func (h *RedisZSetHandler) RedisZSetZRangeByScore(c *gin.Context) {
	var req types.ZSetZRangeByScoreRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request parameters", err)
		return
	}

	// Validate required fields
	if req.Key == "" || req.Min == "" || req.Max == "" {
		response.BadRequest(c, "Key, min and max are required", nil)
		return
	}

	// Call service layer
	data, err := h.zsetService.ZRangeByScore(c.Request.Context(), &req)
	response.JSON(c, data, err)
}

// RedisZSetZRevRangeByScore godoc
// @Summary Redis有序集合ZREVRANGEBYSCORE操作
// @Description 通过分数返回有序集合指定区间内的成员（分数从大到小）
// @Tags Redis ZSet Operations
// @Accept json
// @Produce json
// @Param request body types.ZSetZRevRangeByScoreRequest true "请求参数"
// @Success 200 {object} response.BaseResponse "成功响应"
// @Failure 400 {object} response.BaseResponse "请求参数错误"
// @Failure 500 {object} response.BaseResponse "服务器内部错误"
// @Router /redis/zset/zrevrangebyscore [post]
func (h *RedisZSetHandler) RedisZSetZRevRangeByScore(c *gin.Context) {
	var req types.ZSetZRevRangeByScoreRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request parameters", err)
		return
	}

	// Validate required fields
	if req.Key == "" || req.Min == "" || req.Max == "" {
		response.BadRequest(c, "Key, min and max are required", nil)
		return
	}

	// Call service layer
	data, err := h.zsetService.ZRevRangeByScore(c.Request.Context(), &req)
	response.JSON(c, data, err)
}

// RedisZSetZRem godoc
// @Summary Redis有序集合ZREM操作
// @Description 移除有序集合中的一个或多个成员
// @Tags Redis ZSet Operations
// @Accept json
// @Produce json
// @Param request body types.ZSetZRemRequest true "请求参数"
// @Success 200 {object} response.BaseResponse "成功响应"
// @Failure 400 {object} response.BaseResponse "请求参数错误"
// @Failure 500 {object} response.BaseResponse "服务器内部错误"
// @Router /redis/zset/zrem [post]
func (h *RedisZSetHandler) RedisZSetZRem(c *gin.Context) {
	var req types.ZSetZRemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request parameters", err)
		return
	}

	// Validate required fields
	if req.Key == "" || len(req.Members) == 0 {
		response.BadRequest(c, "Key and members are required", nil)
		return
	}

	// Call service layer
	data, err := h.zsetService.ZRem(c.Request.Context(), &req)
	response.JSON(c, data, err)
}

// RedisZSetZRemRangeByRank godoc
// @Summary Redis有序集合ZREMRANGEBYRANK操作
// @Description 移除有序集合中给定的排名区间的所有成员
// @Tags Redis ZSet Operations
// @Accept json
// @Produce json
// @Param request body types.ZSetZRemRangeByRankRequest true "请求参数"
// @Success 200 {object} response.BaseResponse "成功响应"
// @Failure 400 {object} response.BaseResponse "请求参数错误"
// @Failure 500 {object} response.BaseResponse "服务器内部错误"
// @Router /redis/zset/zremrangebyrank [post]
func (h *RedisZSetHandler) RedisZSetZRemRangeByRank(c *gin.Context) {
	var req types.ZSetZRemRangeByRankRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request parameters", err)
		return
	}

	// Validate required fields
	if req.Key == "" {
		response.BadRequest(c, "Key is required", nil)
		return
	}

	// Call service layer
	data, err := h.zsetService.ZRemRangeByRank(c.Request.Context(), &req)
	response.JSON(c, data, err)
}

// RedisZSetZRemRangeByScore godoc
// @Summary Redis有序集合ZREMRANGEBYSCORE操作
// @Description 移除有序集合中给定的分数区间的所有成员
// @Tags Redis ZSet Operations
// @Accept json
// @Produce json
// @Param request body types.ZSetZRemRangeByScoreRequest true "请求参数"
// @Success 200 {object} response.BaseResponse "成功响应"
// @Failure 400 {object} response.BaseResponse "请求参数错误"
// @Failure 500 {object} response.BaseResponse "服务器内部错误"
// @Router /redis/zset/zremrangebyscore [post]
func (h *RedisZSetHandler) RedisZSetZRemRangeByScore(c *gin.Context) {
	var req types.ZSetZRemRangeByScoreRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request parameters", err)
		return
	}

	// Validate required fields
	if req.Key == "" || req.Min == "" || req.Max == "" {
		response.BadRequest(c, "Key, min and max are required", nil)
		return
	}

	// Call service layer
	data, err := h.zsetService.ZRemRangeByScore(c.Request.Context(), &req)
	response.JSON(c, data, err)
}