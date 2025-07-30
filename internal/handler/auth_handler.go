package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/ct-zh/go-redis-proxy/internal/service"
	"github.com/ct-zh/go-redis-proxy/pkg/response"
	"github.com/ct-zh/go-redis-proxy/pkg/types"
)

// AuthHandler 处理认证相关的HTTP请求
type AuthHandler struct {
	authService       service.AuthService
	tokenRedisService service.TokenRedisService
}

// NewAuthHandler 创建新的认证处理器
func NewAuthHandler(authService service.AuthService, tokenRedisService service.TokenRedisService) *AuthHandler {
	return &AuthHandler{
		authService:       authService,
		tokenRedisService: tokenRedisService,
	}
}

// Connect godoc
// @Summary 建立Redis连接
// @Description 建立Redis连接并返回访问token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body types.ConnectRequest true "连接请求参数"
// @Success 200 {object} response.BaseResponse{data=types.ConnectResponse} "连接成功"
// @Failure 400 {object} response.BaseResponse "请求参数错误"
// @Failure 500 {object} response.BaseResponse "服务器内部错误"
// @Router /auth/connect [post]
func (h *AuthHandler) Connect(c *gin.Context) {
	var req types.ConnectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request parameters", err)
		return
	}

	data, err := h.authService.Connect(c.Request.Context(), &req)
	if err != nil {
		response.InternalServerError(c, "Failed to establish connection", err)
		return
	}

	response.SuccessWithMessage(c, "Connection established successfully", data)
}

// RefreshToken godoc
// @Summary 刷新Token
// @Description 刷新现有的token，延长其有效期
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body types.TokenRefreshRequest true "请求参数"
// @Success 200 {object} response.BaseResponse "成功响应"
// @Failure 400 {object} response.BaseResponse "请求参数错误"
// @Failure 500 {object} response.BaseResponse "服务器内部错误"
// @Router /auth/refresh [post]
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var req types.TokenRefreshRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request parameters", err)
		return
	}

	// Validate required fields
	if req.Token == "" {
		response.BadRequest(c, "Token is required", nil)
		return
	}

	// Call service layer
	resp, err := h.authService.RefreshToken(c.Request.Context(), &req)
	if err != nil {
		response.InternalServerError(c, "Failed to refresh token", err)
		return
	}

	response.SuccessWithMessage(c, "Token refreshed successfully", resp)
}

// Disconnect godoc
// @Summary 断开连接
// @Description 断开指定token的Redis连接，释放资源
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body types.DisconnectRequest true "请求参数"
// @Success 200 {object} response.BaseResponse "成功响应"
// @Failure 400 {object} response.BaseResponse "请求参数错误"
// @Failure 500 {object} response.BaseResponse "服务器内部错误"
// @Router /auth/disconnect [post]
func (h *AuthHandler) Disconnect(c *gin.Context) {
	var req types.DisconnectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request parameters", err)
		return
	}

	// Validate required fields
	if req.Token == "" {
		response.BadRequest(c, "Token is required", nil)
		return
	}

	// Call service layer
	err := h.authService.Disconnect(c.Request.Context(), &req)
	if err != nil {
		response.InternalServerError(c, "Failed to disconnect", err)
		return
	}

	response.SuccessWithMessage(c, "Disconnected successfully", nil)
}

// GetStats godoc
// @Summary 获取连接统计信息
// @Description 获取当前连接池的统计信息
// @Tags Authentication
// @Accept json
// @Produce json
// @Success 200 {object} response.BaseResponse "成功响应"
// @Failure 500 {object} response.BaseResponse "服务器内部错误"
// @Router /auth/stats [get]
func (h *AuthHandler) GetStats(c *gin.Context) {
	// Call service layer
	stats, err := h.authService.GetStats(c.Request.Context())
	if err != nil {
		response.InternalServerError(c, "Failed to get stats", err)
		return
	}

	response.SuccessWithMessage(c, "Stats retrieved successfully", stats)
}

// Token-based Redis operations

// TokenStringGet godoc
// @Summary Token方式Redis字符串GET操作
// @Description 使用token方式根据指定的key获取Redis中存储的字符串值
// @Tags Token Redis Operations
// @Accept json
// @Produce json
// @Param request body types.TokenStringGetRequest true "请求参数"
// @Success 200 {object} response.BaseResponse "成功响应"
// @Failure 400 {object} response.BaseResponse "请求参数错误"
// @Failure 500 {object} response.BaseResponse "服务器内部错误"
// @Router /token/redis/string/get [post]
func (h *AuthHandler) TokenStringGet(c *gin.Context) {
	var req types.TokenStringGetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request parameters", err)
		return
	}

	// Validate required fields
	if req.Token == "" || req.Key == "" {
		response.BadRequest(c, "Token and key are required", nil)
		return
	}

	// Call service layer
	data, err := h.tokenRedisService.StringGet(c.Request.Context(), &req)
	if err != nil {
		response.InternalServerError(c, "Failed to get string value", err)
		return
	}

	response.SuccessWithMessage(c, "String value retrieved successfully", data)
}

