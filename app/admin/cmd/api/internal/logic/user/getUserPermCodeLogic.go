package user

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/errorx"
	"net/http"
	"slash-admin/app/admin/cmd/api/internal/globalkey"

	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserPermCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserPermCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserPermCodeLogic {
	return &GetUserPermCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserPermCodeLogic) GetUserPermCode() (resp *types.PermCodeResp, err error) {
	roleId := l.ctx.Value(globalkey.JWTRoleId)
	if roleId == nil {
		return nil, &errorx.ApiError{Code: http.StatusUnauthorized, Msg: "sys.login.requireLogin"}
	}
	return &types.PermCodeResp{Data: []string{fmt.Sprintf("%v", roleId)}}, nil
}
