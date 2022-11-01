package api

import (
	"context"
	"github.com/zeromicro/go-zero/core/errorx"
	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"
	"slash-admin/app/admin/ent/predicate"
	"slash-admin/app/admin/ent/sysapi"
	"slash-admin/pkg/message"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetApiListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetApiListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApiListLogic {
	return &GetApiListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetApiListLogic) GetApiList(req *types.ApiListReq) (resp *types.ApiListResp, err error) {
	var predicates []predicate.SysApi

	if req.Path != nil {
		predicates = append(predicates, sysapi.PathContains(*req.Path))
	}

	if req.Method != nil {
		predicates = append(predicates, sysapi.Method(*req.Method))
	}

	if req.Group != nil {
		predicates = append(predicates, sysapi.APIGroup(*req.Group))
	}

	if req.Description != nil {
		predicates = append(predicates, sysapi.DescriptionContains(*req.Description))
	}

	page, err := l.svcCtx.EntClient.SysApi.Query().
		Where(predicates...).
		Page(l.ctx, req.PageNo, req.PageSize)

	if err != nil {
		logx.Errorw(message.DatabaseError, logx.Field("detail", err.Error()))
		return nil, errorx.NewApiInternalServerError(message.DatabaseError)
	}

	return &types.ApiListResp{
		Pagination: l.svcCtx.Converter.ConvertPagination(page.PageDetails),
		List:       l.svcCtx.Converter.ConvertSysApiList(page.List),
	}, nil
}
