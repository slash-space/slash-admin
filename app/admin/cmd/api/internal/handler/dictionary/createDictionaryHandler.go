package dictionary

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"slash-admin/app/admin/cmd/api/internal/logic/dictionary"
	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"
)

// swagger:route post /dict dictionary CreateDictionary
//
// "Create dictionary information | 创建字典信息"
//
// "Create dictionary information | 创建字典信息"
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: CreateDictionaryReq
//
//
// Responses:
//  200: SimpleMsgResp
//
//

func CreateDictionaryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateDictionaryReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := dictionary.NewCreateDictionaryLogic(r.Context(), svcCtx)
		resp, err := l.CreateDictionary(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
