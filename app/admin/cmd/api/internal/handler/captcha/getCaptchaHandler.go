package captcha

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"slash-admin/app/admin/cmd/api/internal/logic/captcha"
	"slash-admin/app/admin/cmd/api/internal/svc"
)

// swagger:route get /captcha captcha GetCaptcha
//
//  Get captcha | 获取验证码
//
//  Get captcha | 获取验证码
//
// Responses:
//  200: CaptchaInfoResp
//
//

func GetCaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := captcha.NewGetCaptchaLogic(r.Context(), svcCtx)
		resp, err := l.GetCaptcha()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
