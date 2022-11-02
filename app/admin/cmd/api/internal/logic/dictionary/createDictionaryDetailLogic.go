package dictionary

import (
	"context"
	"github.com/zeromicro/go-zero/core/errorx"
	"slash-admin/app/admin/ent/sysdictionary"
	"slash-admin/pkg/message"
	pType "slash-admin/pkg/types"

	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateDictionaryDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateDictionaryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDictionaryDetailLogic {
	return &CreateDictionaryDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateDictionaryDetailLogic) CreateDictionaryDetail(req *types.CreateDictionaryDetailReq) (resp *types.SimpleMsgResp, err error) {
	// check if the dictionary parent id exists
	exist, err := l.svcCtx.EntClient.SysDictionary.Query().Where(sysdictionary.IDEQ(req.DictionaryID)).Exist(l.ctx)

	if err != nil {
		l.Errorw("check dictionary id failed", logx.Field("details", err.Error()))
		return nil, errorx.NewApiInternalServerError(message.DatabaseError)
	}

	if !exist {
		l.Errorw("check if the dictionary parent id exists", logx.Field("parentId", req.DictionaryID))
		return nil, errorx.NewApiBadRequestError(message.ParentNotExist)
	}

	details, err := l.svcCtx.EntClient.SysDictionaryDetail.Create().
		SetTitle(req.Title).
		SetKey(req.Key).
		SetValue(req.Value).
		SetStatus(pType.Status(req.Status)).
		SetDictionaryID(req.DictionaryID).
		Save(l.ctx)

	if err != nil {
		l.Errorw("create dictionary detail failed", logx.Field("details", err.Error()))
		return nil, errorx.NewApiInternalServerError(message.DatabaseError)
	}

	l.Infow("create dictionary detail success", logx.Field("details", details))

	return &types.SimpleMsgResp{Msg: message.CreateSuccess}, nil
}
