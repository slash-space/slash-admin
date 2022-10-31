package token

import (
	"context"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"
	"slash-admin/app/admin/ent"
	"slash-admin/app/admin/ent/predicate"
	"slash-admin/app/admin/ent/systoken"
	"slash-admin/app/admin/ent/sysuser"
	"slash-admin/pkg/message"
)

type GetTokenListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTokenListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTokenListLogic {
	return &GetTokenListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTokenListLogic) GetTokenList(req *types.TokenListReq) (resp *types.TokenListResp, err error) {
	if req.UUID == nil && req.Username == nil && req.Nickname == nil && req.Email == nil {
		pageResult, err := l.svcCtx.EntClient.SysToken.Query().Page(l.ctx, req.PageNo, req.PageSize)

		if err != nil {
			logx.Errorw(errorx.DatabaseError, logx.Field("detail", err.Error()))
			return nil, errorx.NewApiInternalServerError(errorx.DatabaseError)
		}

		return &types.TokenListResp{
			Pagination: l.svcCtx.Converter.ConvertPagination(pageResult.PageDetails),
			List:       l.svcCtx.Converter.ConvertSysTokenList(pageResult.List),
		}, nil
	}

	var predicates []predicate.SysUser

	if req.UUID != nil {
		predicates = append(predicates, sysuser.UUID(*req.UUID))
	}
	if req.Username != nil {
		predicates = append(predicates, sysuser.Username(*req.Username))
	}
	if req.Nickname != nil {
		predicates = append(predicates, sysuser.Nickname(*req.Nickname))
	}
	if req.Email != nil {
		predicates = append(predicates, sysuser.Email(*req.Email))
	}

	user, err := l.svcCtx.EntClient.SysUser.Query().Where(predicates...).First(l.ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			l.Errorf("query user not found, %v", err)
			return nil, errorx.NewApiBadRequestError(message.UserNotExists)
		}
		l.Errorf("query user error, %v", err)
		return nil, errorx.NewApiBadRequestError(errorx.DatabaseError)
	}

	pageResult, err := l.svcCtx.EntClient.SysToken.Query().
		Where(systoken.UUIDEQ(user.UUID)).
		Page(l.ctx, req.PageNo, req.PageSize)

	if err != nil {
		logx.Errorw(errorx.DatabaseError, logx.Field("detail", err.Error()))
		return nil, errorx.NewApiInternalServerError(errorx.DatabaseError)
	}

	return &types.TokenListResp{
		Pagination: l.svcCtx.Converter.ConvertPagination(pageResult.PageDetails),
		List:       l.svcCtx.Converter.ConvertSysTokenList(pageResult.List),
	}, nil
}
