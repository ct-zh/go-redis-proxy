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
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"

	"github.com/ct-zh/go-redis-proxy/internal/config"
	"github.com/ct-zh/go-redis-proxy/internal/container"
	"github.com/ct-zh/go-redis-proxy/internal/router"
	"github.com/ct-zh/go-redis-proxy/pkg/errors"
	_ "github.com/ct-zh/go-redis-proxy/api/swagger"
)

func main() {
	// Validate error registry before starting server
	if err := errors.ValidateRegistry(); err != nil {
		log.Fatalf("Error code validation failed: %v", err)
	}
	log.Println("Error registry validation passed")

	// Load configuration
	cfg := config.Load()

	// Initialize dependency injection container
	appContainer := container.NewContainer()
	defer func() {
		if err := appContainer.Cleanup(); err != nil {
			log.Printf("Error during cleanup: %v", err)
		}
	}()

	// Create Gin engine
	engine := gin.Default()

	// Setup Swagger routes
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Setup API routes with dependency injection
	router.SetupWithContainer(engine, appContainer)

	// Setup graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Start server in a goroutine
	addr := cfg.GetServerAddr()
	fmt.Printf("Server starting on %s\n", addr)
	fmt.Println("=== Go Redis Proxy Server ===")
	fmt.Println("Architecture: Layered (Handler -> Service -> DAO)")
	fmt.Println("")
	fmt.Println("Available endpoints:")
	fmt.Println("  GET  /ping                               - Ping endpoint")
	fmt.Println("  GET  /health                             - Health check")
	fmt.Println("")
	fmt.Println("=== Redis String Operations ===")
	fmt.Println("  POST /api/v1/redis/string/get           - Redis string get")
	fmt.Println("  POST /api/v1/redis/string/set           - Redis string set")
	fmt.Println("  POST /api/v1/redis/string/del           - Redis string delete")
	fmt.Println("  POST /api/v1/redis/string/exists        - Redis string exists")
	fmt.Println("  POST /api/v1/redis/string/incr          - Redis string increment")
	fmt.Println("  POST /api/v1/redis/string/decr          - Redis string decrement")
	fmt.Println("  POST /api/v1/redis/string/expire        - Redis string expire")
	fmt.Println("")
	fmt.Println("=== Redis List Operations ===")
	fmt.Println("  POST /api/v1/redis/list/lpush           - Redis list left push")
	fmt.Println("  POST /api/v1/redis/list/rpush           - Redis list right push")
	fmt.Println("  POST /api/v1/redis/list/lpop            - Redis list left pop")
	fmt.Println("  POST /api/v1/redis/list/rpop            - Redis list right pop")
	fmt.Println("  POST /api/v1/redis/list/lrem            - Redis list remove")
	fmt.Println("  POST /api/v1/redis/list/lindex          - Redis list index")
	fmt.Println("  POST /api/v1/redis/list/lrange          - Redis list range")
	fmt.Println("  POST /api/v1/redis/list/llen            - Redis list length")
	fmt.Println("  POST /api/v1/redis/list/ltrim           - Redis list trim")
	fmt.Println("")
	fmt.Println("=== Documentation ===")
	fmt.Println("  GET  /swagger/index.html                - API Documentation")
	fmt.Println("")

	go func() {
		if err := engine.Run(addr); err != nil {
			log.Printf("Server failed to start: %v", err)
		}
	}()

	// Wait for shutdown signal
	<-quit
	fmt.Println("\nShutting down server...")

	// Perform graceful shutdown
	if err := appContainer.Cleanup(); err != nil {
		log.Printf("Error during cleanup: %v", err)
	}

	fmt.Println("Server stopped")
}