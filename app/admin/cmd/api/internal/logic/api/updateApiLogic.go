package api

import (
	"context"
	"github.com/zeromicro/go-zero/core/errorx"
	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"
	"slash-admin/app/admin/ent/sysapi"
	"slash-admin/pkg/message"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateApiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateApiLogic {
	return &UpdateApiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateApiLogic) UpdateApi(req *types.UpdateApiReq) (resp *types.SimpleMsgResp, err error) {
	exist, err := l.svcCtx.EntClient.SysApi.Query().Where(sysapi.ID(req.ID)).Exist(l.ctx)

	if err != nil {
		l.Errorw("update api error", logx.Field("detail", err.Error()))
		return nil, errorx.NewApiInternalServerError(message.DatabaseError)
	}

	if !exist {
		l.Errorw("update api error", logx.Field("detail", "api not exist"))
		return nil, errorx.NewApiInternalServerError(message.ApiNotExists)
	}

	err = l.svcCtx.EntClient.SysApi.UpdateOneID(req.ID).
		SetNillablePath(req.Path).
		SetNillableDescription(req.Description).
		SetNillableAPIGroup(req.Group).
		SetNillableMethod(req.Method).
		Exec(l.ctx)

	if err != nil {
		l.Errorw("update api error", logx.Field("detail", err.Error()))
		return nil, errorx.NewApiBadRequestError(message.DatabaseError)
	}

	return &types.SimpleMsgResp{
		Msg: errorx.UpdateSuccess,
	}, nil
}
