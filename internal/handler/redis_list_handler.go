package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/ct-zh/go-redis-proxy/internal/service"
	"github.com/ct-zh/go-redis-proxy/pkg/types"
	"github.com/ct-zh/go-redis-proxy/pkg/response"
)

// RedisListHandler handles HTTP requests for Redis list operations
type RedisListHandler struct {
	listService service.RedisListService
}

// NewRedisListHandler creates a new RedisListHandler instance
func NewRedisListHandler(listService service.RedisListService) *RedisListHandler {
	return &RedisListHandler{
		listService: listService,
	}
}

// RedisListLPush godoc
// @Summary Redis列表LPUSH操作
// @Description 将一个或多个值插入到列表头部
// @Tags Redis List Operations
// @Accept json
// @Produce json
// @Param request body types.ListLPushRequest true "请求参数"
// @Success 200 {object} response.BaseResponse "成功响应"
// @Failure 400 {object} response.BaseResponse "请求参数错误"
// @Failure 500 {object} response.BaseResponse "服务器内部错误"
// @Router /redis/list/lpush [post]
func (h *RedisListHandler) RedisListLPush(c *gin.Context) {
	var req types.ListLPushRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request parameters", err)
		return
	}

	// Validate required fields
	if req.Key == "" || len(req.Values) == 0 {
		response.BadRequest(c, "Key and values are required", nil)
		return
	}

	// Call service layer
	data, err := h.listService.LPush(c.Request.Context(), &req)
	response.JSON(c, data, err)
}

// RedisListRPush godoc
// @Summary Redis列表RPUSH操作
// @Description 将一个或多个值插入到列表尾部
// @Tags Redis List Operations
// @Accept json
// @Produce json
// @Param request body types.ListRPushRequest true "请求参数"
// @Success 200 {object} response.BaseResponse "成功响应"
// @Failure 400 {object} response.BaseResponse "请求参数错误"
// @Failure 500 {object} response.BaseResponse "服务器内部错误"
// @Router /redis/list/rpush [post]
func (h *RedisListHandler) RedisListRPush(c *gin.Context) {
	var req types.ListRPushRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request parameters", err)
		return
	}

	// Validate required fields
	if req.Key == "" || len(req.Values) == 0 {
		response.BadRequest(c, "Key and values are required", nil)
		return
	}

	// Call service layer
	data, err := h.listService.RPush(c.Request.Context(), &req)
	response.JSON(c, data, err)
}

// RedisListLPop godoc
// @Summary Redis列表LPOP操作
// @Description 移出并获取列表的第一个元素
// @Tags Redis List Operations
// @Accept json
// @Produce json
// @Param request body types.ListLPopRequest true "请求参数"
// @Success 200 {object} response.BaseResponse "成功响应"
// @Failure 400 {object} response.BaseResponse "请求参数错误"
// @Failure 500 {object} response.BaseResponse "服务器内部错误"
// @Router /redis/list/lpop [post]
func (h *RedisListHandler) RedisListLPop(c *gin.Context) {
	var req types.ListLPopRequest
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
	data, err := h.listService.LPop(c.Request.Context(), &req)
	response.JSON(c, data, err)
}

// RedisListRPop godoc
// @Summary Redis列表RPOP操作
// @Description 移出并获取列表的最后一个元素
// @Tags Redis List Operations
// @Accept json
// @Produce json
// @Param request body types.ListRPopRequest true "请求参数"
// @Success 200 {object} response.BaseResponse "成功响应"
// @Failure 400 {object} response.BaseResponse "请求参数错误"
// @Failure 500 {object} response.BaseResponse "服务器内部错误"
// @Router /redis/list/rpop [post]
func (h *RedisListHandler) RedisListRPop(c *gin.Context) {
	var req types.ListRPopRequest
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
	data, err := h.listService.RPop(c.Request.Context(), &req)
	response.JSON(c, data, err)
}

// RedisListLRem godoc
// @Summary Redis列表LREM操作
// @Description 根据参数count的值，移除列表中与参数value相等的元素
// @Tags Redis List Operations
// @Accept json
// @Produce json
// @Param request body types.ListLRemRequest true "请求参数"
// @Success 200 {object} response.BaseResponse "成功响应"
// @Failure 400 {object} response.BaseResponse "请求参数错误"
// @Failure 500 {object} response.BaseResponse "服务器内部错误"
// @Router /redis/list/lrem [post]
func (h *RedisListHandler) RedisListLRem(c *gin.Context) {
	var req types.ListLRemRequest
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
	data, err := h.listService.LRem(c.Request.Context(), &req)
	response.JSON(c, data, err)
}

// RedisListLIndex godoc
// @Summary Redis列表LINDEX操作
// @Description 通过索引获取列表中的元素
// @Tags Redis List Operations
// @Accept json
// @Produce json
// @Param request body types.ListLIndexRequest true "请求参数"
// @Success 200 {object} response.BaseResponse "成功响应"
// @Failure 400 {object} response.BaseResponse "请求参数错误"
// @Failure 500 {object} response.BaseResponse "服务器内部错误"
// @Router /redis/list/lindex [post]
func (h *RedisListHandler) RedisListLIndex(c *gin.Context) {
	var req types.ListLIndexRequest
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
	data, err := h.listService.LIndex(c.Request.Context(), &req)
	response.JSON(c, data, err)
}

// RedisListLRange godoc
// @Summary Redis列表LRANGE操作
// @Description 获取列表指定范围内的元素
// @Tags Redis List Operations
// @Accept json
// @Produce json
// @Param request body types.ListLRangeRequest true "请求参数"
// @Success 200 {object} response.BaseResponse "成功响应"
// @Failure 400 {object} response.BaseResponse "请求参数错误"
// @Failure 500 {object} response.BaseResponse "服务器内部错误"
// @Router /redis/list/lrange [post]
func (h *RedisListHandler) RedisListLRange(c *gin.Context) {
	var req types.ListLRangeRequest
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
	data, err := h.listService.LRange(c.Request.Context(), &req)
	response.JSON(c, data, err)
}

// RedisListLLen godoc
// @Summary Redis列表LLEN操作
// @Description 获取列表长度
// @Tags Redis List Operations
// @Accept json
// @Produce json
// @Param request body types.ListLLenRequest true "请求参数"
// @Success 200 {object} response.BaseResponse "成功响应"
// @Failure 400 {object} response.BaseResponse "请求参数错误"
// @Failure 500 {object} response.BaseResponse "服务器内部错误"
// @Router /redis/list/llen [post]
func (h *RedisListHandler) RedisListLLen(c *gin.Context) {
	var req types.ListLLenRequest
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
	data, err := h.listService.LLen(c.Request.Context(), &req)
	response.JSON(c, data, err)
}

// RedisListLTrim godoc
// @Summary Redis列表LTRIM操作
// @Description 对一个列表进行修剪，让列表只保留指定区间内的元素
// @Tags Redis List Operations
// @Accept json
// @Produce json
// @Param request body types.ListLTrimRequest true "请求参数"
// @Success 200 {object} response.BaseResponse "成功响应"
// @Failure 400 {object} response.BaseResponse "请求参数错误"
// @Failure 500 {object} response.BaseResponse "服务器内部错误"
// @Router /redis/list/ltrim [post]
func (h *RedisListHandler) RedisListLTrim(c *gin.Context) {
	var req types.ListLTrimRequest
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
	data, err := h.listService.LTrim(c.Request.Context(), &req)
	response.JSON(c, data, err)
}