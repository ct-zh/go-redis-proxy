package handler

import (
	"net/http"

	"github.com/ct-zh/go-redis-proxy/internal/service"
	"github.com/ct-zh/go-redis-proxy/pkg/types"
	"github.com/gin-gonic/gin"
)

// RedisSetHandler handles set-related requests
type RedisSetHandler struct {
	svc service.RedisSetService
}

// NewRedisSetHandler creates a new RedisSetHandler
func NewRedisSetHandler(svc service.RedisSetService) *RedisSetHandler {
	return &RedisSetHandler{svc: svc}
}

// SAdd handles the SADD command
func (h *RedisSetHandler) SAdd(c *gin.Context) {
	var req types.RedisSAddRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	count, err := h.svc.SAdd(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"count": count})
}

// SRem handles the SREM command
func (h *RedisSetHandler) SRem(c *gin.Context) {
	var req types.RedisSRemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	count, err := h.svc.SRem(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"count": count})
}

// SIsMember handles the SISMEMBER command
func (h *RedisSetHandler) SIsMember(c *gin.Context) {
	var req types.RedisSIsMemberRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exists, err := h.svc.SIsMember(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"exists": exists})
}

// SMembers handles the SMEMBERS command
func (h *RedisSetHandler) SMembers(c *gin.Context) {
	var req types.RedisSMembersRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	members, err := h.svc.SMembers(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"members": members})
}

// SCard handles the SCARD command
func (h *RedisSetHandler) SCard(c *gin.Context) {
	var req types.RedisSCardRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	count, err := h.svc.SCard(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"count": count})
}
