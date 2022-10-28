package captcha

import (
	"context"
	"github.com/mojocn/base64Captcha"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"net/http"
	"slash-admin/app/admin/cmd/api/internal/config"
	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"
	"slash-admin/pkg/captcha"
	"slash-admin/pkg/message"
	"sync"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

var (
	once   sync.Once
	store  *captcha.RedisStore
	driver *base64Captcha.DriverDigit
)

func NewGetCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCaptchaLogic {
	return &GetCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCaptchaLogic) GetCaptcha() (resp *types.CaptchaInfoResp, err error) {
	captchaStore, digit := GetCaptchaStore(l.svcCtx.Config, l.svcCtx.Redis)
	gen := base64Captcha.NewCaptcha(digit, captchaStore)

	id, b64s, err := gen.Generate()

	if err != nil {
		logx.Errorw("fail to generate captcha", logx.Field("detail", err.Error()))
		return nil, errorx.NewApiError(http.StatusInternalServerError, message.CaptchaCreateFail)
	}

	return &types.CaptchaInfoResp{
		CaptchaId: id,
		ImgPath:   b64s,
	}, nil
}

func GetCaptchaStore(c config.Config, r *redis.Redis) (*captcha.RedisStore, *base64Captcha.DriverDigit) {
	once.Do(func() {
		driver = base64Captcha.NewDriverDigit(
			c.Captcha.ImgHeight,
			c.Captcha.ImgWidth,
			c.Captcha.KeyLong,
			0.7,
			80,
		)
		store = captcha.NewCaptchaStoreWithRedis(r)
		logx.Info("once: init captcha store")
	})
	return store, driver
}
