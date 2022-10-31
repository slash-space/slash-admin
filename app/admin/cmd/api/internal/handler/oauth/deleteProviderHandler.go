package oauth

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"slash-admin/app/admin/cmd/api/internal/logic/oauth"
	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"
)

// swagger:route delete /oauth/provider oauth DeleteProvider
//
//  Delete provider information | 删除提供商信息
//
//  Delete provider information | 删除提供商信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDReq
//
//
// Responses:
//  200: SimpleMsgResp
//
//

func DeleteProviderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := oauth.NewDeleteProviderLogic(r.Context(), svcCtx)
		resp, err := l.DeleteProvider(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
