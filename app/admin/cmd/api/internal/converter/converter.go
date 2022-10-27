package converter

import (
	"slash-admin/app/admin/cmd/api/internal/types"
	"slash-admin/app/admin/ent"
	pType "slash-admin/pkg/types"
)

// goverter:converter
// goverter:extend StatusToUint8
type Converter interface {
	ConvertPagination(input *ent.PageDetails) (output *types.Pagination)
	ConvertSysRoleToRoleInfo(input *ent.SysRole) (output *types.RoleInfo)

	ConvertSysRoleToRoleInfoList(input []*ent.SysRole) (output []*types.RoleInfo)

	// goverter:ignore CreatedAt UpdatedAt DeletedAt
	ConvertRoleInfoToCreateSysRoleInput(input *types.RoleInfo) (output *ent.CreateSysRoleInput)
	// goverter:ignore CreatedAt UpdatedAt DeletedAt
	ConvertRoleInfoToUpdateSysRoleInput(input *types.RoleInfo) (output *ent.UpdateSysRoleInput)
}

func StatusToUint8(v pType.Status) uint8 {
	return uint8(v)
}
