package dictionary

import (
	"context"
	"github.com/zeromicro/go-zero/core/errorx"
	"slash-admin/app/admin/ent"
	"slash-admin/app/admin/ent/sysdictionarydetail"
	"slash-admin/pkg/message"
	"slash-admin/pkg/utils"

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
	err = utils.WithTx(l.ctx, l.svcCtx.EntClient, func(tx *ent.Tx) error {
		tx.SysDictionaryDetail.Delete().Where(sysdictionarydetail.DictionaryIDEQ(req.ID)).ExecX(l.ctx)
		tx.SysDictionary.DeleteOneID(req.ID).ExecX(l.ctx)
		return nil
	})

	if err != nil {
		l.Errorw("Delete dictionary failed", logx.Field("DictionaryId", req.ID))
		return nil, errorx.NewApiInternalServerError(message.DatabaseError)
	}

	l.Infow("Delete dictionary success", logx.Field("DictionaryId", req.ID))
	return &types.SimpleMsgResp{Msg: message.DeleteSuccess}, nil
}
