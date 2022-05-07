package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

// Fields of the Account.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).Unique(),
		field.String("access_token").SchemaType(map[string]string{
			dialect.Postgres: "text",
		}),
		field.String("refresh_token").SchemaType(map[string]string{
			dialect.Postgres: "text",
		}),
		field.String("token_type"),
		field.Time("expires_at"),
		field.String("providerAccountId"),
		field.String("scope"),
		field.String("provider"),
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("accounts").Unique(),
	}
}

func (Account) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("provider", "providerAccountId").
			Unique(),
	}
}
