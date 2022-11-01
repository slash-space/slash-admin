package role

import (
	"context"
	"github.com/zeromicro/go-zero/core/errorx"
	"slash-admin/pkg/message"

	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleListLogic {
	return &GetRoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRoleListLogic) GetRoleList(req *types.PageReq) (resp *types.RoleListResp, err error) {
	page, err := l.svcCtx.EntClient.SysRole.Query().Page(l.ctx, req.PageNo, req.PageSize)

	if err != nil {
		l.Errorw(message.DatabaseError, logx.Field("detail", err.Error()))
		return nil, errorx.NewApiInternalServerError(message.DatabaseError)
	}

	return &types.RoleListResp{
		Pagination: l.svcCtx.Converter.ConvertPagination(page.PageDetails),
		List:       l.svcCtx.Converter.ConvertSysRoleToRoleInfoList(page.List),
	}, nil

}
