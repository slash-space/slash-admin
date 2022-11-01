package token

import (
	"context"
	"github.com/zeromicro/go-zero/core/errorx"
	"slash-admin/app/admin/cmd/api/internal/globalkey"
	"slash-admin/app/admin/ent"
	"slash-admin/app/admin/ent/systoken"
	"slash-admin/pkg/message"
	pType "slash-admin/pkg/types"
	"time"

	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetTokenStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetTokenStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetTokenStatusLogic {
	return &SetTokenStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetTokenStatusLogic) SetTokenStatus(req *types.SetBooleanStatusReq) (resp *types.SimpleMsgResp, err error) {
	err = l.svcCtx.EntClient.SysToken.UpdateOneID(req.ID).
		SetStatus(pType.Status(req.Status)).
		Exec(l.ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			logx.Errorw("Update token status failed, please check the token id", logx.Field("TokenId", req.ID))
			return nil, errorx.NewApiBadRequestError(errorx.UpdateFailed)
		}
		l.Errorw("Update token status failed", logx.Field("TokenId", req.ID), logx.Field("err", err))
		return nil, errorx.NewApiInternalServerError(message.DatabaseError)
	}

	err = l.UpdateTokenInRedis(req.ID, pType.Status(req.Status))

	if err != nil {
		l.Errorf("Update token status failed, update token in redis failed", logx.Field("TokenId", req.ID), logx.Field("err", err))
		return nil, errorx.NewApiInternalServerError(message.RedisError)
	}

	return &types.SimpleMsgResp{Msg: errorx.UpdateSuccess}, nil
}

func (l *SetTokenStatusLogic) UpdateTokenInRedis(id uint64, status pType.Status) error {
	// add into redis
	tokenData, err := l.svcCtx.EntClient.SysToken.Query().Where(systoken.ID(id)).First(l.ctx)

	if err != nil {
		return err
	}

	if status == 0 {
		expired := tokenData.ExpiredAt.Unix() - time.Now().Unix()
		err := l.svcCtx.Redis.Setex(globalkey.GetBlackListToken(tokenData.Token), "1", int(expired))
		if err != nil {
			return err
		}
	} else if status == 1 {
		_, err := l.svcCtx.Redis.Del(globalkey.GetBlackListToken(tokenData.Token))
		if err != nil {
			return err
		}
	}

	return nil
}
