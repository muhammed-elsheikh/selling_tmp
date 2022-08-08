package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(70),
		field.Int("age").Positive(),
		field.String("email").Unique(),
		field.String("phone").Optional(),
		field.String("national_id").Optional(),
		field.String("local_address").Optional(),
		field.String("card_id").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("card", Card.Type).Unique(),
		edge.To("orders", Order.Type),
	}
}
