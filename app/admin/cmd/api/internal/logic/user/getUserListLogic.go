package user

import (
	"context"
	"github.com/zeromicro/go-zero/core/errorx"
	"slash-admin/app/admin/ent/predicate"
	"slash-admin/app/admin/ent/sysuser"
	"slash-admin/pkg/message"

	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserListLogic) GetUserList(req *types.GetUserListReq) (resp *types.UserListResp, err error) {
	var predicates []predicate.SysUser
	if req.Username != "" {
		predicates = append(predicates, sysuser.UsernameContains(req.Username))
	}
	if req.Nickname != "" {
		predicates = append(predicates, sysuser.NicknameContains(req.Nickname))
	}
	if req.Email != "" {
		predicates = append(predicates, sysuser.EmailContains(req.Email))
	}
	if req.Mobile != "" {
		predicates = append(predicates, sysuser.MobileContains(req.Mobile))
	}
	if req.RoleId != 0 {
		predicates = append(predicates, sysuser.RoleIDEQ(req.RoleId))
	}

	pageResult, err := l.svcCtx.EntClient.SysUser.Query().Where(predicates...).Page(l.ctx, req.PageNo, req.PageSize)

	if err != nil {
		logx.Errorw(message.DatabaseError, logx.Field("detail", err.Error()))
		return nil, errorx.NewApiInternalServerError(message.DatabaseError)
	}

	return &types.UserListResp{
		Pagination: l.svcCtx.Converter.ConvertPagination(pageResult.PageDetails),
		List:       l.svcCtx.Converter.ConvertSysUserList(pageResult.List),
	}, nil
}
