package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ct-zh/go-redis-proxy/pkg/types"
)

// List operations

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
func (h *RedisHandler) RedisListLPush(c *gin.Context) {
	var req types.ListLPushRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: "Invalid request parameters"})
		return
	}

	// Validate required fields
	if req.Key == "" {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: "Key is required"})
		return
	}
	if len(req.Values) == 0 {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: "Values are required"})
		return
	}

	// Call service layer
	response, err := h.listService.LPush(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
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
func (h *RedisHandler) RedisListRPush(c *gin.Context) {
	var req types.ListRPushRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: "Invalid request parameters"})
		return
	}

	// Validate required fields
	if req.Key == "" {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: "Key is required"})
		return
	}
	if len(req.Values) == 0 {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: "Values are required"})
		return
	}

	// Call service layer
	response, err := h.listService.RPush(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
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
func (h *RedisHandler) RedisListLPop(c *gin.Context) {
	var req types.ListLPopRequest
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
	response, err := h.listService.LPop(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
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
func (h *RedisHandler) RedisListRPop(c *gin.Context) {
	var req types.ListRPopRequest
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
	response, err := h.listService.RPop(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
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
func (h *RedisHandler) RedisListLRem(c *gin.Context) {
	var req types.ListLRemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: "Invalid request parameters"})
		return
	}

	// Validate required fields
	if req.Key == "" {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: "Key is required"})
		return
	}
	if req.Value == "" {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: "Value is required"})
		return
	}

	// Call service layer
	response, err := h.listService.LRem(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
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
func (h *RedisHandler) RedisListLIndex(c *gin.Context) {
	var req types.ListLIndexRequest
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
	response, err := h.listService.LIndex(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
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
func (h *RedisHandler) RedisListLRange(c *gin.Context) {
	var req types.ListLRangeRequest
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
	response, err := h.listService.LRange(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
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
func (h *RedisHandler) RedisListLLen(c *gin.Context) {
	var req types.ListLLenRequest
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
	response, err := h.listService.LLen(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
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
func (h *RedisHandler) RedisListLTrim(c *gin.Context) {
	var req types.ListLTrimRequest
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
	response, err := h.listService.LTrim(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}