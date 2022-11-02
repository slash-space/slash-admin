package dictionary

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"slash-admin/app/admin/cmd/api/internal/logic/dictionary"
	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"
)

// swagger:route post /dict/detail dictionary CreateDictionaryDetail
//
// "Create dictionary KV information | 创建字典键值信息"
//
// "Create dictionary KV information | 创建字典键值信息"
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: CreateDictionaryDetailReq
//
//
// Responses:
//  200: SimpleMsgResp
//
//

func CreateDictionaryDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateDictionaryDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := dictionary.NewCreateDictionaryDetailLogic(r.Context(), svcCtx)
		resp, err := l.CreateDictionaryDetail(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
