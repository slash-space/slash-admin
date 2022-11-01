package oauth

import (
	"context"
	"github.com/zeromicro/go-zero/core/errorx"
	"slash-admin/pkg/message"

	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProviderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProviderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProviderListLogic {
	return &GetProviderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProviderListLogic) GetProviderList(req *types.PageReq) (resp *types.ProviderListResp, err error) {
	pageResult, err := l.svcCtx.EntClient.SysOauthProvider.Query().Page(l.ctx, req.PageNo, req.PageSize)

	if err != nil {
		l.Error("query oauth provider list error: %v", err)
		return nil, errorx.NewApiInternalServerError(message.DatabaseError)
	}

	return &types.ProviderListResp{
		Pagination: l.svcCtx.Converter.ConvertPagination(pageResult.PageDetails),
		List:       l.svcCtx.Converter.ConvertSysOauthProviderList(pageResult.List),
	}, nil
}
