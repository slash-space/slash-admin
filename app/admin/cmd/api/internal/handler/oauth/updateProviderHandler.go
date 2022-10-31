package oauth

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"slash-admin/app/admin/cmd/api/internal/logic/oauth"
	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"
)

// swagger:route put /oauth/provider oauth UpdateProvider
//
//  Update provider information | 更新提供商信息
//
//  Update provider information | 更新提供商信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpdateProviderReq
//
//
// Responses:
//  200: SimpleMsgResp
//
//

func UpdateProviderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateProviderReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := oauth.NewUpdateProviderLogic(r.Context(), svcCtx)
		resp, err := l.UpdateProvider(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
