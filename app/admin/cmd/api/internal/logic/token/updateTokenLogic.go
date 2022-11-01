package token

import (
	"context"
	"github.com/zeromicro/go-zero/core/errorx"
	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"
	"slash-admin/app/admin/ent/systoken"
	"slash-admin/pkg/message"
	pType "slash-admin/pkg/types"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTokenLogic {
	return &UpdateTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateTokenLogic) UpdateToken(req *types.UpdateTokenReq) (resp *types.SimpleMsgResp, err error) {
	exist, err := l.svcCtx.EntClient.SysToken.Query().Where(systoken.IDEQ(req.ID)).Exist(l.ctx)

	if err != nil {
		logx.Errorw(message.DatabaseError, logx.Field("detail", err.Error()))
		return nil, errorx.NewApiInternalServerError(message.DatabaseError)
	}

	if !exist {
		logx.Errorw(errorx.TargetNotExist, logx.Field("detail", req.ID))
		return nil, errorx.NewApiInternalServerError(errorx.TargetNotExist)
	}

	var expiredTime *time.Time

	if req.ExpiredAt != nil {
		t := time.Unix(*req.ExpiredAt, 0)
		expiredTime = &t
	}

	_, err = l.svcCtx.EntClient.SysToken.UpdateOneID(req.ID).
		SetNillableUUID(req.UUID).
		SetNillableToken(req.Token).
		SetNillableStatus((*pType.Status)(req.Status)).
		SetNillableSource(req.Source).
		SetNillableExpiredAt(expiredTime).
		Save(l.ctx)

	if err != nil {
		logx.Errorw(message.DatabaseError, logx.Field("detail", err.Error()))
		return nil, errorx.NewApiInternalServerError(err.Error())
	}

	return &types.SimpleMsgResp{Msg: errorx.UpdateSuccess}, nil

}
