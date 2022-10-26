package core

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"slash-admin/app/admin/cmd/api/internal/logic/core"
	"slash-admin/app/admin/cmd/api/internal/svc"
)

// swagger:route get /core/init/database core InitDatabase
//
// "Initialize database | 初始化数据库"
//
// "初始化数据库"
//
// Responses:
//  200: SimpleMsgResp
//
//

func InitDatabaseHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := core.NewInitDatabaseLogic(r.Context(), svcCtx)
		resp, err := l.InitDatabase()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
