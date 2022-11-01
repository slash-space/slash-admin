package api

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"slash-admin/app/admin/cmd/api/internal/logic/api"
	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"
)

// swagger:route put /api api UpdateApi
//
//  Update API information | 创建或更新API
//
//  Update API information | 创建或更新API
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpdateApiReq
//
//
// Responses:
//  200: SimpleMsgResp
//
//

func UpdateApiHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateApiReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := api.NewUpdateApiLogic(r.Context(), svcCtx)
		resp, err := l.UpdateApi(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
