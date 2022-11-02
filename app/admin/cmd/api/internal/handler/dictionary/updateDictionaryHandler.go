package dictionary

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"slash-admin/app/admin/cmd/api/internal/logic/dictionary"
	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"
)

// swagger:route put /dict dictionary UpdateDictionary
//
// "Update dictionary information | 更新字典信息"
//
// "Update dictionary information | 更新字典信息"
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpdateDictionaryReq
//
//
// Responses:
//  200: SimpleMsgResp
//
//

func UpdateDictionaryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateDictionaryReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := dictionary.NewUpdateDictionaryLogic(r.Context(), svcCtx)
		resp, err := l.UpdateDictionary(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
