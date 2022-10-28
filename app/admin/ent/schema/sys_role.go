package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"slash-admin/pkg/types"
	"time"
)

type SysRole struct {
	ent.Schema
}

func (SysRole) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("name").Comment("角色名"),
		field.String("value").Unique().Comment("角色值，用于前端权限控制"),
		field.String("default_router").Default("dashboard").Comment("默认登录页面"),

		field.Uint8("status").
			SchemaType(map[string]string{dialect.MySQL: "tinyint unsigned"}).
			GoType(types.Status(0)).
			Optional().
			Default(0).
			Comment("0=开启 1=禁用"),

		field.String("remark").Comment("备注"),
		field.Uint32("order_no").Default(0).Comment("排序编号"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Optional().Nillable(),
	}
}

func (SysRole) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("menus", SysMenu.Type).Ref("roles"),
		edge.To("role", SysUser.Type).Comment("用户"),
	}
}

func (SysRole) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("deleted_at"),
	}
}

func (SysRole) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_role"},
	}
}
