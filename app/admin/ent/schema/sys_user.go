package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"slash-admin/pkg/types"
	"time"
)

type SysUser struct {
	ent.Schema
}

func (SysUser) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("uuid").Comment("用户 UUID"),
		field.String("username").Unique().Comment("登录名"),
		field.String("password").Comment("密码"),
		field.String("nickname").Unique().Comment("昵称"),
		field.String("side_mode").Default("dark").Comment("布局方式"),
		field.String("avatar").Comment("头像路径"),
		field.String("base_color").Default("#fff").Comment("后台页面色调"),
		field.String("active_color").Default("#1890ff").Comment("当前激活的颜色设定"),
		field.Uint32("role_id").Default(2).Comment("角色ID"),
		field.String("mobile").Comment("手机号"),
		field.String("email").Comment("邮箱号"),

		field.Uint8("status").
			SchemaType(map[string]string{dialect.MySQL: "tinyint unsigned"}).
			GoType(types.Status(0)).
			Optional().
			Comment("0=正常 1=禁用"),

		field.Time("created_at").Immutable().Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Optional().Nillable(),
	}
}

func (SysUser) Edges() []ent.Edge {
	return nil
}

func (SysUser) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("deleted_at"),
	}
}

func (SysUser) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_user"},
	}
}
