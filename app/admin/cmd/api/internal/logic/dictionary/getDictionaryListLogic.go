package dictionary

import (
	"context"

	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDictionaryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDictionaryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDictionaryListLogic {
	return &GetDictionaryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDictionaryListLogic) GetDictionaryList(req *types.DictionaryListReq) (resp *types.DictionaryListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
