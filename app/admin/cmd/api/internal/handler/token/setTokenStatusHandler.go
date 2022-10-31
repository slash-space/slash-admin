package token

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"slash-admin/app/admin/cmd/api/internal/logic/token"
	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"
)

// swagger:route post /token/status token SetTokenStatus
//
// "Set token status | 设置token状态"
//
// "Set token status"
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: SetBooleanStatusReq
//
//
// Responses:
//  200: SimpleMsgResp
//
//

func SetTokenStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SetBooleanStatusReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := token.NewSetTokenStatusLogic(r.Context(), svcCtx)
		resp, err := l.SetTokenStatus(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
