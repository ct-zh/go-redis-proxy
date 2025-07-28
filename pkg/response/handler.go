package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ct-zh/go-redis-proxy/pkg/errors"
)

// Result 标准响应结构
type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// Handler 封装的响应处理器
type Handler struct {
	*gin.Context
}

// NewHandler 创建响应处理器
func NewHandler(c *gin.Context) *Handler {
	return &Handler{Context: c}
}

// Response 统一响应方法
// data: 业务数据，自动包装到response.data字段
// err: 错误信息，支持BusinessError或标准error
func (h *Handler) Response(data interface{}, err error) {
	if err != nil {
		h.handleError(data, err)
		return
	}
	
	// 成功响应
	h.JSON(http.StatusOK, Result{
		Code: 200,
		Msg:  "Success",
		Data: data,
	})
}

// Success 成功响应的便捷方法
func (h *Handler) Success(data interface{}) {
	h.Response(data, nil)
}

// Error 错误响应的便捷方法
func (h *Handler) Error(err error) {
	h.Response(nil, err)
}

// handleError 处理错误响应
func (h *Handler) handleError(data interface{}, err error) {
	// 尝试转换为BusinessError
	if bizErr, ok := err.(errors.BusinessError); ok {
		h.JSON(http.StatusOK, Result{
			Code: bizErr.Code(),
			Msg:  bizErr.Message(),
			Data: data,
		})
		return
	}
	
	// 标准error，返回500
	h.JSON(http.StatusOK, Result{
		Code: errors.CodeInternalError,
		Msg:  err.Error(),
		Data: data,
	})
}

// BadRequest 400错误响应
func (h *Handler) BadRequest(message string) {
	h.JSON(http.StatusBadRequest, Result{
		Code: errors.CodeInvalidParams,
		Msg:  message,
	})
}

// Unauthorized 401错误响应
func (h *Handler) Unauthorized(message string) {
	h.JSON(http.StatusUnauthorized, Result{
		Code: errors.CodeUnauthorized,
		Msg:  message,
	})
}

// InternalError 500错误响应
func (h *Handler) InternalError(message string) {
	h.JSON(http.StatusInternalServerError, Result{
		Code: errors.CodeInternalError,
		Msg:  message,
	})
}

// 全局便捷函数

// Respond 全局响应函数
func Respond(c *gin.Context, data interface{}, err error) {
	handler := NewHandler(c)
	handler.Response(data, err)
}
