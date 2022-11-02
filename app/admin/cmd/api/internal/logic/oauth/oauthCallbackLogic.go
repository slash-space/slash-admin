package oauth

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/errorx"
	"golang.org/x/oauth2"
	"io"
	"net/http"
	"slash-admin/app/admin/cmd/api/internal/logic/user"
	"slash-admin/app/admin/ent"
	"slash-admin/app/admin/ent/sysoauthprovider"
	"slash-admin/app/admin/ent/sysuser"
	"slash-admin/pkg/message"
	pType "slash-admin/pkg/types"
	"strings"
	"time"

	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OauthCallbackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

type userInfo struct {
	Email    string `json:"email"`
	NickName string `json:"nickName"`
	Picture  string `json:"picture"`
}

func NewOauthCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OauthCallbackLogic {
	return &OauthCallbackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OauthCallbackLogic) OauthCallback(req *types.OauthCallbackParamReq) (resp *types.CallbackResp, err error) {
	provider := strings.Split(req.State, "-")[1]

	if _, ok := providerConfig[provider]; !ok {
		target, err := l.svcCtx.EntClient.SysOauthProvider.Query().
			Where(sysoauthprovider.Name(provider)).
			First(l.ctx)

		if err != nil {
			if ent.IsNotFound(err) {
				logx.Errorw("provider not found", logx.Field("detail", target))
				return nil, errorx.NewApiBadRequestError(errorx.TargetNotExist)
			}
			l.Errorw("query provider error", logx.Field("detail", err.Error()))
			return nil, errorx.NewApiInternalServerError(message.DatabaseError)
		}

		providerConfig[provider] = oauth2.Config{
			ClientID:     target.ClientID,
			ClientSecret: target.ClientSecret,
			Endpoint: oauth2.Endpoint{
				AuthURL:   target.AuthURL,
				TokenURL:  target.TokenURL,
				AuthStyle: oauth2.AuthStyle(target.AuthStyle),
			},
			RedirectURL: target.RedirectURL,
			Scopes:      strings.Split(target.Scopes, " "),
		}
		if _, ok := userInfoURL[target.Name]; !ok {
			userInfoURL[target.Name] = target.InfoURL
		}
	}

	// get user information
	content, err := getUserInfo(providerConfig[provider], userInfoURL[provider], req.Code)
	if err != nil {
		return nil, errorx.NewApiBadRequestError(err.Error())
	}

	// find or register user
	var u userInfo
	err = json.Unmarshal(content, &u)

	if err != nil {
		l.Errorf("unmarshal user info error: %s", err.Error())
		return nil, errorx.NewApiInternalServerError(message.OAuthUserUnmarshalError)
	}

	if u.Email != "" {
		targetUser, err := l.svcCtx.EntClient.SysUser.Query().Where(sysuser.Email(u.Email)).WithRole().First(l.ctx)

		if err != nil {
			if ent.IsNotFound(err) {
				l.Errorf("user not found: %s", err.Error())
				return nil, errorx.NewApiBadRequestError(message.OAuthUserBindError)
			}
			l.Errorf("query user error: %s", err.Error())
			return nil, errorx.NewApiInternalServerError(message.DatabaseError)
		}

		token, err := user.GetSysUserJwtToken(
			targetUser,
			l.svcCtx.Config.Auth.AccessSecret,
			time.Now().Unix(),
			l.svcCtx.Config.Auth.AccessExpire,
		)

		if err != nil {
			l.Errorf("get user token error: %s", err.Error())
			return nil, errorx.NewApiInternalServerError(message.GenerateTokenError)
		}

		expiredAt := time.Now().Add(time.Second * 259200).Unix()
		source := strings.Split(req.State, "-")[1]

		_, err = l.svcCtx.EntClient.SysToken.Create().
			SetUUID(targetUser.UUID).
			SetToken(token).
			SetStatus(pType.StatusNormal).
			SetSource(source).
			SetExpiredAt(time.Unix(expiredAt, 0)).
			Save(l.ctx)

		if err != nil {
			logx.Errorw(message.DatabaseError, logx.Field("detail", err.Error()))
			return nil, errorx.NewApiInternalServerError(err.Error())
		}

		roleMeta := &types.RoleMetaInfo{
			RoleName: "",
			Value:    "",
		}
		if targetUser.Edges.Role != nil {
			roleMeta.RoleName = targetUser.Edges.Role.Name
			roleMeta.Value = targetUser.Edges.Role.Value
		}

		return &types.CallbackResp{
			UserId:    targetUser.UUID,
			Role:      roleMeta,
			Token:     token,
			ExpiredAt: uint64(expiredAt),
		}, nil
	}

	return nil, errorx.NewApiBadRequestError(message.OAuthUserEmailBindError)
}

func getUserInfo(c oauth2.Config, infoURL string, code string) ([]byte, error) {
	token, err := c.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	var response *http.Response
	if c.Endpoint.AuthStyle == 1 {
		response, err = http.Get(infoURL + token.AccessToken)
		if err != nil {
			return nil, fmt.Errorf("failed getting user info: %s", err.Error())
		}
	} else if c.Endpoint.AuthStyle == 2 {
		client := &http.Client{}
		request, err := http.NewRequest("GET", infoURL, nil)
		if err != nil {
			return nil, fmt.Errorf("failed getting user info: %s", err.Error())
		}

		request.Header.Set("Accept", "application/json")
		request.Header.Set("Authorization", "Bearer "+token.AccessToken)

		response, err = client.Do(request)
		if err != nil {
			return nil, fmt.Errorf("failed getting user info: %s", err.Error())
		}
	}

	defer response.Body.Close()
	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	return contents, nil
}
