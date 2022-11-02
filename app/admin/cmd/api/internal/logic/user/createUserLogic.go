package user

import (
	"context"
	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/errorx"
	"slash-admin/app/admin/ent/sysuser"
	"slash-admin/pkg/message"
	pType "slash-admin/pkg/types"

	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserLogic) CreateUser(req *types.CreateUserReq) (resp *types.SimpleMsgResp, err error) {
	// check username or email is exist
	exist, err := l.svcCtx.EntClient.SysUser.Query().
		Where(sysuser.Or(sysuser.Username(req.Username), sysuser.Email(req.Email))).
		Exist(l.ctx)

	if err != nil {
		l.Errorw("query user error", logx.Field("err", err))
		return nil, errorx.NewApiBadRequestError(message.DatabaseError)
	}

	if exist {
		logx.Errorw("username or email address had been used",
			logx.Field("username", req.Username),
			logx.Field("email", req.Email),
		)
		return nil, errorx.NewApiBadRequestError(message.UserAlreadyExists)
	}

	// create user
	err = l.svcCtx.EntClient.SysUser.Create().
		SetUUID(uuid.NewString()).
		SetUsername(req.Username).
		SetNickname(req.Nickname).
		SetPassword(req.Password).
		SetEmail(req.Email).
		SetNillableRoleID(req.RoleID).
		SetNillableAvatar(req.Avatar).
		SetNillableMobile(req.Mobile).
		SetStatus(pType.StatusNormal).
		Exec(l.ctx)

	if err != nil {
		l.Errorw("create user error", logx.Field("err", err))
		return nil, errorx.NewApiBadRequestError(message.DatabaseError)
	}

	return &types.SimpleMsgResp{Msg: errorx.Success}, nil
}
