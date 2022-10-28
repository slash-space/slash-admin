package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"slash-admin/app/admin/cmd/api/internal/logic/user"
	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"
)

// swagger:route post /user user CreateOrUpdateUser
//
// "Create or update user's information | 新增或更新用户"
//
// "新增或更新用户"
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: CreateOrUpdateUserReq
//
//
// Responses:
//  200: SimpleMsgResp
//
//

func CreateOrUpdateUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateOrUpdateUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewCreateOrUpdateUserLogic(r.Context(), svcCtx)
		resp, err := l.CreateOrUpdateUser(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
