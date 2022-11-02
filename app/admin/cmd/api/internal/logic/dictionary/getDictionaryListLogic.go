package dictionary

import (
	"context"
	"github.com/zeromicro/go-zero/core/errorx"
	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"
	"slash-admin/app/admin/ent/predicate"
	"slash-admin/app/admin/ent/sysdictionary"
	"slash-admin/pkg/message"

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
	var predicates []predicate.SysDictionary

	if req.Name != nil {
		predicates = append(predicates, sysdictionary.NameContains(*req.Name))
	}

	if req.Title != nil {
		predicates = append(predicates, sysdictionary.TitleContains(*req.Title))
	}

	pageResult, err := l.svcCtx.EntClient.SysDictionary.Query().
		Where(predicates...).
		Page(l.ctx, req.PageNo, req.PageSize)

	if err != nil {
		l.Errorw("query dictionary list error", logx.Field("details", err.Error()))
		return nil, errorx.NewApiInternalServerError(message.DatabaseError)
	}

	return &types.DictionaryListResp{
		Pagination: l.svcCtx.Converter.ConvertPagination(pageResult.PageDetails),
		List:       l.svcCtx.Converter.ConvertSysDictionaryList(pageResult.List),
	}, nil
}
