package role

import (
	"context"
	"github.com/zeromicro/go-zero/core/errorx"
	"net/http"
	"slash-admin/app/admin/ent/sysrole"
	pType "slash-admin/pkg/types"

	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleLogic {
	return &UpdateRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateRoleLogic) UpdateRole(req *types.UpdateRoleReq) (resp *types.SimpleMsgResp, err error) {
	// check if the role exists
	exist, err := l.svcCtx.EntClient.SysRole.Query().Where(sysrole.ID(req.ID)).Exist(l.ctx)

	if err != nil {
		logx.Errorw(errorx.DatabaseError, logx.Field("detail", err.Error()))
		return nil, errorx.NewApiError(http.StatusInternalServerError, errorx.DatabaseError)
	}

	if !exist {
		l.Errorf("role not found, role id: %d", req.ID)
		return nil, errorx.NewApiError(http.StatusBadRequest, errorx.TargetNotExist)
	}
	data, err := l.svcCtx.EntClient.SysRole.UpdateOneID(req.ID).
		SetNillableName(req.Name).
		SetNillableValue(req.Value).
		SetNillableDefaultRouter(req.DefaultRouter).
		SetNillableStatus((*pType.Status)(req.Status)).
		SetNillableRemark(req.Remark).
		SetNillableOrderNo(req.OrderNo).
		Save(l.ctx)

	if err != nil {
		logx.Errorw(errorx.DatabaseError, logx.Field("detail", err.Error()))
		return nil, errorx.NewApiError(http.StatusInternalServerError, errorx.UpdateFailed)
	}

	err = UpdateRoleInfoInRedis(l.ctx, l.svcCtx.Redis, l.svcCtx.EntClient)

	if err != nil {
		logx.Errorw("fail to update the role info in Redis", logx.Field("detail", err.Error()))
		return nil, err
	}
	logx.Infow("update role successfully", logx.Field("detail", data))
	return &types.SimpleMsgResp{Msg: errorx.UpdateSuccess}, nil
}
