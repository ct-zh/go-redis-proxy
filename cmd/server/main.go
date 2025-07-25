package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/ct-zh/go-redis-proxy/internal/config"
	"github.com/ct-zh/go-redis-proxy/internal/router"
)

func main() {
	// 加载配置
	cfg := config.Load()

	// 创建Gin引擎
	engine := gin.Default()

	// 设置路由
	router.Setup(engine)

	// 启动服务器
	addr := cfg.GetServerAddr()
	fmt.Printf("Server starting on %s\n", addr)
	fmt.Println("Available endpoints:")
	fmt.Println("  GET /ping   - Ping endpoint")
	fmt.Println("  GET /health - Health check")
	fmt.Println("  POST /api/v1/redis/string/get - Redis string get")

	if err := engine.Run(addr); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
