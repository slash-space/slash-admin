package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"slash-admin/app/admin/cmd/api/internal/config"
	"slash-admin/app/admin/cmd/api/internal/middleware"
)

type ServiceContext struct {
	Config    config.Config
	Authority rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		Authority: middleware.NewAuthorityMiddleware().Handle,
	}
}
