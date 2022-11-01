package oauth

import (
	"context"
	"github.com/zeromicro/go-zero/core/errorx"
	"golang.org/x/oauth2"
	"slash-admin/app/admin/ent"
	"slash-admin/app/admin/ent/sysoauthprovider"
	"slash-admin/pkg/message"
	"strings"

	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

var providerConfig = make(map[string]oauth2.Config)

// userInfoURL used to store infoURL in database | 用来存储获取用户信息网址数据
var userInfoURL = make(map[string]string)

type OauthLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOauthLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OauthLoginLogic {
	return &OauthLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OauthLoginLogic) OauthLogin(req *types.OauthLoginReq) (resp *types.RedirectResp, err error) {

	provider, err := l.svcCtx.EntClient.SysOauthProvider.
		Query().
		Where(sysoauthprovider.Name(req.Provider)).
		First(l.ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			l.Errorw("oauth provider not found", logx.Field("provider", req.Provider))
			return nil, errorx.NewApiBadRequestError(message.ProviderNotExists)
		}
		l.Error("query oauth provider error: %v", err)
		return nil, errorx.NewApiInternalServerError(message.DatabaseError)
	}

	var config oauth2.Config

	if v, ok := providerConfig[provider.Name]; ok {
		config = v
	} else {
		providerConfig[provider.Name] = oauth2.Config{
			ClientID:     provider.ClientID,
			ClientSecret: provider.ClientSecret,
			Endpoint: oauth2.Endpoint{
				AuthURL:   provider.AuthURL,
				TokenURL:  provider.TokenURL,
				AuthStyle: oauth2.AuthStyle(provider.AuthStyle),
			},
			RedirectURL: provider.RedirectURL,
			Scopes:      strings.Split(provider.Scopes, " "),
		}
		config = providerConfig[provider.Name]
	}

	if _, ok := userInfoURL[provider.Name]; !ok {
		userInfoURL[provider.Name] = provider.InfoURL
	}

	url := config.AuthCodeURL(req.State)

	l.Infow("oauth login url redirect", logx.Field("url", url))

	return &types.RedirectResp{
		URL: url,
	}, nil
}
