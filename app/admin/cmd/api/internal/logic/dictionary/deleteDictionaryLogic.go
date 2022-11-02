package dictionary

import (
	"context"

	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteDictionaryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteDictionaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDictionaryLogic {
	return &DeleteDictionaryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteDictionaryLogic) DeleteDictionary(req *types.IDReq) (resp *types.SimpleMsgResp, err error) {
	// todo: add your logic here and delete this line

	return
}
