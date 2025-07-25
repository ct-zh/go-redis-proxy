package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	redis "github.com/go-redis/redis/v8"
	redismock "github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"
)

func TestRedisStringGet(t *testing.T) {
	db, mock := redismock.NewClientMock()

	// Mock a successful GET
	mock.ExpectGet("my-key").SetVal("my-value")

	// Mock a key that doesn't exist
	mock.ExpectGet("non-existent-key").SetErr(redis.Nil)

	// Mock a generic error
	mock.ExpectGet("error-key").SetErr(redis.ErrClosed)

	gin.SetMode(gin.TestMode)

	t.Run("Successful Get", func(t *testing.T) {
		rr := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rr)
		body := `{"addr": "localhost:6379", "key": "my-key"}`
		c.Request, _ = http.NewRequest("POST", "/api/v1/redis/string/get", bytes.NewBufferString(body))
		c.Request.Header.Set("Content-Type", "application/json")

		RedisStringGet(db)(c)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.JSONEq(t, `{"code":0,"data":{"value":"my-value"},"msg":"success"}`, rr.Body.String())
	})

	t.Run("Key Not Found", func(t *testing.T) {
		rr := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rr)
		body := `{"addr": "localhost:6379", "key": "non-existent-key"}`
		c.Request, _ = http.NewRequest("POST", "/api/v1/redis/string/get", bytes.NewBufferString(body))
		c.Request.Header.Set("Content-Type", "application/json")

		RedisStringGet(db)(c)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.JSONEq(t, `{"code":0,"data":{"value":null},"msg":"success"}`, rr.Body.String())
	})

	t.Run("Redis Error", func(t *testing.T) {
		rr := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rr)
		body := `{"addr": "localhost:6379", "key": "error-key"}`
		c.Request, _ = http.NewRequest("POST", "/api/v1/redis/string/get", bytes.NewBufferString(body))
		c.Request.Header.Set("Content-Type", "application/json")

		RedisStringGet(db)(c)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Contains(t, rr.Body.String(), "error connecting to redis")
	})

	t.Run("Invalid Request Body", func(t *testing.T) {
		rr := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rr)
		body := `{"addr": "localhost:6379", "key": "my-key"` // Invalid JSON
		c.Request, _ = http.NewRequest("POST", "/api/v1/redis/string/get", bytes.NewBufferString(body))
		c.Request.Header.Set("Content-Type", "application/json")

		RedisStringGet(db)(c)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})
}
