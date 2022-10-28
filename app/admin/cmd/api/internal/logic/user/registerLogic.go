package user

import (
	"context"
	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/errorx"
	"slash-admin/app/admin/cmd/api/internal/logic/captcha"
	"slash-admin/pkg/message"
	"slash-admin/pkg/utils"

	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.SimpleMsgResp, err error) {
	store, _ := captcha.GetCaptchaStore(l.svcCtx.Config, l.svcCtx.RedisClient)
	ok := store.Verify(req.CaptchaId, req.Captcha, true)
	if !ok {
		return nil, errorx.NewApiBadRequestError(message.WrongCaptcha)
	}
	save, err := l.svcCtx.EntClient.SysUser.
		Create().
		SetUUID(uuid.NewString()).
		SetUsername(req.Username).
		SetNickname(req.Username).
		SetPassword(utils.BcryptEncrypt(req.Password)).
		SetEmail(req.Email).
		Save(l.ctx)

	if err != nil {
		l.Errorf("register logic: create user err: %s", err.Error())
		return nil, errorx.NewApiInternalServerError(errorx.DatabaseError)
	}

	l.Infow("register logic: create user success", logx.Field("user", save))

	return &types.SimpleMsgResp{
		Msg: errorx.Success,
	}, nil
}
