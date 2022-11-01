package token

import (
	"context"
	"github.com/zeromicro/go-zero/core/errorx"
	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"
	"slash-admin/pkg/message"
	pType "slash-admin/pkg/types"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTokenLogic {
	return &CreateTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateTokenLogic) CreateToken(req *types.CreateTokenReq) (resp *types.SimpleMsgResp, err error) {

	_, err = l.svcCtx.EntClient.SysToken.Create().
		SetUUID(req.UUID).
		SetToken(req.Token).
		SetStatus(pType.Status(req.Status)).
		SetSource(req.Source).
		SetExpiredAt(time.Unix(req.ExpireAt, 0)).
		Save(l.ctx)

	if err != nil {
		logx.Errorw(message.DatabaseError, logx.Field("detail", err.Error()))
		return nil, errorx.NewApiInternalServerError(err.Error())
	}
	return &types.SimpleMsgResp{Msg: errorx.CreateSuccess}, nil
}
