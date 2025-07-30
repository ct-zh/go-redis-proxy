package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ct-zh/go-redis-proxy/internal/service"
	"github.com/ct-zh/go-redis-proxy/pkg/types"
)

// SimpleTokenController 简化的Token管理控制器
type SimpleTokenController struct {
	tokenService *service.SimpleTokenService
}

// NewSimpleTokenController 创建简化的Token管理控制器
func NewSimpleTokenController(tokenService *service.SimpleTokenService) *SimpleTokenController {
	return &SimpleTokenController{
		tokenService: tokenService,
	}
}

// CreateToken 创建Token
// @Summary 创建Token
// @Description 根据Redis连接配置创建Token
// @Tags Token管理
// @Accept json
// @Produce json
// @Param request body types.ConnectRequest true "连接请求"
// @Success 200 {object} types.ConnectResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/token/create [post]
func (c *SimpleTokenController) CreateToken(ctx *gin.Context) {
	var req types.ConnectRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request format",
			"details": err.Error(),
		})
		return
	}

	response, err := c.tokenService.CreateToken(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create token",
			"details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// RefreshToken 刷新Token
// @Summary 刷新Token
// @Description 刷新现有Token的有效期
// @Tags Token管理
// @Accept json
// @Produce json
// @Param request body types.TokenRefreshRequest true "刷新请求"
// @Success 200 {object} types.ConnectResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/token/refresh [post]
func (c *SimpleTokenController) RefreshToken(ctx *gin.Context) {
	var req types.TokenRefreshRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request format",
			"details": err.Error(),
		})
		return
	}

	response, err := c.tokenService.RefreshToken(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to refresh token",
			"details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// ValidateToken 验证Token
// @Summary 验证Token
// @Description 验证Token是否有效
// @Tags Token管理
// @Accept json
// @Produce json
// @Param token query string true "Token"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/token/validate [get]
func (c *SimpleTokenController) ValidateToken(ctx *gin.Context) {
	token := ctx.Query("token")
	if token == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Token is required",
		})
		return
	}

	valid, err := c.tokenService.ValidateToken(ctx.Request.Context(), token)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to validate token",
			"details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"valid": valid,
		"token": token,
	})
}

// DeleteToken 删除Token
// @Summary 删除Token
// @Description 删除指定的Token
// @Tags Token管理
// @Accept json
// @Produce json
// @Param request body types.DisconnectRequest true "断开连接请求"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/token/delete [post]
func (c *SimpleTokenController) DeleteToken(ctx *gin.Context) {
	var req types.DisconnectRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request format",
			"details": err.Error(),
		})
		return
	}

	err := c.tokenService.DeleteToken(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to delete token",
			"details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Token deleted successfully",
	})
}

// GetStats 获取连接统计
// @Summary 获取连接统计
// @Description 获取当前连接池和Token的统计信息
// @Tags Token管理
// @Produce json
// @Success 200 {object} types.ConnectionStats
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/token/stats [get]
func (c *SimpleTokenController) GetStats(ctx *gin.Context) {
	stats, err := c.tokenService.GetConnectionStats(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to get stats",
			"details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, stats)
}

// HealthCheck 健康检查
// @Summary 健康检查
// @Description 检查Token管理服务的健康状态
// @Tags Token管理
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/token/health [get]
func (c *SimpleTokenController) HealthCheck(ctx *gin.Context) {
	err := c.tokenService.HealthCheck(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": "unhealthy",
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "healthy",
		"message": "Token service is running normally",
	})
}