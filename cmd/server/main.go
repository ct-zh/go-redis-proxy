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
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/ct-zh/go-redis-proxy/api/swagger"
	"github.com/ct-zh/go-redis-proxy/internal/config"
	"github.com/ct-zh/go-redis-proxy/internal/container"
	"github.com/ct-zh/go-redis-proxy/internal/router"
	"github.com/ct-zh/go-redis-proxy/pkg/errors"
	"github.com/ct-zh/go-redis-proxy/pkg/logger"
)

func main() {
	// Validate error registry before starting server
	if err := errors.ValidateRegistry(); err != nil {
		log.Fatalf("Error code validation failed: %v", err)
	}
	log.Println("Error registry validation passed")

	// Load configuration
	cfg := config.Load()

	// Initialize logger
	loggerConfig := logger.LoggerConfig{
		Level:    cfg.Log.Level,
		Dir:      cfg.Log.Dir,
		MaxSize:  cfg.Log.MaxSize,
		MaxAge:   cfg.Log.MaxAge,
		Compress: cfg.Log.Compress,
	}
	if err := logger.Init(loggerConfig); err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	logger.Info("Logger initialized successfully", nil)

	// Initialize dependency injection container
	appContainer, cleanup, err := container.InitializeContainer()
	if err != nil {
		log.Fatalf("Failed to initialize container: %v", err)
	}
	defer cleanup()

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
	go func() {
		if err := engine.Run(addr); err != nil {
			log.Printf("Server failed to start: %v", err)
		}
	}()

	// Wait for shutdown signal
	<-quit
	fmt.Println("\nShutting down server...")

	fmt.Println("Server stopped")
}
