package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"slash-admin/app/admin/cmd/api/internal/config"
	"slash-admin/app/admin/cmd/api/internal/middleware"
	"slash-admin/app/admin/ent"
)

type ServiceContext struct {
	RedisClient *redis.Redis
	EntClient   *ent.Client
	Config      config.Config
	Authority   rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	var (
		redisClient *redis.Redis
		entClient   *ent.Client
		err         error
	)

	if redisClient, err = c.Redis.NewRedis(); err != nil {
		panic(err)
	}
	if entClient, err = c.Database.NewDatabase(c.Redis); err != nil {
		panic(err)
	}

	return &ServiceContext{
		RedisClient: redisClient,
		EntClient:   entClient,
		Config:      c,
		Authority:   middleware.NewAuthorityMiddleware(nil, redisClient).Handle,
	}
}
