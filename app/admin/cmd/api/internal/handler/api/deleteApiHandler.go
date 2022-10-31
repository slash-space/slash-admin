package api

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"slash-admin/app/admin/cmd/api/internal/logic/api"
	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"
)

// swagger:route delete /api api DeleteApi
//
// "删除API信息"
//
// "Delete API information | 删除API信息"
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

func DeleteApiHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := api.NewDeleteApiLogic(r.Context(), svcCtx)
		resp, err := l.DeleteApi(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
