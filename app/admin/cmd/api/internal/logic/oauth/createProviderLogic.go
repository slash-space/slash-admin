package oauth

import (
	"context"
	"github.com/zeromicro/go-zero/core/errorx"
	"slash-admin/pkg/message"

	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateProviderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateProviderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateProviderLogic {
	return &CreateProviderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateProviderLogic) CreateProvider(req *types.CreateProviderReq) (resp *types.SimpleMsgResp, err error) {
	err = l.svcCtx.EntClient.SysOauthProvider.Create().
		SetName(req.Name).
		SetClientID(req.ClientID).
		SetClientSecret(req.ClientSecret).
		SetRedirectURL(req.RedirectURL).
		SetScopes(req.Scopes).
		SetAuthURL(req.AuthURL).
		SetTokenURL(req.TokenURL).
		SetAuthStyle(req.AuthStyle).
		SetInfoURL(req.InfoURL).
		Exec(context.Background())

	if err != nil {
		l.Errorf("create provider error: %v", err)
		return nil, errorx.NewApiInternalServerError(message.DatabaseError)
	}

	return &types.SimpleMsgResp{Msg: errorx.CreateSuccess}, nil
}
