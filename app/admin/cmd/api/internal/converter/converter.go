package converter

import (
	"slash-admin/app/admin/cmd/api/internal/types"
	"slash-admin/app/admin/ent"
	pType "slash-admin/pkg/types"
)

//go:generate go run github.com/jmattheis/goverter/cmd/goverter slash-admin/app/admin/cmd/api/internal/converter

// goverter:converter
// goverter:extend StatusToUint8
type Converter interface {
	// goverter:matchIgnoreCase
	ConvertPagination(input *ent.PageDetails) (output *types.Pagination)

	// goverter:matchIgnoreCase
	ConvertSysRoleToRoleInfo(input *ent.SysRole) (output *types.RoleInfo)
	ConvertSysRoleToRoleInfoList(input []*ent.SysRole) (output []*types.RoleInfo)
}

func StatusToUint8(v pType.Status) uint8 {
	return uint8(v)
}
