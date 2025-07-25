package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ct-zh/go-redis-proxy/internal/config"
	"github.com/ct-zh/go-redis-proxy/internal/router"
)

func main() {
	// 加载配置
	cfg := config.Load()
	
	// 设置路由
	handler := router.Setup()
	
	// 启动服务器
	addr := cfg.GetServerAddr()
	fmt.Printf("Server starting on %s\n", addr)
	fmt.Println("Available endpoints:")
	fmt.Println("  GET /ping   - Ping endpoint")
	fmt.Println("  GET /health - Health check")
	
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}