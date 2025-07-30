package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ct-zh/go-redis-proxy/internal/service"
	"github.com/ct-zh/go-redis-proxy/pkg/types"
)

// DynamicStringController 支持动态连接的字符串操作控制器
type DynamicStringController struct {
	stringService *service.DynamicRedisStringService
}

// NewDynamicStringController 创建新的动态字符串控制器
func NewDynamicStringController(stringService *service.DynamicRedisStringService) *DynamicStringController {
	return &DynamicStringController{
		stringService: stringService,
	}
}

// Get 处理GET请求
func (c *DynamicStringController) Get(ctx *gin.Context) {
	var req types.StringGetRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := c.stringService.Get(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

// Set 处理SET请求
func (c *DynamicStringController) Set(ctx *gin.Context) {
	var req types.StringSetRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := c.stringService.Set(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

// Del 处理DEL请求
func (c *DynamicStringController) Del(ctx *gin.Context) {
	var req types.StringDelRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := c.stringService.Del(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}