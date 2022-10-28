// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	captcha "slash-admin/app/admin/cmd/api/internal/handler/captcha"
	core "slash-admin/app/admin/cmd/api/internal/handler/core"
	role "slash-admin/app/admin/cmd/api/internal/handler/role"
	user "slash-admin/app/admin/cmd/api/internal/handler/user"
	"slash-admin/app/admin/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/core/health",
				Handler: core.HealthCheckHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/core/init/database",
				Handler: core.InitDatabaseHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/role",
					Handler: role.CreateOrUpdateRoleHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/role",
					Handler: role.DeleteRoleHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/role/list",
					Handler: role.GetRoleListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/role/status",
					Handler: role.SetRoleStatusHandler(serverCtx),
				},
			}...,
		),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/captcha",
				Handler: captcha.GetCaptchaHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: user.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/register",
				Handler: user.RegisterHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/user/change-password",
					Handler: user.ChangePasswordHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/user/info",
					Handler: user.GetUserInfoHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user",
					Handler: user.CreateOrUpdateUserHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/list",
					Handler: user.GetUserListHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/user",
					Handler: user.DeleteUserHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/user/perm",
					Handler: user.GetUserPermCodeHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/user/profile",
					Handler: user.GetUserProfileHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/profile",
					Handler: user.UpdateUserProfileHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/user/logout",
					Handler: user.LogoutHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
