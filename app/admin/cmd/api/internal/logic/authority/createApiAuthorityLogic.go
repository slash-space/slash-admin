package authority

import (
	"context"
	"github.com/zeromicro/go-zero/core/errorx"
	"net/http"
	"strconv"

	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateApiAuthorityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateApiAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateApiAuthorityLogic {
	return &CreateApiAuthorityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateApiAuthorityLogic) CreateApiAuthority(req *types.CreateApiAuthorityReq) (resp *types.SimpleMsgResp, err error) {
	// clear old policies
	roleIdString := strconv.Itoa(int(req.RoleId))
	var oldPolicies [][]string
	oldPolicies = l.svcCtx.Casbin.GetFilteredPolicy(0, roleIdString)
	if len(oldPolicies) != 0 {
		removeResult, err := l.svcCtx.Casbin.RemoveFilteredPolicy(0, roleIdString)
		if err != nil {
			return nil, &errorx.ApiError{
				Code: http.StatusInternalServerError,
				Msg:  err.Error(),
			}
		}
		if !removeResult {
			return nil, errorx.NewApiError(http.StatusInternalServerError, "cannot clear old policies")
		}
	}
	// add new policies
	var policies [][]string
	for _, v := range req.Data {
		policies = append(policies, []string{roleIdString, v.Path, v.Method})
	}
	addResult, err := l.svcCtx.Casbin.AddPolicies(policies)
	if err != nil {
		return nil, err
	}
	if addResult {
		return &types.SimpleMsgResp{Msg: errorx.UpdateSuccess}, nil
	} else {
		return nil, errorx.NewApiError(http.StatusBadRequest, errorx.UpdateFailed)
	}
}
