package svc

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"slash-admin/app/admin/cmd/api/internal/config"
	"slash-admin/app/admin/cmd/api/internal/converter"
	"slash-admin/app/admin/cmd/api/internal/converter/generated"
	"slash-admin/app/admin/cmd/api/internal/middleware"
	"slash-admin/app/admin/ent"
)

type ServiceContext struct {
	Redis     *redis.Redis
	EntClient *ent.Client
	Config    config.Config
	Authority rest.Middleware
	Converter converter.Converter
}

func NewServiceContext(c config.Config) *ServiceContext {
	var (
		redisClient *redis.Redis
		entClient   *ent.Client
		err         error
	)
	if redisClient, err = c.Redis.NewRedis(); err != nil {
		logx.Errorf("初始化redis失败: ", err)
		panic(err)
	}
	if entClient, err = c.Database.NewDatabase(c.Redis); err != nil {
		logx.Errorf("初始化ent失败: ", err)
		panic(err)
	}

	return &ServiceContext{
		Redis:     redisClient,
		EntClient: entClient,
		Config:    c,
		Authority: middleware.NewAuthorityMiddleware(nil, redisClient).Handle,
		Converter: &generated.ConverterImpl{},
	}
}
