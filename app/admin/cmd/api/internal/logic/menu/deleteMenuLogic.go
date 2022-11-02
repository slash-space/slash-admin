package menu

import (
	"context"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"
	"slash-admin/app/admin/ent"
	"slash-admin/app/admin/ent/sysmenu"
	"slash-admin/pkg/message"
)

type DeleteMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMenuLogic {
	return &DeleteMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteMenuLogic) DeleteMenu(req *types.IDReq) (resp *types.SimpleMsgResp, err error) {

	exist, err := l.svcCtx.EntClient.SysMenu.Query().Where(sysmenu.ParentID(req.ID)).Exist(l.ctx)

	if err != nil {
		l.Errorw("delete menu error", logx.Field("id", req.ID), logx.Field("error", err))
		return nil, errorx.NewApiInternalServerError(message.DatabaseError)
	}

	if exist {
		l.Errorw("Delete Menu failed, check the id", logx.Field("menuId", req.ID))
		return nil, errorx.NewApiBadRequestError(message.ChildrenExistError)
	}

	err = l.svcCtx.EntClient.SysMenu.DeleteOneID(req.ID).Exec(l.ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			l.Errorw("Delete Menu failed, check the id", logx.Field("menuParamId", req.ID))
			return nil, errorx.NewApiBadRequestError(message.TargetNotExist)
		}
		l.Errorw("delete menu error", logx.Field("id", req.ID), logx.Field("error", err))
		return nil, errorx.NewApiInternalServerError(message.DatabaseError)
	}

	l.Infow("delete menu successfully", logx.Field("menuId", req.ID))

	return &types.SimpleMsgResp{Msg: message.DeleteSuccess}, nil
}
