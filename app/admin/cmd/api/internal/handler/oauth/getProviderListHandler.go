package oauth

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"slash-admin/app/admin/cmd/api/internal/logic/oauth"
	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"
)

// swagger:route post /oauth/provider/list oauth GetProviderList
//
// "Get provider list | 获取提供商列表"
//
// "获取提供商列表"
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: PageReq
//
//
// Responses:
//  200: ProviderListResp
//
//

func GetProviderListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PageReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := oauth.NewGetProviderListLogic(r.Context(), svcCtx)
		resp, err := l.GetProviderList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
