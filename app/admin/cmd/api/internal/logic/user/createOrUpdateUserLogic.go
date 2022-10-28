package user

import (
	"context"
	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/errorx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"slash-admin/app/admin/ent/sysuser"
	"slash-admin/pkg/message"
	pType "slash-admin/pkg/types"
	"slash-admin/pkg/utils"

	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrUpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateOrUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrUpdateUserLogic {
	return &CreateOrUpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOrUpdateUserLogic) CreateOrUpdateUser(req *types.CreateOrUpdateUserReq) (resp *types.SimpleMsgResp, err error) {
	if req.ID != 0 {
		return l.UpdateUser(req)
	}
	return l.CreateUser(req)
}

func (l *CreateOrUpdateUserLogic) CreateUser(req *types.CreateOrUpdateUserReq) (resp *types.SimpleMsgResp, err error) {
	if req.Username == "" || req.Email == "" || req.Password == "" {
		l.Errorw("create user error", logx.Field("err", "username or email or password is empty"))
		return nil, errorx.NewApiBadRequestError(message.UserParamsError)
	}

	// check username or email is exist
	exist, err := l.svcCtx.EntClient.SysUser.Query().
		Where(sysuser.Or(sysuser.Username(req.Username), sysuser.Email(req.Email))).
		Exist(l.ctx)

	if err != nil {
		l.Errorw("query user error", logx.Field("err", err))
		return nil, errorx.NewApiBadRequestError(errorx.DatabaseError)
	}

	if exist {
		logx.Errorw("username or email address had been used",
			logx.Field("username", req.Username),
			logx.Field("email", req.Email),
		)
		return nil, status.Error(codes.InvalidArgument, message.UserAlreadyExists)
	}

	// create user
	err = l.svcCtx.EntClient.SysUser.Create().
		SetUUID(uuid.NewString()).
		SetUsername(req.Username).
		SetNickname(req.Nickname).
		SetPassword(req.Password).
		SetEmail(req.Email).
		SetRoleID(req.RoleID).
		SetAvatar(req.Avatar).
		SetMobile(req.Mobile).
		SetStatus(pType.StatusNormal).
		Exec(l.ctx)

	if err != nil {
		l.Errorw("create user error", logx.Field("err", err))
		return nil, errorx.NewApiBadRequestError(errorx.DatabaseError)
	}

	return &types.SimpleMsgResp{Msg: errorx.Success}, nil
}

func (l *CreateOrUpdateUserLogic) UpdateUser(req *types.CreateOrUpdateUserReq) (resp *types.SimpleMsgResp, err error) {

	updateOne := l.svcCtx.EntClient.SysUser.UpdateOneID(req.ID)

	if req.Username != "" {
		updateOne.SetUsername(req.Username)
	}

	if req.Nickname != "" {
		updateOne.SetNickname(req.Nickname)
	}

	if req.Password != "" {
		updateOne.SetPassword(utils.BcryptEncrypt(req.Password))
	}

	if req.Email != "" {
		updateOne.SetEmail(req.Email)
	}

	if req.RoleID != 0 {
		updateOne.SetRoleID(req.RoleID)
	}

	if req.Avatar != "" {
		updateOne.SetAvatar(req.Avatar)
	}

	if req.Mobile != "" {
		updateOne.SetMobile(req.Mobile)
	}

	if req.Status != nil {
		updateOne.SetStatus(pType.Status(*req.Status))
	}

	err = updateOne.Exec(l.ctx)

	if err != nil {
		l.Errorw("update user error", logx.Field("err", err))
		return nil, errorx.NewApiBadRequestError(errorx.DatabaseError)
	}

	logx.Infow("update user successfully", logx.Field("detail", req))

	return &types.SimpleMsgResp{Msg: errorx.Success}, nil
}
