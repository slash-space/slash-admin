package token

import (
	"context"
	"github.com/zeromicro/go-zero/core/errorx"

	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTokenLogic {
	return &DeleteTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteTokenLogic) DeleteToken(req *types.IDReq) (resp *types.SimpleMsgResp, err error) {
	err = l.svcCtx.EntClient.SysToken.DeleteOneID(req.ID).Exec(l.ctx)

	if err != nil {
		l.Errorw("Delete token failed", logx.Field("detail", err.Error()))
		return nil, errorx.NewApiBadRequestError(err.Error())
	}

	l.Infow("Delete token successfully", logx.Field("TokenId", req.ID))

	return &types.SimpleMsgResp{Msg: "Delete token successfully"}, nil
}
