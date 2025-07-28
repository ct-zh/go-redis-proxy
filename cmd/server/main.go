// @title Go Redis Proxy API
// @version 1.0.0
// @description Redis HTTP代理服务API文档
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api/v1

// @schemes http https

package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"

	"github.com/ct-zh/go-redis-proxy/internal/config"
	"github.com/ct-zh/go-redis-proxy/internal/router"
	_ "github.com/ct-zh/go-redis-proxy/api/swagger"
)

func main() {
	// 加载配置
	cfg := config.Load()

	// 创建Gin引擎
	engine := gin.Default()

	// 设置Swagger路由
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 设置API路由
	router.Setup(engine)

	// 启动服务器
	addr := cfg.GetServerAddr()
	fmt.Printf("Server starting on %s\n", addr)
	fmt.Println("Available endpoints:")
	fmt.Println("  GET /ping   - Ping endpoint")
	fmt.Println("  GET /health - Health check")
	fmt.Println("  POST /api/v1/redis/string/get - Redis string get")
	fmt.Println("  POST /api/v1/redis/string/set - Redis string set")
	fmt.Println("  POST /api/v1/redis/string/del - Redis string delete")
	fmt.Println("  POST /api/v1/redis/string/exists - Redis string exists")
	fmt.Println("  POST /api/v1/redis/string/incr - Redis string increment")
	fmt.Println("  POST /api/v1/redis/string/decr - Redis string decrement")
	fmt.Println("  POST /api/v1/redis/string/expire - Redis string expire")
	fmt.Println("  POST /api/v1/redis/list/lpush - Redis list left push")
	fmt.Println("  POST /api/v1/redis/list/rpush - Redis list right push")
	fmt.Println("  POST /api/v1/redis/list/lpop - Redis list left pop")
	fmt.Println("  POST /api/v1/redis/list/rpop - Redis list right pop")
	fmt.Println("  POST /api/v1/redis/list/lrem - Redis list remove")
	fmt.Println("  POST /api/v1/redis/list/lindex - Redis list index")
	fmt.Println("  POST /api/v1/redis/list/lrange - Redis list range")
	fmt.Println("  POST /api/v1/redis/list/llen - Redis list length")
	fmt.Println("  POST /api/v1/redis/list/ltrim - Redis list trim")
	fmt.Println("  GET /swagger/index.html - API Documentation")

	if err := engine.Run(addr); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
