package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/ct-zh/go-redis-proxy/internal/service"
	"github.com/ct-zh/go-redis-proxy/pkg/types"
	"github.com/ct-zh/go-redis-proxy/pkg/response"
)

// RedisHashHandler handles HTTP requests for Redis hash operations
type RedisHashHandler struct {
	hashService service.RedisHashService
}

// NewRedisHashHandler creates a new RedisHashHandler instance
func NewRedisHashHandler(hashService service.RedisHashService) *RedisHashHandler {
	return &RedisHashHandler{
		hashService: hashService,
	}
}

// RedisHashHSet godoc
// @Summary Redis哈希表HSET操作
// @Description 为哈希表中的字段赋值
// @Tags Redis Hash Operations
// @Accept json
// @Produce json
// @Param request body types.HashHSetRequest true "请求参数"
// @Success 200 {object} response.BaseResponse "成功响应"
// @Failure 400 {object} response.BaseResponse "请求参数错误"
// @Failure 500 {object} response.BaseResponse "服务器内部错误"
// @Router /redis/hash/hset [post]
func (h *RedisHashHandler) RedisHashHSet(c *gin.Context) {
	var req types.HashHSetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request parameters", err)
		return
	}

	// Validate required fields
	if req.Key == "" || len(req.Fields) == 0 {
		response.BadRequest(c, "Key and fields are required", nil)
		return
	}

	// Call service layer
	data, err := h.hashService.HSet(c.Request.Context(), &req)
	response.JSON(c, data, err)
}

// RedisHashHGet godoc
// @Summary Redis哈希表HGET操作
// @Description 获取哈希表中指定字段的值
// @Tags Redis Hash Operations
// @Accept json
// @Produce json
// @Param request body types.HashHGetRequest true "请求参数"
// @Success 200 {object} response.BaseResponse "成功响应"
// @Failure 400 {object} response.BaseResponse "请求参数错误"
// @Failure 500 {object} response.BaseResponse "服务器内部错误"
// @Router /redis/hash/hget [post]
func (h *RedisHashHandler) RedisHashHGet(c *gin.Context) {
	var req types.HashHGetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request parameters", err)
		return
	}

	// Validate required fields
	if req.Key == "" || req.Field == "" {
		response.BadRequest(c, "Key and field are required", nil)
		return
	}

	// Call service layer
	data, err := h.hashService.HGet(c.Request.Context(), &req)
	response.JSON(c, data, err)
}

// RedisHashHMGet godoc
// @Summary Redis哈希表HMGET操作
// @Description 获取哈希表中一个或多个字段的值
// @Tags Redis Hash Operations
// @Accept json
// @Produce json
// @Param request body types.HashHMGetRequest true "请求参数"
// @Success 200 {object} response.BaseResponse "成功响应"
// @Failure 400 {object} response.BaseResponse "请求参数错误"
// @Failure 500 {object} response.BaseResponse "服务器内部错误"
// @Router /redis/hash/hmget [post]
func (h *RedisHashHandler) RedisHashHMGet(c *gin.Context) {
	var req types.HashHMGetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request parameters", err)
		return
	}

	// Validate required fields
	if req.Key == "" || len(req.Fields) == 0 {
		response.BadRequest(c, "Key and fields are required", nil)
		return
	}

	// Call service layer
	data, err := h.hashService.HMGet(c.Request.Context(), &req)
	response.JSON(c, data, err)
}

// RedisHashHGetAll godoc
// @Summary Redis哈希表HGETALL操作
// @Description 获取哈希表中所有的字段和值
// @Tags Redis Hash Operations
// @Accept json
// @Produce json
// @Param request body types.HashHGetAllRequest true "请求参数"
// @Success 200 {object} response.BaseResponse "成功响应"
// @Failure 400 {object} response.BaseResponse "请求参数错误"
// @Failure 500 {object} response.BaseResponse "服务器内部错误"
// @Router /redis/hash/hgetall [post]
func (h *RedisHashHandler) RedisHashHGetAll(c *gin.Context) {
	var req types.HashHGetAllRequest
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
	data, err := h.hashService.HGetAll(c.Request.Context(), &req)
	response.JSON(c, data, err)
}

