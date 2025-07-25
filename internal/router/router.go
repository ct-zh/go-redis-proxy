package router

import (
	"net/http"

	"github.com/ct-zh/go-redis-proxy/internal/handler"
)

func Setup() *http.ServeMux {
	mux := http.NewServeMux()
	
	// 注册ping路由
	mux.HandleFunc("/ping", handler.Ping)
	mux.HandleFunc("/health", handler.Ping) // 健康检查使用相同处理器

	// Redis string
	mux.HandleFunc("/api/v1/redis/string/get", handler.RedisStringGet(nil))
	
	return mux
}