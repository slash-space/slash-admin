package core

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"slash-admin/app/admin/cmd/api/internal/logic/core"
	"slash-admin/app/admin/cmd/api/internal/svc"
)

// swagger:route get /core/health core HealthCheck
//
// "Check the system status | 检查系统状态"
//
// "Check the system status | 检查系统状态"
//

func HealthCheckHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := core.NewHealthCheckLogic(r.Context(), svcCtx)
		err := l.HealthCheck()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
