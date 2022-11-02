package dictionary

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"slash-admin/app/admin/cmd/api/internal/logic/dictionary"
	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"
)

// swagger:route delete /dict/detail dictionary DeleteDictionaryDetail
//
// "Delete dictionary KV information | 删除字典键值信息"
//
// "Delete dictionary KV information | 删除字典键值信息"
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDReq
//
//
// Responses:
//  200: SimpleMsgResp
//
//

func DeleteDictionaryDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := dictionary.NewDeleteDictionaryDetailLogic(r.Context(), svcCtx)
		resp, err := l.DeleteDictionaryDetail(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
