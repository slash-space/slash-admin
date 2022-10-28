package user

import (
	"context"
	"github.com/zeromicro/go-zero/core/errorx"
	"slash-admin/app/admin/cmd/api/internal/globalkey"
	"slash-admin/app/admin/ent"
	"slash-admin/app/admin/ent/sysuser"
	"slash-admin/pkg/message"

	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserProfileLogic {
	return &GetUserProfileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserProfileLogic) GetUserProfile() (resp *types.ProfileResp, err error) {

	UUID := l.ctx.Value(globalkey.JWTUserId).(string)

	user, err := l.svcCtx.EntClient.SysUser.Query().Where(sysuser.UUIDEQ(UUID)).First(l.ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errorx.NewApiBadRequestError(message.UserNotExists)
		}
		return nil, errorx.NewApiBadRequestError(errorx.DatabaseError)
	}

	resp.User = l.svcCtx.Converter.ConvertSysUser(user)

	return
}
