package api

import (
	"context"
	"github.com/zeromicro/go-zero/core/errorx"
	"slash-admin/app/admin/ent"
	"slash-admin/pkg/message"

	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteApiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteApiLogic {
	return &DeleteApiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteApiLogic) DeleteApi(req *types.IDReq) (resp *types.SimpleMsgResp, err error) {
	err = l.svcCtx.EntClient.SysApi.DeleteOneID(req.ID).Exec(l.ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			l.Errorw("delete api error", logx.Field("detail", err.Error()))
			return nil, errorx.NewApiBadRequestError(message.ApiNotExists)
		}
		l.Errorw("delete api error", logx.Field("detail", err.Error()))
		return nil, errorx.NewApiInternalServerError(message.DatabaseError)
	}

	return &types.SimpleMsgResp{Msg: errorx.Success}, nil
}
