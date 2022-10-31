package oauth

import (
	"context"
	"github.com/zeromicro/go-zero/core/errorx"
	"slash-admin/app/admin/ent/sysoauthprovider"
	"slash-admin/pkg/message"

	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProviderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateProviderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProviderLogic {
	return &UpdateProviderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateProviderLogic) UpdateProvider(req *types.UpdateProviderReq) (resp *types.SimpleMsgResp, err error) {
	// check if provider exists
	exist, err := l.svcCtx.EntClient.SysOauthProvider.
		Query().
		Where(sysoauthprovider.IDEQ(req.ID)).
		Exist(l.ctx)

	if err != nil {
		l.Errorf("check provider exist error: %v", err)
		return nil, errorx.NewApiInternalServerError(errorx.DatabaseError)
	}

	if !exist {
		l.Errorf("provider not exist: %v", err)
		return nil, errorx.NewApiBadRequestError(message.ProviderNotExists)
	}

	err = l.svcCtx.EntClient.SysOauthProvider.UpdateOneID(req.ID).
		SetNillableName(req.Name).
		SetNillableClientID(req.ClientID).
		SetNillableClientSecret(req.ClientSecret).
		SetNillableRedirectURL(req.RedirectURL).
		SetNillableScopes(req.Scopes).
		SetNillableAuthURL(req.AuthURL).
		SetNillableTokenURL(req.TokenURL).
		SetNillableAuthStyle(req.AuthStyle).
		SetNillableInfoURL(req.InfoURL).
		Exec(l.ctx)

	if err != nil {
		l.Errorf("update provider error: %v", err)
		return nil, errorx.NewApiInternalServerError(errorx.DatabaseError)
	}

	l.Infow("update provider success", logx.Field("detail", req))

	return &types.SimpleMsgResp{Msg: errorx.UpdateSuccess}, nil
}
