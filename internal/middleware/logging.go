package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/ct-zh/go-redis-proxy/pkg/logger"
)

// responseWriter 包装gin的ResponseWriter以捕获响应数据
type responseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w responseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// LoggingMiddleware 日志中间件
func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 记录开始时间
		startTime := time.Now()

		// 读取请求体
		var requestBody []byte
		if c.Request.Body != nil {
			requestBody, _ = io.ReadAll(c.Request.Body)
			// 重新设置请求体，以便后续处理器可以读取
			c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		}

		// 包装ResponseWriter以捕获响应
		responseBuffer := &bytes.Buffer{}
		writer := &responseWriter{
			ResponseWriter: c.Writer,
			body:           responseBuffer,
		}
		c.Writer = writer

		// 处理请求
		c.Next()

		// 计算请求耗时
		duration := time.Since(startTime)

		// 准备日志字段
		fields := logrus.Fields{
			"method":        c.Request.Method,
			"path":          c.Request.URL.Path,
			"query":         c.Request.URL.RawQuery,
			"status_code":   c.Writer.Status(),
			"client_ip":     c.ClientIP(),
			"user_agent":    c.Request.UserAgent(),
			"duration_ms":   duration.Milliseconds(),
			"request_size":  len(requestBody),
			"response_size": responseBuffer.Len(),
		}

		// 添加请求体（如果是JSON格式且不为空）
		if len(requestBody) > 0 && isJSONContent(c.Request.Header.Get("Content-Type")) {
			var requestJSON interface{}
			if err := json.Unmarshal(requestBody, &requestJSON); err == nil {
				fields["request_body"] = requestJSON
			} else {
				fields["request_body"] = string(requestBody)
			}
		}

		// 添加响应体（如果是JSON格式且不为空）
		responseBody := responseBuffer.Bytes()
		if len(responseBody) > 0 && isJSONContent(c.Writer.Header().Get("Content-Type")) {
			var responseJSON interface{}
			if err := json.Unmarshal(responseBody, &responseJSON); err == nil {
				fields["response_body"] = responseJSON
			} else {
				fields["response_body"] = string(responseBody)
			}
		}

		// 根据状态码决定日志级别
		message := "HTTP Request"
		if c.Writer.Status() >= 500 {
			// 服务器错误
			logger.Error(message, fields)
		} else if c.Writer.Status() >= 400 {
			// 客户端错误
			logger.Warn(message, fields)
		} else {
			// 正常请求
			logger.Access(message, fields)
		}

		// 记录详细的调试信息
		if logger.GetLogger() != nil {
			debugFields := logrus.Fields{
				"headers": c.Request.Header,
				"params":  c.Params,
			}
			logger.Debug("Request Details", debugFields)
		}
	}
}

// isJSONContent 检查内容类型是否为JSON
func isJSONContent(contentType string) bool {
	return contentType == "application/json" || 
		   contentType == "application/json; charset=utf-8" ||
		   contentType == "text/json"
}

// RequestIDMiddleware 为每个请求生成唯一ID的中间件
func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := generateRequestID()
		c.Set("request_id", requestID)
		c.Header("X-Request-ID", requestID)
		c.Next()
	}
}

// generateRequestID 生成请求ID
func generateRequestID() string {
	return time.Now().Format("20060102150405") + "-" + randomString(8)
}

// randomString 生成随机字符串
func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[time.Now().UnixNano()%int64(len(charset))]
	}
	return string(b)
}

// RecoveryMiddleware 恢复中间件，捕获panic并记录日志
func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				fields := logrus.Fields{
					"error":     err,
					"method":    c.Request.Method,
					"path":      c.Request.URL.Path,
					"client_ip": c.ClientIP(),
				}
				
				if requestID, exists := c.Get("request_id"); exists {
					fields["request_id"] = requestID
				}

				logger.Error("Panic recovered", fields)
				
				// 返回500错误
				c.JSON(500, gin.H{
					"code": 500,
					"msg":  "Internal Server Error",
					"data": nil,
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}