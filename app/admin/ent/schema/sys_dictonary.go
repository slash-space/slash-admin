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

type SysDictionary struct {
	ent.Schema
}

func (SysDictionary) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("title").Comment("the title shown in the UI"),
		field.String("name").Comment("the name of dictionary for search"),

		field.Uint8("status").
			SchemaType(map[string]string{dialect.MySQL: "tinyint unsigned"}).
			GoType(types.Status(0)).
			Optional().
			Comment("0=禁用 1=开启"),

		field.String("desc").Comment("the descriptions of dictionary"),
		field.Time("created_at").Immutable().Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Optional().Nillable(),
	}
}

func (SysDictionary) Edges() []ent.Edge {
	return nil
}

func (SysDictionary) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("deleted_at"),
	}
}

func (SysDictionary) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_dictionary"},
	}
}