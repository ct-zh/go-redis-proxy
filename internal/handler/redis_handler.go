package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ct-zh/go-redis-proxy/internal/service"
	"github.com/ct-zh/go-redis-proxy/pkg/types"
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
// @Success 200 {object} types.StringGetResponse "成功响应"
// @Failure 400 {object} types.ErrorResponse "请求参数错误"
// @Failure 500 {object} types.ErrorResponse "服务器内部错误"
// @Router /redis/string/get [post]
func (h *RedisHandler) RedisStringGet(c *gin.Context) {
	var req types.StringGetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: "Invalid request parameters"})
		return
	}

	// Validate required fields
	if req.Key == "" {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: "Key is required"})
		return
	}

	// Call service layer
	response, err := h.stringService.Get(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
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
func (h *RedisHandler) RedisStringSet(c *gin.Context) {
	var req types.StringSetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: "Invalid request parameters"})
		return
	}

	// Validate required fields
	if req.Key == "" {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: "Key is required"})
		return
	}

	// Call service layer
	response, err := h.stringService.Set(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
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
func (h *RedisHandler) RedisStringDel(c *gin.Context) {
	var req types.StringDelRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: "Invalid request parameters"})
		return
	}

	// Validate required fields
	if req.Key == "" {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: "Key is required"})
		return
	}

	// Call service layer
	response, err := h.stringService.Del(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
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
func (h *RedisHandler) RedisStringExists(c *gin.Context) {
	var req types.StringExistsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: "Invalid request parameters"})
		return
	}

	// Validate required fields
	if req.Key == "" {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: "Key is required"})
		return
	}

	// Call service layer
	response, err := h.stringService.Exists(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
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
func (h *RedisHandler) RedisStringIncr(c *gin.Context) {
	var req types.StringIncrRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: "Invalid request parameters"})
		return
	}

	// Validate required fields
	if req.Key == "" {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: "Key is required"})
		return
	}

	// Call service layer
	response, err := h.stringService.Incr(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
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
func (h *RedisHandler) RedisStringDecr(c *gin.Context) {
	var req types.StringDecrRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: "Invalid request parameters"})
		return
	}

	// Validate required fields
	if req.Key == "" {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: "Key is required"})
		return
	}

	// Call service layer
	response, err := h.stringService.Decr(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
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
func (h *RedisHandler) RedisStringExpire(c *gin.Context) {
	var req types.StringExpireRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: "Invalid request parameters"})
		return
	}

	// Validate required fields
	if req.Key == "" {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: "Key is required"})
		return
	}
	if req.TTL <= 0 {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: "TTL must be greater than 0"})
		return
	}

	// Call service layer
	response, err := h.stringService.Expire(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}