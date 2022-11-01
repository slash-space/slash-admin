package token

import (
	"context"
	"github.com/zeromicro/go-zero/core/errorx"
	"slash-admin/app/admin/cmd/api/internal/globalkey"
	"slash-admin/app/admin/ent/systoken"
	"slash-admin/pkg/message"
	"time"

	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogoutLogic) Logout(req *types.UUIDReq) (resp *types.SimpleMsgResp, err error) {
	UUID := l.ctx.Value(globalkey.JWTUserId).(string)

	_, err = l.svcCtx.EntClient.SysToken.Update().Where(systoken.UUIDEQ(UUID)).SetStatus(0).Save(l.ctx)

	if err != nil {
		logx.Errorw("logout: set token status to disabled", logx.Field("detail", err.Error()))
		return nil, errorx.NewApiBadRequestError(message.DatabaseError)
	}

	allTokens, err := l.svcCtx.EntClient.SysToken.Query().
		Where(systoken.UUIDEQ(UUID)).
		Where(systoken.StatusEQ(0)).
		Where(systoken.ExpiredAtGT(time.Now())).
		All(l.ctx)

	if err != nil {
		logx.Errorw("logout: get all not expired tokens", logx.Field("detail", err.Error()))
		return nil, errorx.NewApiBadRequestError(message.DatabaseError)
	}

	for _, v := range allTokens {
		err := l.svcCtx.Redis.Set(globalkey.GetBlackListToken(v.Token), "1")
		if err != nil {
			logx.Errorw("logout: set token to redis", logx.Field("detail", err.Error()))
			return nil, errorx.NewApiBadRequestError(message.RedisError)
		}
	}

	return &types.SimpleMsgResp{Msg: errorx.UpdateSuccess}, nil
}
