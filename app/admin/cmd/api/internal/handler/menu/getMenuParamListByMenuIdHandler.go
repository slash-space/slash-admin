package menu

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"slash-admin/app/admin/cmd/api/internal/logic/menu"
	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"
)

// swagger:route post /menu/param/list menu GetMenuParamListByMenuId
//
// "Get menu extra parameters by menu ID | 获取某个菜单的额外参数列表"
//
// "Get menu extra parameters by menu ID | 获取某个菜单的额外参数列表"
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDReq
//
//
// Responses:
//  200: MenuParamListByMenuIdResp
//
//

func GetMenuParamListByMenuIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := menu.NewGetMenuParamListByMenuIdLogic(r.Context(), svcCtx)
		resp, err := l.GetMenuParamListByMenuId(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
