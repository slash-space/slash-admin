package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type BsUser struct {
	ent.Schema
}

func (BsUser) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
	}
}

// Fields of the BsUser.
func (BsUser) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Annotations(entproto.Field(1)),
		field.String("username").
			Unique().
			Comment("登录名").
			Annotations(entproto.Field(2)),

		field.Uint64("points").
			Annotations(entproto.Field(3)),

		field.Enum("status").
			Values("pending", "active").
			Annotations(
				entproto.Field(4),
				entproto.Enum(map[string]int32{
					"pending": 1,
					"active":  2,
				}),
			),
	}
}
