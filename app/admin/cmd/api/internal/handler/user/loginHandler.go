package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"slash-admin/app/admin/cmd/api/internal/logic/user"
	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"
)

// swagger:route post /user/login user Login
//
//  Login | 登录
//
//  Login | 登录
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: LoginReq
//
//
// Responses:
//  200: LoginResp
//
//

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
