package dictionary

import (
	"context"
	"github.com/zeromicro/go-zero/core/errorx"
	"slash-admin/pkg/message"
	pType "slash-admin/pkg/types"

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

	details, err := l.svcCtx.EntClient.SysDictionary.Create().
		SetTitle(req.Title).
		SetName(req.Name).
		SetStatus(pType.Status(req.Status)).
		SetDesc(req.Description).
		Save(l.ctx)

	if err != nil {
		l.Errorw("create dictionary error", logx.Field("detail", err.Error()))
		return nil, errorx.NewApiBadRequestError(message.DatabaseError)
	}

	l.Infow("create dictionary success", logx.Field("detail", details))

	return &types.SimpleMsgResp{Msg: message.CreateSuccess}, nil
}
