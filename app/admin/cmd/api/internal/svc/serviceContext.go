package svc

import (
	"context"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/casbin/casbin/v2"
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
	Casbin    *casbin.SyncedEnforcer
}

func NewServiceContext(c config.Config) *ServiceContext {
	var (
		redisClient    *redis.Redis
		entClient      *ent.Client
		casbinEnforcer *casbin.SyncedEnforcer
		err            error
	)
	if redisClient, err = c.Redis.NewRedis(); err != nil {
		logx.Errorf("初始化redis失败: ", err)
		panic(err)
	}
	if entClient, err = c.Database.NewDatabase(c.Redis); err != nil {
		logx.Errorf("初始化ent失败: ", err)
		panic(err)
	}

	if c.Database.AutoMigrate {
		if err = entClient.Schema.Create(context.Background(), schema.WithForeignKeys(false)); err != nil {
			logx.Errorf("初始化ent schema失败: ", err)
			panic(err)
		}
	}

	if casbinEnforcer, err = c.Casbin.NewCasbin(entClient); err != nil {
		logx.Errorf("初始化casbin失败: ", err)
		panic(err)
	}

	return &ServiceContext{
		Redis:     redisClient,
		EntClient: entClient,
		Config:    c,
		Authority: middleware.NewAuthorityMiddleware(casbinEnforcer, redisClient).Handle,
		Converter: &generated.ConverterImpl{},
		Casbin:    casbinEnforcer,
	}
}
