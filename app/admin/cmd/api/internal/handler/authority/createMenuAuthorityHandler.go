package authority

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"slash-admin/app/admin/cmd/api/internal/logic/authority"
	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"
)

// swagger:route post /authority/menu authority CreateMenuAuthority
//
// "Create menu authorization information | 创建菜单权限"
//
// "Create menu authorization information | 创建菜单权限"
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: CreateMenuAuthorityReq
//
//
// Responses:
//  200: SimpleMsgResp
//
//

func CreateMenuAuthorityHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateMenuAuthorityReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := authority.NewCreateMenuAuthorityLogic(r.Context(), svcCtx)
		resp, err := l.CreateMenuAuthority(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
