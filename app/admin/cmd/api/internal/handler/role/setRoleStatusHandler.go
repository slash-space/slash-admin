package role

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"slash-admin/app/admin/cmd/api/internal/logic/role"
	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"
)

// swagger:route post /role/status role SetRoleStatus
//
// "Set role status | 设置角色状态"
//
// "设置角色状态"
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: SetStatusReq
//
//
// Responses:
//  200: SimpleMsgResp
//
//

func SetRoleStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SetStatusReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := role.NewSetRoleStatusLogic(r.Context(), svcCtx)
		resp, err := l.SetRoleStatus(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
