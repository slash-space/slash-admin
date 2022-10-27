package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

type SysMenuParam struct {
	ent.Schema
}

func (SysMenuParam) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("menu_id").Comment("menu id"),
		field.String("type").Comment("参数类型"),
		field.String("key").Comment("参数键"),
		field.String("value").Comment("参数值"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Optional().Nillable(),
	}
}

func (SysMenuParam) Edges() []ent.Edge {
	return nil
}

func (SysMenuParam) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("deleted_at"),
	}
}

func (SysMenuParam) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_menu_param"},
	}
}
