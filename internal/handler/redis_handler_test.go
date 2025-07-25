
package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redismock/v8"
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

	handler := RedisStringGet(db)

	t.Run("Successful Get", func(t *testing.T) {
		body := `{"addr": "localhost:6379", "key": "my-key"}`
		req, _ := http.NewRequest("POST", "/api/v1/redis/string/get", strings.NewReader(body))
		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.JSONEq(t, `{"code":0,"data":{"value":"my-value"},"msg":"success"}`, rr.Body.String())
	})

	t.Run("Key Not Found", func(t *testing.T) {
		body := `{"addr": "localhost:6379", "key": "non-existent-key"}`
		req, _ := http.NewRequest("POST", "/api/v1/redis/string/get", strings.NewReader(body))
		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.JSONEq(t, `{"code":0,"data":{"value":null},"msg":"success"}`, rr.Body.String())
	})

	t.Run("Redis Error", func(t *testing.T) {
		body := `{"addr": "localhost:6379", "key": "error-key"}`
		req, _ := http.NewRequest("POST", "/api/v1/redis/string/get", strings.NewReader(body))
		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Contains(t, rr.Body.String(), "error connecting to redis")
	})

	t.Run("Invalid Request Body", func(t *testing.T) {
		body := `{"addr": "localhost:6379", "key": "my-key"` // Invalid JSON
		req, _ := http.NewRequest("POST", "/api/v1/redis/string/get", strings.NewReader(body))
		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})
}
