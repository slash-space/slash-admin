package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"time"
)

type SysApi struct {
	ent.Schema
}

func (SysApi) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("path").Comment("API path"),
		field.String("description").Comment("API description"),
		field.String("api_group").Comment("API group"),
		field.String("method").Default("POST").Comment("HTTP method"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Optional().Nillable(),
	}
}

func (SysApi) Edges() []ent.Edge {
	return nil
}

func (SysApi) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_api"},
	}
}
