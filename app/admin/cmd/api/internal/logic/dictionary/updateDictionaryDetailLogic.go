package dictionary

import (
	"context"
	"github.com/zeromicro/go-zero/core/errorx"
	"slash-admin/app/admin/ent"
	"slash-admin/app/admin/ent/sysdictionarydetail"
	"slash-admin/pkg/message"
	pType "slash-admin/pkg/types"

	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDictionaryDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateDictionaryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDictionaryDetailLogic {
	return &UpdateDictionaryDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateDictionaryDetailLogic) UpdateDictionaryDetail(req *types.UpdateDictionaryDetailReq) (resp *types.SimpleMsgResp, err error) {
	// check if the dictionary parent id exists
	exist, err := l.svcCtx.EntClient.SysDictionaryDetail.Query().Where(sysdictionarydetail.IDEQ(req.ID)).Exist(l.ctx)

	if err != nil {
		l.Errorw("check dictionary details id failed", logx.Field("details", err.Error()))
		return nil, errorx.NewApiInternalServerError(message.DatabaseError)
	}

	if !exist {
		l.Errorw("check if the dictionary details id exists", logx.Field("id", req.ID))
		return nil, errorx.NewApiBadRequestError(message.TargetNotExist)
	}

	details, err := l.svcCtx.EntClient.SysDictionaryDetail.UpdateOneID(req.ID).
		SetNillableTitle(req.Title).
		SetNillableKey(req.Key).
		SetNillableValue(req.Value).
		SetNillableStatus((*pType.Status)(req.Status)).
		Save(l.ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			l.Errorw("update dictionary detail failed", logx.Field("details", err.Error()))
			return nil, errorx.NewApiBadRequestError(message.UpdateFailed)
		}

		l.Errorw("create dictionary detail failed", logx.Field("details", err.Error()))
		return nil, errorx.NewApiInternalServerError(message.DatabaseError)
	}

	l.Infow("update dictionary detail success", logx.Field("details", details))

	return &types.SimpleMsgResp{Msg: message.UpdateSuccess}, nil
}
