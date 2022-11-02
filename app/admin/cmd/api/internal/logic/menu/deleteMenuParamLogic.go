package menu

import (
	"context"
	"github.com/zeromicro/go-zero/core/errorx"
	"slash-admin/app/admin/ent"
	"slash-admin/pkg/message"

	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMenuParamLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteMenuParamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMenuParamLogic {
	return &DeleteMenuParamLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteMenuParamLogic) DeleteMenuParam(req *types.IDReq) (resp *types.SimpleMsgResp, err error) {
	err = l.svcCtx.EntClient.SysMenuParam.DeleteOneID(req.ID).Exec(l.ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			l.Errorw("Delete Menu parameter failed, check the id", logx.Field("menuParamId", req.ID))
			return nil, errorx.NewApiBadRequestError(message.TargetNotExist)
		}
		l.Errorw("delete menu parameter error", logx.Field("id", req.ID), logx.Field("error", err))
		return nil, errorx.NewApiInternalServerError(message.DatabaseError)
	}

	return &types.SimpleMsgResp{Msg: message.DeleteSuccess}, nil
}
