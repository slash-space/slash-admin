package authority

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"
	"slash-admin/app/admin/ent"
	"slash-admin/app/admin/ent/sysmenu"
	"slash-admin/pkg/message"
	"slash-admin/pkg/utils"
	"strings"
)

type CreateMenuAuthorityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateMenuAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMenuAuthorityLogic {
	return &CreateMenuAuthorityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateMenuAuthorityLogic) CreateMenuAuthority(req *types.CreateMenuAuthorityReq) (resp *types.SimpleMsgResp, err error) {

	err = utils.WithTx(l.ctx, l.svcCtx.EntClient, func(tx *ent.Tx) error {
		deleteString := fmt.Sprintf("DELETE FROM %s where role_id = %d", sysmenu.RolesTable, req.RoleId)
		_, err := tx.ExecContext(l.ctx, deleteString)
		if err != nil {
			return err
		}

		var insertString strings.Builder
		insertString.WriteString(fmt.Sprintf("insert into %s values ", sysmenu.RolesTable))
		for i := 0; i < len(req.MenuIds); i++ {
			if i != len(req.MenuIds)-1 {
				insertString.WriteString(fmt.Sprintf("(%d, %d),", req.MenuIds[i], req.RoleId))
			} else {
				insertString.WriteString(fmt.Sprintf("(%d, %d);", req.MenuIds[i], req.RoleId))
			}
		}
		_, err = tx.ExecContext(l.ctx, insertString.String())
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		l.Errorf("CreateMenuAuthority error: %v", err)
		return nil, errorx.NewApiInternalServerError(message.DatabaseError)
	}

	logx.Infow(errorx.UpdateSuccess, logx.Field("details", req.MenuIds), logx.Field("role_id", req.RoleId))

	return &types.SimpleMsgResp{Msg: errorx.UpdateSuccess}, nil
}
