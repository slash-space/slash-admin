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

type UpdateDictionaryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateDictionaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDictionaryLogic {
	return &UpdateDictionaryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateDictionaryLogic) UpdateDictionary(req *types.UpdateDictionaryReq) (resp *types.SimpleMsgResp, err error) {

	exist, err := l.svcCtx.EntClient.SysDictionary.Query().Where(sysdictionary.IDEQ(req.ID)).Exist(l.ctx)

	if err != nil {
		l.Errorw("update dictionary error", logx.Field("detail", err.Error()))
		return nil, errorx.NewApiInternalServerError(message.DatabaseError)
	}

	if !exist {
		l.Errorw("update dictionary error", logx.Field("detail", "dictionary not exist"))
		return nil, errorx.NewApiInternalServerError(message.TargetNotExist)
	}

	details, err := l.svcCtx.EntClient.SysDictionary.UpdateOneID(req.ID).
		SetNillableTitle(req.Title).
		SetNillableName(req.Name).
		SetNillableStatus((*pType.Status)(req.Status)).
		SetNillableDesc(req.Description).
		Save(l.ctx)

	if err != nil {
		l.Errorw("update dictionary error", logx.Field("detail", err.Error()))
		return nil, errorx.NewApiBadRequestError(message.DatabaseError)
	}

	l.Infow("update dictionary success", logx.Field("detail", details))

	return &types.SimpleMsgResp{Msg: message.UpdateSuccess}, nil
}
