package authority

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"slash-admin/app/admin/cmd/api/internal/logic/authority"
	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"
)

// swagger:route post /authority/api authority CreateApiAuthority
//
//  Create or update API authorization information | 创建或更新API权限
//
//  Create or update API authorization information | 创建或更新API权限
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: CreateApiAuthorityReq
//
//
// Responses:
//  200: SimpleMsgResp
//
//

func CreateApiAuthorityHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateApiAuthorityReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := authority.NewCreateApiAuthorityLogic(r.Context(), svcCtx)
		resp, err := l.CreateApiAuthority(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
