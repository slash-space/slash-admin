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

type SysDictionaryDetail struct {
	ent.Schema
}

func (SysDictionaryDetail) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("title").Comment("the title shown in the UI"),
		field.String("key").Comment("key"),
		field.String("value").Comment("value"),

		field.Uint8("status").
			SchemaType(map[string]string{dialect.MySQL: "tinyint unsigned"}).
			GoType(types.Status(0)).
			Optional().
			Comment("0=禁用 1=开启"),

		field.String("remark").Comment("备注"),
		field.Uint32("order_no").Default(0).Comment("排序编号"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Optional().Nillable(),
	}
}

func (SysDictionaryDetail) Edges() []ent.Edge {
	return nil
}

func (SysDictionaryDetail) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("deleted_at"),
	}
}

func (SysDictionaryDetail) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_dictionary_detail"},
	}
}
