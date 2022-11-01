package user

import (
	"context"
	"github.com/zeromicro/go-zero/core/errorx"
	"slash-admin/app/admin/cmd/api/internal/globalkey"
	"slash-admin/app/admin/ent/sysuser"
	"slash-admin/pkg/message"

	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserProfileLogic {
	return &UpdateUserProfileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserProfileLogic) UpdateUserProfile(req *types.UpdateProfileReq) (resp *types.SimpleMsgResp, err error) {
	uuid := l.ctx.Value(globalkey.JWTUserId).(string)

	updateOne := l.svcCtx.EntClient.SysUser.Update().
		Where(sysuser.UUID(uuid)).
		SetNillableAvatar(req.Avatar).
		SetNillableMobile(req.Mobile).
		SetNillableEmail(req.Email)

	if req.Nickname != nil {
		updateOne.SetNickname(*req.Nickname)
	}

	affect, err := updateOne.Save(l.ctx)

	if err != nil {
		l.Errorf("UpdateUserProfile error: %v", err)
		return nil, errorx.NewApiBadRequestError(message.DatabaseError)
	}

	if affect == 0 {
		return nil, errorx.NewApiBadRequestError(errorx.UpdateFailed)
	}

	return &types.SimpleMsgResp{
		Msg: errorx.UpdateSuccess,
	}, nil
}
