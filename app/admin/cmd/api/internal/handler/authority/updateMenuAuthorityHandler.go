package authority

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"slash-admin/app/admin/cmd/api/internal/logic/authority"
	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"
)

// swagger:route put /authority/menu authority UpdateMenuAuthority
//
// "Update menu authorization information | 更新菜单权限"
//
// "Update menu authorization information | 更新菜单权限"
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpdateMenuAuthorityReq
//
//
// Responses:
//  200: SimpleMsgResp
//
//

func UpdateMenuAuthorityHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateMenuAuthorityReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := authority.NewUpdateMenuAuthorityLogic(r.Context(), svcCtx)
		resp, err := l.UpdateMenuAuthority(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
