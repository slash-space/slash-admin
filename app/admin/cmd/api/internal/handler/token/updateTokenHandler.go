package token

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"slash-admin/app/admin/cmd/api/internal/logic/token"
	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"
)

// swagger:route put /token token UpdateToken
//
//  Update Token information | 更新Token
//
//  Update Token information | 更新Token
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpdateTokenReq
//
//
// Responses:
//  200: SimpleMsgResp
//
//

func UpdateTokenHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateTokenReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := token.NewUpdateTokenLogic(r.Context(), svcCtx)
		resp, err := l.UpdateToken(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
