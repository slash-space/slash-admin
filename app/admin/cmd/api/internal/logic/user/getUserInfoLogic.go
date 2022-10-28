package user

import (
	"context"
	"github.com/zeromicro/go-zero/core/errorx"
	"net/http"
	"slash-admin/app/admin/cmd/api/internal/globalkey"
	"slash-admin/app/admin/ent"
	"slash-admin/app/admin/ent/sysuser"
	"slash-admin/pkg/message"

	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo() (resp *types.GetUserInfoResp, err error) {
	uuid := l.ctx.Value(globalkey.JWTUserId).(string)

	if uuid == "" {
		l.Errorf("GetUserInfoLogic.GetUserInfo: %s", "Please log in")
		return nil, errorx.NewApiError(http.StatusUnauthorized, "Please log in")
	}

	user, err := l.svcCtx.EntClient.SysUser.Query().Where(sysuser.UUIDEQ(uuid)).First(l.ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			l.Errorf("GetUserInfoLogic.GetUserInfo: %s", message.UserNotExists)
			return nil, errorx.NewApiBadRequestError(message.UserNotExists)
		}
		return nil, errorx.NewApiBadRequestError(errorx.DatabaseError)
	}

	var roleMeta = &types.RoleMetaInfo{}

	if user.Edges.Role != nil {
		roleMeta.RoleName = user.Edges.Role.Name
		roleMeta.Value = user.Edges.Role.Value
	}

	return &types.GetUserInfoResp{
		User:  l.svcCtx.Converter.ConvertSysUser(user),
		Roles: roleMeta,
	}, nil
}
