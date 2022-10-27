package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

type SysOauthProvider struct {
	ent.Schema
}

func (SysOauthProvider) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("name").Comment("the provider's name"),
		field.String("client_id").Comment("the client id"),
		field.String("client_secret").Comment("the client secret"),
		field.String("redirect_url").Comment("the redirect url"),
		field.String("scopes").Comment("the scopes"),
		field.String("auth_url").Comment("the auth url of the provider"),
		field.String("token_url").Comment("the token url of the provider"),
		field.Uint8("auth_style").Comment("the auth style, 0: auto detect 1: third party log in 2: log in with username and password"),
		field.String("info_url").Comment("the URL to request user information by token"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Optional().Nillable(),
	}
}

func (SysOauthProvider) Edges() []ent.Edge {
	return nil
}

func (SysOauthProvider) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("deleted_at"),
	}
}

func (SysOauthProvider) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_oauth_provider"},
	}
}
