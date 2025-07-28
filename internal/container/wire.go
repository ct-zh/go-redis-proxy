//go:build wireinject
// +build wireinject

package container

import (
	"github.com/ct-zh/go-redis-proxy/internal/dao"
	"github.com/ct-zh/go-redis-proxy/internal/handler"
	"github.com/ct-zh/go-redis-proxy/internal/service"
	"github.com/google/wire"
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

// buildProvider
var buildProvider = wire.NewSet(
	wire.Bind(new(dao.RedisDAO), new(*dao.RedisDAOImpl)),
	dao.NewRedisDAO,

	wire.Bind(new(service.RedisStringService), new(*service.RedisStringServiceImpl)),
	service.NewRedisStringService,
	wire.Bind(new(service.RedisListService), new(*service.RedisListServiceImpl)),
	service.NewRedisListService,
	wire.Bind(new(service.RedisSetService), new(*service.RedisSetServiceImpl)),
	service.NewRedisSetService,

	handler.NewRedisHandler,
	handler.NewRedisListHandler,
	handler.NewRedisSetHandler,

	wire.Struct(new(Container), "*"),
)

func InitializeContainer() (*Container, func(), error) {
	wire.Build(buildProvider)
	return nil, nil, nil
}

