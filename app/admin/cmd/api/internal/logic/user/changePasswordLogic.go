package user

import (
	"context"
	"github.com/zeromicro/go-zero/core/errorx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"slash-admin/app/admin/cmd/api/internal/globalkey"
	"slash-admin/app/admin/ent"
	"slash-admin/app/admin/ent/sysuser"
	"slash-admin/pkg/message"
	"slash-admin/pkg/utils"

	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangePasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangePasswordLogic {
	return &ChangePasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangePasswordLogic) ChangePassword(req *types.ChangePasswordReq) (resp *types.SimpleMsgResp, err error) {
	uuid := l.ctx.Value(globalkey.JWTUserId).(string)

	target, err := l.svcCtx.EntClient.SysUser.Query().Where(sysuser.UUIDEQ(uuid)).First(l.ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			l.Infow("user not found", logx.Field("uuid", uuid))
			return nil, errorx.NewApiBadRequestError(message.UserNotExists)
		}
		l.Errorw("query user error", logx.Field("uuid", uuid), logx.Field("error", err))
		return nil, errorx.NewApiBadRequestError(errorx.DatabaseError)
	}

	if ok := utils.BcryptCheck(req.OldPassword, target.Password); !ok {
		l.Infow("password not match", logx.Field("uuid", uuid))
		return nil, status.Errorf(codes.InvalidArgument, message.WrongPassword)
	}

	newPassword := utils.BcryptEncrypt(req.NewPassword)

	err = l.svcCtx.EntClient.SysUser.Update().Where(sysuser.UUIDEQ(uuid)).SetPassword(newPassword).Exec(l.ctx)

	if err != nil {
		l.Errorw("update user error", logx.Field("uuid", uuid), logx.Field("error", err))
		return nil, errorx.NewApiBadRequestError(errorx.DatabaseError)
	}

	logx.Infow("change password successful", logx.Field("UUID", target.UUID))

	return &types.SimpleMsgResp{
		Msg: errorx.UpdateSuccess,
	}, nil

}
