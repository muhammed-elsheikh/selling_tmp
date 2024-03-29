package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id"),
		field.Int("product_id"),
		field.Int("quantity"),
		field.Float("total"),
		field.Time("order_date").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Unique().Ref("orders"),
		edge.From("product", Product.Type).Required().Field("product_id").Unique().Ref("order"),
	}
}
