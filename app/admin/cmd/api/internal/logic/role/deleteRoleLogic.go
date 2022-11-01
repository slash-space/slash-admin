package role

import (
	"context"
	"github.com/zeromicro/go-zero/core/errorx"
	"net/http"
	"slash-admin/app/admin/ent/sysuser"
	"slash-admin/pkg/message"

	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRoleLogic {
	return &DeleteRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteRoleLogic) DeleteRole(req *types.IDReq) (resp *types.SimpleMsgResp, err error) {
	// check users has this role
	exist, err := l.svcCtx.EntClient.SysUser.Query().Where(sysuser.RoleIDEQ(req.ID)).Exist(l.ctx)

	if err != nil {
		l.Errorf("check role exist error: %v", err)
		return nil, errorx.NewApiInternalServerError(message.DatabaseError)
	}

	if exist {
		logx.Errorw("delete role failed, please check the users who belongs to the role had been deleted",
			logx.Field("roleId", req.ID))
		return nil, errorx.NewApiError(http.StatusBadRequest, message.UserExists)
	}

	err = l.svcCtx.EntClient.SysRole.DeleteOneID(req.ID).Exec(l.ctx)

	if err != nil {
		l.Errorf("SysRole delete error: %v", err)
		return nil, errorx.NewApiInternalServerError(message.DatabaseError)
	}

	l.Infow("delete role successfully", logx.Field("roleId", req.ID))

	return &types.SimpleMsgResp{
		Msg: errorx.DeleteSuccess,
	}, nil
}