// RedisHashHDel godoc
// @Summary Redis哈希表HDEL操作
// @Description 删除哈希表中一个或多个字段
// @Tags Redis Hash Operations
// @Accept json
// @Produce json
// @Param request body types.HashHDelRequest true "请求参数"
// @Success 200 {object} response.BaseResponse "成功响应"
// @Failure 400 {object} response.BaseResponse "请求参数错误"
// @Failure 500 {object} response.BaseResponse "服务器内部错误"
// @Router /redis/hash/hdel [post]
func (h *RedisHashHandler) RedisHashHDel(c *gin.Context) {
	var req types.HashHDelRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request parameters", err)
		return
	}

	// Validate required fields
	if req.Key == "" || len(req.Fields) == 0 {
		response.BadRequest(c, "Key and fields are required", nil)
		return
	}

	// Call service layer
	data, err := h.hashService.HDel(c.Request.Context(), &req)
	response.JSON(c, data, err)
}

// RedisHashHExists godoc
// @Summary Redis哈希表HEXISTS操作
// @Description 查看哈希表的指定字段是否存在
// @Tags Redis Hash Operations
// @Accept json
// @Produce json
// @Param request body types.HashHExistsRequest true "请求参数"
// @Success 200 {object} response.BaseResponse "成功响应"
// @Failure 400 {object} response.BaseResponse "请求参数错误"
// @Failure 500 {object} response.BaseResponse "服务器内部错误"
// @Router /redis/hash/hexists [post]
func (h *RedisHashHandler) RedisHashHExists(c *gin.Context) {
	var req types.HashHExistsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request parameters", err)
		return
	}

	// Validate required fields
	if req.Key == "" || req.Field == "" {
		response.BadRequest(c, "Key and field are required", nil)
		return
	}

	// Call service layer
	data, err := h.hashService.HExists(c.Request.Context(), &req)
	response.JSON(c, data, err)
}

// RedisHashHLen godoc
// @Summary Redis哈希表HLEN操作
// @Description 获取哈希表中字段的数量
// @Tags Redis Hash Operations
// @Accept json
// @Produce json
// @Param request body types.HashHLenRequest true "请求参数"
// @Success 200 {object} response.BaseResponse "成功响应"
// @Failure 400 {object} response.BaseResponse "请求参数错误"
// @Failure 500 {object} response.BaseResponse "服务器内部错误"
// @Router /redis/hash/hlen [post]
func (h *RedisHashHandler) RedisHashHLen(c *gin.Context) {
	var req types.HashHLenRequest
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
	data, err := h.hashService.HLen(c.Request.Context(), &req)
	response.JSON(c, data, err)
}

// RedisHashHKeys godoc
// @Summary Redis哈希表HKEYS操作
// @Description 获取哈希表中的所有字段名
// @Tags Redis Hash Operations
// @Accept json
// @Produce json
// @Param request body types.HashHKeysRequest true "请求参数"
// @Success 200 {object} response.BaseResponse "成功响应"
// @Failure 400 {object} response.BaseResponse "请求参数错误"
// @Failure 500 {object} response.BaseResponse "服务器内部错误"
// @Router /redis/hash/hkeys [post]
func (h *RedisHashHandler) RedisHashHKeys(c *gin.Context) {
	var req types.HashHKeysRequest
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
	data, err := h.hashService.HKeys(c.Request.Context(), &req)
	response.JSON(c, data, err)
}

// RedisHashHVals godoc
// @Summary Redis哈希表HVALS操作
// @Description 获取哈希表中所有字段的值
// @Tags Redis Hash Operations
// @Accept json
// @Produce json
// @Param request body types.HashHValsRequest true "请求参数"
// @Success 200 {object} response.BaseResponse "成功响应"
// @Failure 400 {object} response.BaseResponse "请求参数错误"
// @Failure 500 {object} response.BaseResponse "服务器内部错误"
// @Router /redis/hash/hvals [post]
func (h *RedisHashHandler) RedisHashHVals(c *gin.Context) {
	var req types.HashHValsRequest
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
	data, err := h.hashService.HVals(c.Request.Context(), &req)
	response.JSON(c, data, err)
}

// RedisHashHIncrBy godoc
// @Summary Redis哈希表HINCRBY操作
// @Description 为哈希表中的字段值加上指定增量值
// @Tags Redis Hash Operations
// @Accept json
// @Produce json
// @Param request body types.HashHIncrByRequest true "请求参数"
// @Success 200 {object} response.BaseResponse "成功响应"
// @Failure 400 {object} response.BaseResponse "请求参数错误"
// @Failure 500 {object} response.BaseResponse "服务器内部错误"
// @Router /redis/hash/hincrby [post]
func (h *RedisHashHandler) RedisHashHIncrBy(c *gin.Context) {
	var req types.HashHIncrByRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request parameters", err)
		return
	}

	// Validate required fields
	if req.Key == "" || req.Field == "" {
		response.BadRequest(c, "Key and field are required", nil)
		return
	}

	// Call service layer
	data, err := h.hashService.HIncrBy(c.Request.Context(), &req)
	response.JSON(c, data, err)
}