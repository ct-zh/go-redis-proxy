package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Ping godoc
// @Summary Ping endpoint
// @Description Ping服务，用于健康检查
// @Tags Health
// @Accept json
// @Produce json
// @Success 200 {object} types.PingResponse "成功响应"
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"message":   "pong",
			"timestamp": time.Now().Format(time.RFC3339),
		},
	})
}
