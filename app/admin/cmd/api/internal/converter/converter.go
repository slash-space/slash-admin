package converter

import (
	"slash-admin/app/admin/cmd/api/internal/types"
	"slash-admin/app/admin/ent"
	pType "slash-admin/pkg/types"
	"time"
)

// goverter:converter
// goverter:extend StatusToUint8
// goverter:extend TimeToUnixMilli
type Converter interface {
	ConvertPagination(input *ent.PageDetails) (output *types.Pagination)

	ConvertSysUser(input *ent.SysUser) (output *types.UserInfo)
	ConvertSysUserList(input []*ent.SysUser) (output []*types.UserInfo)

	ConvertSysRoleToRoleInfo(input *ent.SysRole) (output *types.RoleInfo)

	ConvertSysRoleToRoleInfoList(input []*ent.SysRole) (output []*types.RoleInfo)
}

func StatusToUint8(v pType.Status) uint8 {
	return uint8(v)
}

func TimeToUnixMilli(v time.Time) int64 {
	return v.UnixMilli()
}
