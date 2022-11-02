package dictionary

import (
	"context"

	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDetailByDictionaryNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDetailByDictionaryNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDetailByDictionaryNameLogic {
	return &GetDetailByDictionaryNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDetailByDictionaryNameLogic) GetDetailByDictionaryName(req *types.DictionaryDetailReq) (resp *types.DictionaryDetailListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
