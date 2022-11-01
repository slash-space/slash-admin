package authority

import (
	"context"
	"github.com/zeromicro/go-zero/core/errorx"
	"slash-admin/app/admin/ent"
	"slash-admin/app/admin/ent/sysrole"
	"slash-admin/pkg/message"

	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuAuthorityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMenuAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuAuthorityLogic {
	return &GetMenuAuthorityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMenuAuthorityLogic) GetMenuAuthority(req *types.IDReq) (resp *types.MenuAuthorityInfoResp, err error) {
	role, err := l.svcCtx.EntClient.SysRole.Query().Where(sysrole.IDEQ(req.ID)).WithMenus().First(l.ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			l.Errorf("query role by id %d not found", req.ID)
			return nil, errorx.NewApiBadRequestError(message.RoleNotExists)
		}
		return nil, errorx.NewApiBadRequestError(message.DatabaseError)
	}

	resp = &types.MenuAuthorityInfoResp{
		RoleId:  req.ID,
		MenuIds: []uint64{},
	}

	if role.Edges.Menus != nil {
		for _, v := range role.Edges.Menus {
			resp.MenuIds = append(resp.MenuIds, v.ID)
		}
	}
	return
}
