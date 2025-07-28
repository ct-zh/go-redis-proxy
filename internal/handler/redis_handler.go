package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/ct-zh/go-redis-proxy/internal/service"
	"github.com/ct-zh/go-redis-proxy/pkg/types"
	"github.com/ct-zh/go-redis-proxy/pkg/response"
)

// RedisHandler handles HTTP requests for Redis operations
type RedisHandler struct {
	stringService service.RedisStringService
	listService   service.RedisListService
}

// NewRedisHandler creates a new RedisHandler instance
func NewRedisHandler(stringService service.RedisStringService, listService service.RedisListService) *RedisHandler {
	return &RedisHandler{
		stringService: stringService,
		listService:   listService,
	}
}

// String operations

// RedisStringGet godoc
// @Summary Redis字符串GET操作
// @Description 根据指定的key获取Redis中存储的字符串值
// @Tags Redis String Operations
// @Accept json
// @Produce json
// @Param request body types.StringGetRequest true "请求参数"
// @Success 200 {object} response.BaseResponse "成功响应"
// @Failure 400 {object} response.BaseResponse "请求参数错误"
// @Failure 500 {object} response.BaseResponse "服务器内部错误"
// @Router /redis/string/get [post]
func (h *RedisHandler) RedisStringGet(c *gin.Context) {
	var req types.StringGetRequest
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
	data, err := h.stringService.Get(c.Request.Context(), &req)
	response.JSON(c, data, err)
}

// RedisStringSet godoc
// @Summary Redis字符串SET操作
// @Description 设置指定key的字符串值，支持TTL过期时间
// @Tags Redis String Operations
// @Accept json
// @Produce json
// @Param request body types.StringSetRequest true "请求参数"
// @Success 200 {object} response.BaseResponse "成功响应"
// @Failure 400 {object} response.BaseResponse "请求参数错误"
// @Failure 500 {object} response.BaseResponse "服务器内部错误"
// @Router /redis/string/set [post]
func (h *RedisHandler) RedisStringSet(c *gin.Context) {
	var req types.StringSetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request parameters", err)
		return
	}

	// Validate required fields
	if req.Key == "" || req.Value == "" {
		response.BadRequest(c, "Key and value are required", nil)
		return
	}

	// Call service layer
	data, err := h.stringService.Set(c.Request.Context(), &req)
	response.JSON(c, data, err)
}

// RedisStringDel godoc
// @Summary Redis字符串DEL操作
// @Description 删除指定的key
// @Tags Redis String Operations
// @Accept json
// @Produce json
// @Param request body types.StringDelRequest true "请求参数"
// @Success 200 {object} response.BaseResponse "成功响应"
// @Failure 400 {object} response.BaseResponse "请求参数错误"
// @Failure 500 {object} response.BaseResponse "服务器内部错误"
// @Router /redis/string/del [post]
func (h *RedisHandler) RedisStringDel(c *gin.Context) {
	var req types.StringDelRequest
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
	data, err := h.stringService.Del(c.Request.Context(), &req)
	response.JSON(c, data, err)
}

// RedisStringExists godoc
// @Summary Redis字符串EXISTS操作
// @Description 检查指定key是否存在
// @Tags Redis String Operations
// @Accept json
// @Produce json
// @Param request body types.StringExistsRequest true "请求参数"
// @Success 200 {object} response.BaseResponse "成功响应"
// @Failure 400 {object} response.BaseResponse "请求参数错误"
// @Failure 500 {object} response.BaseResponse "服务器内部错误"
// @Router /redis/string/exists [post]
func (h *RedisHandler) RedisStringExists(c *gin.Context) {
	var req types.StringExistsRequest
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
	data, err := h.stringService.Exists(c.Request.Context(), &req)
	response.JSON(c, data, err)
}

// RedisStringIncr godoc
// @Summary Redis字符串INCR操作
// @Description 将指定key的值加1
// @Tags Redis String Operations
// @Accept json
// @Produce json
// @Param request body types.StringIncrRequest true "请求参数"
// @Success 200 {object} response.BaseResponse "成功响应"
// @Failure 400 {object} response.BaseResponse "请求参数错误"
// @Failure 500 {object} response.BaseResponse "服务器内部错误"
// @Router /redis/string/incr [post]
func (h *RedisHandler) RedisStringIncr(c *gin.Context) {
	var req types.StringIncrRequest
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
	data, err := h.stringService.Incr(c.Request.Context(), &req)
	response.JSON(c, data, err)
}

// RedisStringDecr godoc
// @Summary Redis字符串DECR操作
// @Description 将指定key的值减1
// @Tags Redis String Operations
// @Accept json
// @Produce json
// @Param request body types.StringDecrRequest true "请求参数"
// @Success 200 {object} response.BaseResponse "成功响应"
// @Failure 400 {object} response.BaseResponse "请求参数错误"
// @Failure 500 {object} response.BaseResponse "服务器内部错误"
// @Router /redis/string/decr [post]
func (h *RedisHandler) RedisStringDecr(c *gin.Context) {
	var req types.StringDecrRequest
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
	data, err := h.stringService.Decr(c.Request.Context(), &req)
	response.JSON(c, data, err)
}

// RedisStringExpire godoc
// @Summary Redis字符串EXPIRE操作
// @Description 为指定key设置过期时间
// @Tags Redis String Operations
// @Accept json
// @Produce json
// @Param request body types.StringExpireRequest true "请求参数"
// @Success 200 {object} response.BaseResponse "成功响应"
// @Failure 400 {object} response.BaseResponse "请求参数错误"
// @Failure 500 {object} response.BaseResponse "服务器内部错误"
// @Router /redis/string/expire [post]
func (h *RedisHandler) RedisStringExpire(c *gin.Context) {
	var req types.StringExpireRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request parameters", err)
		return
	}

	// Validate required fields
	if req.Key == "" || req.TTL <= 0 {
		response.BadRequest(c, "Key and valid TTL are required", nil)
		return
	}

	// Call service layer
	data, err := h.stringService.Expire(c.Request.Context(), &req)
	response.JSON(c, data, err)
}