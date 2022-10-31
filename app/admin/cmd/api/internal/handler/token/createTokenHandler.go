package token

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"slash-admin/app/admin/cmd/api/internal/logic/token"
	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"
)

// swagger:route post /token token CreateToken
//
//  Create Token information | 创建Token
//
//  Create Token information | 创建Token
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: CreateTokenReq
//
//
// Responses:
//  200: SimpleMsgResp
//
//

func CreateTokenHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateTokenReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := token.NewCreateTokenLogic(r.Context(), svcCtx)
		resp, err := l.CreateToken(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
