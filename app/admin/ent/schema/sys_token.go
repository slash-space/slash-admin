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

type SysToken struct {
	ent.Schema
}

func (SysToken) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("uuid").Comment(" 用户的UUID"),
		field.String("token").Comment(" Token 字符串"),
		field.String("source").Comment(" Token 来源 （本地为core, 第三方如github等）"),

		field.Uint8("status").
			SchemaType(map[string]string{dialect.MySQL: "tinyint unsigned"}).
			GoType(types.Status(0)).
			Optional().
			Comment("0=禁用 1=正常"),

		field.Time("expired_at").Comment(" 过期时间"),
		field.Time("created_at").Immutable().Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Optional().Nillable(),
	}
}

func (SysToken) Edges() []ent.Edge {
	return nil
}

func (SysToken) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("uuid"),
		index.Fields("deleted_at"),
		index.Fields("expired_at"),
	}
}

func (SysToken) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_token"},
	}
}
