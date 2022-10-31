package oauth

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"slash-admin/app/admin/cmd/api/internal/logic/oauth"
	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"
)

// swagger:route post /oauth/login oauth OauthLogin
//
// "Oauth log in | Oauth 登录"
//
// "Oauth log in | Oauth 登录"
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: OauthLoginReq
//
//
// Responses:
//  200: RedirectResp
//
//

func OauthLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OauthLoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := oauth.NewOauthLoginLogic(r.Context(), svcCtx)
		resp, err := l.OauthLogin(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
