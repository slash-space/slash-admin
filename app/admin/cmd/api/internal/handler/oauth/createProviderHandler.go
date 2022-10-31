package oauth

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"slash-admin/app/admin/cmd/api/internal/logic/oauth"
	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"
)

// swagger:route post /oauth/provider oauth CreateProvider
//
// "Create provider information | 创建提供商信息"
//
// "创建提供商信息"
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: CreateProviderReq
//
//
// Responses:
//  200: SimpleMsgResp
//
//

func CreateProviderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateProviderReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := oauth.NewCreateProviderLogic(r.Context(), svcCtx)
		resp, err := l.CreateProvider(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
