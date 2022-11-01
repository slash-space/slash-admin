package authority

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"slash-admin/app/admin/cmd/api/internal/logic/authority"
	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"
)

// swagger:route put /authority/api authority UpdateApiAuthority
//
//  Get API authorization information | 获取API权限
//
//  Get API authorization information | 获取API权限
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpdateApiAuthorityReq
//
//
// Responses:
//  200: SimpleMsgResp
//
//

func UpdateApiAuthorityHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateApiAuthorityReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := authority.NewUpdateApiAuthorityLogic(r.Context(), svcCtx)
		resp, err := l.UpdateApiAuthority(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
