package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"message":   "pong",
			"timestamp": time.Now().Format(time.RFC3339),
		},
	})
}
