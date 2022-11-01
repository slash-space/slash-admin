package oauth

import (
	"context"
	"github.com/zeromicro/go-zero/core/errorx"
	"slash-admin/app/admin/ent"
	"slash-admin/pkg/message"

	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteProviderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteProviderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProviderLogic {
	return &DeleteProviderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteProviderLogic) DeleteProvider(req *types.IDReq) (resp *types.SimpleMsgResp, err error) {
	err = l.svcCtx.EntClient.SysOauthProvider.DeleteOneID(req.ID).Exec(l.ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			logx.Errorw("Delete provider failed, check the id", logx.Field("ProviderId", req.ID))
			return nil, errorx.NewApiBadRequestError(errorx.DeleteFailed)
		}
		l.Errorw("delete provider error", logx.Field("detail", err.Error()))
		return nil, errorx.NewApiInternalServerError(message.DatabaseError)
	}

	l.Infow("Delete provider successfully", logx.Field("ProviderId", req.ID))

	return &types.SimpleMsgResp{Msg: errorx.DeleteSuccess}, nil
}
