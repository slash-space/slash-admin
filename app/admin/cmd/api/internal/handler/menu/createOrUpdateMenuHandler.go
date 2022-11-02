package menu

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"slash-admin/app/admin/cmd/api/internal/logic/menu"
	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"
)

// swagger:route post /menu menu CreateOrUpdateMenu
//
// "Create or update menu information | 创建或更新菜单"
//
// "Create or update menu information | 创建或更新菜单"
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: CreateOrUpdateMenuReq
//
//
// Responses:
//  200: SimpleMsgResp
//
//

func CreateOrUpdateMenuHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateOrUpdateMenuReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := menu.NewCreateOrUpdateMenuLogic(r.Context(), svcCtx)
		resp, err := l.CreateOrUpdateMenu(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
