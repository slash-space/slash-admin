package user

import (
	"context"
	"github.com/zeromicro/go-zero/core/errorx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"slash-admin/app/admin/cmd/api/internal/globalkey"
	"slash-admin/app/admin/cmd/api/internal/logic/captcha"
	"slash-admin/app/admin/ent"
	"slash-admin/app/admin/ent/sysuser"
	"slash-admin/pkg/message"
	"slash-admin/pkg/utils"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	store, _ := captcha.GetCaptchaStore(l.svcCtx.Config, l.svcCtx.Redis)

	if ok := store.Verify(req.CaptchaId, req.Captcha, true); !ok {
		return nil, errorx.NewApiBadRequestError(message.WrongCaptcha)
	}

	user, err := l.svcCtx.EntClient.SysUser.
		Query().
		Where(sysuser.UsernameEQ(req.Username)).
		WithRole().
		First(l.ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			l.Errorf("login logic: user not found: %s", err.Error())
			return nil, errorx.NewApiBadRequestError(message.UserNotExists)
		}
		l.Errorf("login logic: query user err: %s", err.Error())
		return nil, errorx.NewApiInternalServerError(errorx.DatabaseError)
	}

	if ok := utils.BcryptCheck(req.Password, user.Password); !ok {
		logx.Errorw("wrong password", logx.Field("detail", req))
		return nil, status.Error(codes.InvalidArgument, message.WrongUsernameOrPassword)
	}

	roleMeta := &types.RoleMetaInfo{
		RoleName: "",
		Value:    "",
	}
	roleId := globalkey.RoleMemberID

	if user.Edges.Role != nil {
		roleId = int(user.Edges.Role.ID)
		roleMeta.RoleName = user.Edges.Role.Name
		roleMeta.Value = user.Edges.Role.Value
	}

	// 生成token
	token, err := GetJwtToken(
		l.svcCtx.Config.Auth.AccessSecret,
		user.UUID,
		time.Now().Unix(),
		l.svcCtx.Config.Auth.AccessExpire,
		roleId,
	)

	if err != nil {
		l.Errorf("login logic: generate token err: %s", err.Error())
		return nil, errorx.NewApiInternalServerError(message.GenerateTokenError)
	}

	// add token into database
	expiredAt := time.Now().Add(time.Second * 259200)

	err = l.svcCtx.EntClient.SysToken.Create().
		SetUUID(user.UUID).
		SetToken(token).
		SetSource("core").
		SetStatus(0).
		SetExpiredAt(expiredAt).
		Exec(l.ctx)

	if err != nil {
		l.Errorf("login logic: save token err: %s", err.Error())
		return nil, errorx.NewApiInternalServerError(message.GenerateTokenError)
	}

	l.Infow("log in successfully", logx.Field("UUID", user.UUID))

	return &types.LoginResp{
		User:      l.svcCtx.Converter.ConvertSysUser(user),
		Role:      roleMeta,
		Token:     token,
		ExpiredAt: expiredAt.Unix(),
	}, nil

}

func GetJwtToken(secretKey, uuid string, iat, seconds int64, roleId int) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims[globalkey.JWTUserId] = uuid
	claims[globalkey.JWTRoleId] = roleId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func GetSysUserJwtToken(user *ent.SysUser, secretKey string, iat, seconds int64) (string, error) {
	roleId := globalkey.RoleMemberID
	if user.Edges.Role != nil {
		roleId = int(user.Edges.Role.ID)
	}
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims[globalkey.JWTUserId] = user.UUID
	claims[globalkey.JWTRoleId] = roleId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
