package role

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"slash-admin/app/admin/cmd/api/internal/logic/role"
	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"
)

// swagger:route post /role role CreateOrUpdateRole
//
// "Create or update role information | 创建或更新角色"
//
// "创建或更新角色"
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: RoleInfo
//
//
// Responses:
//  200: SimpleMsgResp
//
//

func CreateOrUpdateRoleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RoleInfo
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := role.NewCreateOrUpdateRoleLogic(r.Context(), svcCtx)
		resp, err := l.CreateOrUpdateRole(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
