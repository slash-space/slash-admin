package api

import (
	"context"
	"github.com/zeromicro/go-zero/core/errorx"

	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateApiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateApiLogic {
	return &CreateApiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateApiLogic) CreateApi(req *types.CreateApiReq) (resp *types.SimpleMsgResp, err error) {

	err = l.svcCtx.EntClient.SysApi.Create().
		SetPath(req.Path).
		SetDescription(req.Description).
		SetAPIGroup(req.Group).
		SetMethod(req.Method).
		Exec(l.ctx)

	if err != nil {
		l.Errorw("create api error", logx.Field("detail", err.Error()))
		return nil, errorx.NewApiInternalServerError(errorx.DatabaseError)
	}

	l.Infow("create api success",
		logx.Field("path", req.Path),
		logx.Field("desc", req.Description),
		logx.Field("group", req.Group),
		logx.Field("method", req.Method),
	)

	return &types.SimpleMsgResp{Msg: errorx.CreateSuccess}, nil

}
