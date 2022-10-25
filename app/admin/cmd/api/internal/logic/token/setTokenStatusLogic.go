package token

import (
	"context"

	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetTokenStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetTokenStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetTokenStatusLogic {
	return &SetTokenStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetTokenStatusLogic) SetTokenStatus(req *types.SetBooleanStatusReq) (resp *types.SimpleMsg, err error) {
	// todo: add your logic here and delete this line

	return
}
