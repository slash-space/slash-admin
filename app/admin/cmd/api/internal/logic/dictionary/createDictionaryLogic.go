package dictionary

import (
	"context"

	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateDictionaryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateDictionaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDictionaryLogic {
	return &CreateDictionaryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateDictionaryLogic) CreateDictionary(req *types.CreateDictionaryReq) (resp *types.SimpleMsgResp, err error) {
	// todo: add your logic here and delete this line

	return
}
