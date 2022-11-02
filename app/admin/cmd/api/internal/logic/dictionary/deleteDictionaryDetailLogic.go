package dictionary

import (
	"context"
	"github.com/zeromicro/go-zero/core/errorx"
	"slash-admin/app/admin/ent"
	"slash-admin/pkg/message"

	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteDictionaryDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteDictionaryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDictionaryDetailLogic {
	return &DeleteDictionaryDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteDictionaryDetailLogic) DeleteDictionaryDetail(req *types.IDReq) (resp *types.SimpleMsgResp, err error) {

	err = l.svcCtx.EntClient.SysDictionaryDetail.DeleteOneID(req.ID).Exec(l.ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			l.Errorw("Delete dictionary detail failed, check the id", logx.Field("DetailId", req.ID))
			return nil, errorx.NewApiBadRequestError(message.DeleteFailed)
		}

		l.Errorw("Delete dictionary detail failed", logx.Field("DetailId", req.ID))
		return nil, errorx.NewApiInternalServerError(message.DatabaseError)
	}

	l.Infow("Delete dictionary detail success", logx.Field("DetailId", req.ID))

	return &types.SimpleMsgResp{Msg: message.DeleteSuccess}, nil
}
