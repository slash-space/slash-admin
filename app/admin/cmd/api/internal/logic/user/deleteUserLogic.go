package user

import (
	"context"
	"github.com/zeromicro/go-zero/core/errorx"
	"slash-admin/app/admin/ent"
	"slash-admin/pkg/message"

	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserLogic) DeleteUser(req *types.IDReq) (resp *types.SimpleMsgResp, err error) {
	err = l.svcCtx.EntClient.SysUser.DeleteOneID(req.ID).Exec(l.ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			logx.Errorw("delete user failed, please check the user id", logx.Field("userId", req.ID))
			return &types.SimpleMsgResp{Msg: errorx.DeleteFailed}, nil
		}
		logx.Errorw(message.DatabaseError, logx.Field("detail", err.Error()))
		return nil, errorx.NewApiBadRequestError(message.DatabaseError)
	}

	logx.Infow("delete user successfully", logx.Field("userId", req.ID))

	return &types.SimpleMsgResp{Msg: errorx.DeleteSuccess}, nil
}