// TokenStringSet godoc
// @Summary Token方式Redis字符串SET操作
// @Description 使用token方式设置指定key的字符串值，支持TTL过期时间
// @Tags Token Redis Operations
// @Accept json
// @Produce json
// @Param request body types.TokenStringSetRequest true "请求参数"
// @Success 200 {object} response.BaseResponse "成功响应"
// @Failure 400 {object} response.BaseResponse "请求参数错误"
// @Failure 500 {object} response.BaseResponse "服务器内部错误"
// @Router /token/redis/string/set [post]
func (h *AuthHandler) TokenStringSet(c *gin.Context) {
	var req types.TokenStringSetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request parameters", err)
		return
	}

	// Validate required fields
	if req.Token == "" || req.Key == "" || req.Value == "" {
		response.BadRequest(c, "Token, key and value are required", nil)
		return
	}

	// Call service layer
	data, err := h.tokenRedisService.StringSet(c.Request.Context(), &req)
	if err != nil {
		response.InternalServerError(c, "Failed to set string value", err)
		return
	}

	response.SuccessWithMessage(c, "String value set successfully", data)
}

// TokenStringDel godoc
// @Summary Token方式Redis字符串DEL操作
// @Description 使用token方式删除指定的key
// @Tags Token Redis Operations
// @Accept json
// @Produce json
// @Param request body types.TokenStringDelRequest true "请求参数"
// @Success 200 {object} response.BaseResponse "成功响应"
// @Failure 400 {object} response.BaseResponse "请求参数错误"
// @Failure 500 {object} response.BaseResponse "服务器内部错误"
// @Router /token/redis/string/del [post]
func (h *AuthHandler) TokenStringDel(c *gin.Context) {
	var req types.TokenStringDelRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request parameters", err)
		return
	}

	// Validate required fields
	if req.Token == "" || req.Key == "" {
		response.BadRequest(c, "Token and key are required", nil)
		return
	}

	// Call service layer
	data, err := h.tokenRedisService.StringDel(c.Request.Context(), &req)
	if err != nil {
		response.InternalServerError(c, "Failed to delete string value", err)
		return
	}

	response.SuccessWithMessage(c, "String value deleted successfully", data)
}

// TokenHashHSet godoc
// @Summary Token方式Redis哈希HSET操作
// @Description 使用token方式设置哈希表中指定字段的值
// @Tags Token Redis Operations
// @Accept json
// @Produce json
// @Param request body types.TokenHashHSetRequest true "请求参数"
// @Success 200 {object} response.BaseResponse "成功响应"
// @Failure 400 {object} response.BaseResponse "请求参数错误"
// @Failure 500 {object} response.BaseResponse "服务器内部错误"
// @Router /token/redis/hash/hset [post]
func (h *AuthHandler) TokenHashHSet(c *gin.Context) {
	var req types.TokenHashHSetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request parameters", err)
		return
	}

	// Validate required fields
	if req.Token == "" || req.Key == "" || req.Field == "" || req.Value == "" {
		response.BadRequest(c, "Token, key, field and value are required", nil)
		return
	}

	// Call service layer
	data, err := h.tokenRedisService.HashHSet(c.Request.Context(), &req)
	if err != nil {
		response.InternalServerError(c, "Failed to set hash field", err)
		return
	}

	response.SuccessWithMessage(c, "Hash field set successfully", data)
}

// TokenHashHGet godoc
// @Summary Token方式Redis哈希HGET操作
// @Description 使用token方式获取哈希表中指定字段的值
// @Tags Token Redis Operations
// @Accept json
// @Produce json
// @Param request body types.TokenHashHGetRequest true "请求参数"
// @Success 200 {object} response.BaseResponse "成功响应"
// @Failure 400 {object} response.BaseResponse "请求参数错误"
// @Failure 500 {object} response.BaseResponse "服务器内部错误"
// @Router /token/redis/hash/hget [post]
func (h *AuthHandler) TokenHashHGet(c *gin.Context) {
	var req types.TokenHashHGetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request parameters", err)
		return
	}

	// Validate required fields
	if req.Token == "" || req.Key == "" || req.Field == "" {
		response.BadRequest(c, "Token, key and field are required", nil)
		return
	}

	// Call service layer
	data, err := h.tokenRedisService.HashHGet(c.Request.Context(), &req)
	if err != nil {
		response.InternalServerError(c, "Failed to get hash field", err)
		return
	}

	response.SuccessWithMessage(c, "Hash field retrieved successfully", data)
}