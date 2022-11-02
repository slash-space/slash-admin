package menu

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"slash-admin/app/admin/cmd/api/internal/logic/menu"
	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"
)

// swagger:route post /menu/param menu CreateOrUpdateMenuParam
//
// "Create or update menu parameters | 创建或更新菜单参数"
//
// "Create or update menu parameters | 创建或更新菜单参数"
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: CreateOrUpdateMenuParamsReq
//
//
// Responses:
//  200: SimpleMsgResp
//
//

func CreateOrUpdateMenuParamHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateOrUpdateMenuParamsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := menu.NewCreateOrUpdateMenuParamLogic(r.Context(), svcCtx)
		resp, err := l.CreateOrUpdateMenuParam(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
