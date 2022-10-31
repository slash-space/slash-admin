package oauth

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"slash-admin/app/admin/cmd/api/internal/logic/oauth"
	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"
)

// swagger:route get /oauth/login/callback oauth OauthCallback
//
// "Oauth log in callback route | Oauth 登录回调接口"
//
// "Oauth 登录回调接口"
//
// Responses:
//  200: CallbackResp
//
//

func OauthCallbackHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OauthCallbackParamReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := oauth.NewOauthCallbackLogic(r.Context(), svcCtx)
		resp, err := l.OauthCallback(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
