package role

import (
	"context"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"slash-admin/app/admin/cmd/api/internal/globalkey"
	"slash-admin/app/admin/ent"
	"slash-admin/pkg/message"
	pType "slash-admin/pkg/types"
	"strconv"

	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRoleLogic {
	return &CreateRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateRoleLogic) CreateRole(req *types.CreateRoleReq) (resp *types.SimpleMsgResp, err error) {
	data, err := l.svcCtx.EntClient.SysRole.Create().
		SetValue(req.Value).
		SetName(req.Name).
		SetDefaultRouter(req.DefaultRouter).
		SetStatus(pType.Status(req.Status)).
		SetRemark(req.Remark).
		SetOrderNo(req.OrderNo).
		Save(l.ctx)

	if err != nil {
		logx.Errorw(message.DatabaseError, logx.Field("detail", err.Error()))
		return nil, errorx.NewApiInternalServerError(message.DatabaseError)
	}

	err = UpdateRoleInfoInRedis(l.ctx, l.svcCtx.Redis, l.svcCtx.EntClient)

	if err != nil {
		logx.Errorw("fail to update the role info in Redis", logx.Field("detail", err.Error()))
		return nil, err
	}

	logx.Infow("create role successfully", logx.Field("detail", data))
	return &types.SimpleMsgResp{Msg: errorx.CreateSuccess}, nil
}

func UpdateRoleInfoInRedis(ctx context.Context, redis *redis.Redis, entClient *ent.Client) error {
	roleData, err := entClient.SysRole.Query().All(ctx)

	if err != nil {
		logx.Errorw(message.DatabaseError, logx.Field("detail", err.Error()))
		return errorx.NewApiInternalServerError(message.DatabaseError)
	}

	for _, v := range roleData {
		err := redis.Hset(globalkey.RoleList, globalkey.GetRoleListID(v.ID), v.Name)
		err = redis.Hset(globalkey.RoleList, globalkey.GetRoleListValue(v.ID), v.Value)
		err = redis.Hset(globalkey.RoleList, globalkey.GetRoleListStatus(v.ID), strconv.Itoa(int(v.Status)))
		if err != nil {
			return errorx.NewApiInternalServerError(message.RedisError)
		}
	}

	return nil
}
