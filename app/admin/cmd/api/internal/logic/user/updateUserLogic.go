package user

import (
	"context"
	"github.com/zeromicro/go-zero/core/errorx"
	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"
	pType "slash-admin/pkg/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserLogic) UpdateUser(req *types.UpdateUserReq) (resp *types.SimpleMsgResp, err error) {

	err = l.svcCtx.EntClient.SysUser.UpdateOneID(req.ID).
		SetNillableUsername(req.Username).
		SetNillableNickname(req.Nickname).
		SetNillablePassword(req.Password).
		SetNillableEmail(req.Email).
		SetNillableRoleID(req.RoleID).
		SetNillableAvatar(req.Avatar).
		SetNillableMobile(req.Mobile).
		SetNillableStatus((*pType.Status)(req.Status)).
		Exec(l.ctx)

	if err != nil {
		l.Errorw("update user error", logx.Field("err", err))
		return nil, errorx.NewApiBadRequestError(errorx.DatabaseError)
	}

	logx.Infow("update user successfully", logx.Field("detail", req))

	return &types.SimpleMsgResp{Msg: errorx.Success}, nil
}
