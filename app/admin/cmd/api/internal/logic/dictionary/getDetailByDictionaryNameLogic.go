package dictionary

import (
	"context"
	"github.com/zeromicro/go-zero/core/errorx"
	"slash-admin/app/admin/ent/sysdictionary"
	"slash-admin/pkg/message"

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

	dict, err := l.svcCtx.EntClient.SysDictionary.Query().
		Where(sysdictionary.NameEQ(req.Name)).
		WithDetails().
		First(l.ctx)

	if err != nil {
		l.Errorw("query dictionary detail error", logx.Field("details", err.Error()))
		return nil, errorx.NewApiInternalServerError(message.DatabaseError)
	}

	return &types.DictionaryDetailListResp{
		Pagination: &types.Pagination{
			Total: len(dict.Edges.Details),
		},
		List: l.svcCtx.Converter.ConvertSysDictionaryDetailList(dict.Edges.Details),
	}, nil
}
