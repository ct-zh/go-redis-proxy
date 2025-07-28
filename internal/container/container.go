package container

import (
	"github.com/ct-zh/go-redis-proxy/internal/dao"
	"github.com/ct-zh/go-redis-proxy/internal/handler"
	"github.com/ct-zh/go-redis-proxy/internal/service"
)

// Container holds all application dependencies
type Container struct {
	// DAO layer
	RedisDAO dao.RedisDAO

	// Service layer
	StringService service.RedisStringService
	ListService   service.RedisListService
	SetService    service.RedisSetService

	// Handler layer
	RedisHandler     *handler.RedisHandler
	RedisListHandler *handler.RedisListHandler
	RedisSetHandler  *handler.RedisSetHandler
}

// NewContainer creates and initializes a new container with all dependencies
func NewContainer() *Container {
	container := &Container{}

	// Initialize DAO layer
	container.RedisDAO = dao.NewRedisDAO()

	// Initialize Service layer with DAO dependencies
	container.StringService = service.NewRedisStringService(container.RedisDAO)
	container.ListService = service.NewRedisListService(container.RedisDAO)
	container.SetService = service.NewRedisSetService(container.RedisDAO)

	// Initialize Handler layer with Service dependencies
	container.RedisHandler = handler.NewRedisHandler(
		container.StringService,
		container.ListService,
	)
	container.RedisListHandler = handler.NewRedisListHandler(
		container.ListService,
	)

	container.RedisSetHandler = handler.NewRedisSetHandler(
		container.SetService,
	)

	return container
}

// Cleanup performs cleanup operations for all components
func (c *Container) Cleanup() error {
	// Close DAO connections
	if c.RedisDAO != nil {
		if err := c.RedisDAO.Close(); err != nil {
			return err
		}
	}
	return nil
}