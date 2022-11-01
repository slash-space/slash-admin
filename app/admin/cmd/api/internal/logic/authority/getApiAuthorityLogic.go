package authority

import (
	"context"
	"strconv"

	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetApiAuthorityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetApiAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApiAuthorityLogic {
	return &GetApiAuthorityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetApiAuthorityLogic) GetApiAuthority(req *types.IDReq) (resp *types.ApiAuthorityListResp, err error) {
	roleId := strconv.Itoa(int(req.ID))
	data := l.svcCtx.Casbin.GetFilteredPolicy(0, roleId)
	resp = &types.ApiAuthorityListResp{}
	resp.Pagination = &types.Pagination{
		Total: len(data),
	}
	for _, v := range data {
		resp.List = append(resp.List, &types.ApiAuthorityInfo{
			Path:   v[1],
			Method: v[2],
		})
	}
	return resp, nil
}
