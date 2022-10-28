package role

import (
	"context"
	"github.com/zeromicro/go-zero/core/errorx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"slash-admin/app/admin/cmd/api/internal/globalkey"
	"slash-admin/app/admin/ent"
	"slash-admin/app/admin/ent/sysrole"
	"strconv"

	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrUpdateRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateOrUpdateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrUpdateRoleLogic {
	return &CreateOrUpdateRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOrUpdateRoleLogic) CreateOrUpdateRole(req *types.RoleInfo) (resp *types.SimpleMsgResp, err error) {
	if req.ID == 0 {
		output := l.svcCtx.Converter.ConvertRoleInfoToCreateSysRoleInput(req)
		data, err := l.svcCtx.EntClient.SysRole.Create().Copy(output).Save(l.ctx)

		if err != nil {
			logx.Errorw(errorx.DatabaseError, logx.Field("detail", err.Error()))
			return nil, errorx.NewApiError(http.StatusInternalServerError, errorx.DatabaseError)
		}

		err = l.UpdateRoleInfoInRedis()
		if err != nil {
			logx.Errorw("fail to update the role info in Redis", logx.Field("detail", err.Error()))
			return nil, err
		}

		logx.Infow("create role successfully", logx.Field("detail", data))
		return &types.SimpleMsgResp{Msg: errorx.CreateSuccess}, nil
	} else {
		// check if the role exists
		exist, err := l.svcCtx.EntClient.SysRole.Query().Where(sysrole.ID(req.ID)).Exist(l.ctx)
		if err != nil {
			logx.Errorw(errorx.DatabaseError, logx.Field("detail", err.Error()))
			return nil, errorx.NewApiError(http.StatusInternalServerError, errorx.DatabaseError)
		}
		if !exist {
			return nil, errorx.NewApiError(http.StatusBadRequest, errorx.TargetNotExist)
		}

		output := l.svcCtx.Converter.ConvertRoleInfoToUpdateSysRoleInput(req)
		data, err := l.svcCtx.EntClient.SysRole.UpdateOneID(req.ID).Copy(output, ent.JsonbReplace).Save(l.ctx)

		if err != nil {
			logx.Errorw(errorx.DatabaseError, logx.Field("detail", err.Error()))
			return nil, errorx.NewApiError(http.StatusInternalServerError, errorx.UpdateFailed)
		}
		logx.Infow("update role successfully", logx.Field("detail", data))
		return &types.SimpleMsgResp{Msg: errorx.UpdateSuccess}, nil
	}
}

func (l *CreateOrUpdateRoleLogic) UpdateRoleInfoInRedis() error {
	roleData, err := l.svcCtx.EntClient.SysRole.Query().All(l.ctx)

	if err != nil {
		logx.Errorw(errorx.DatabaseError, logx.Field("detail", err.Error()))
		return errorx.NewApiError(http.StatusInternalServerError, errorx.DatabaseError)
	}
	for _, v := range roleData {
		err := l.svcCtx.Redis.Hset(globalkey.RoleList, globalkey.GetRoleListID(v.ID), v.Name)
		err = l.svcCtx.Redis.Hset(globalkey.RoleList, globalkey.GetRoleListValue(v.ID), v.Value)
		err = l.svcCtx.Redis.Hset(globalkey.RoleList, globalkey.GetRoleListStatus(v.ID), strconv.Itoa(int(v.Status)))
		if err != nil {
			return status.Error(codes.Internal, errorx.RedisError)
		}
	}
	return nil
}
