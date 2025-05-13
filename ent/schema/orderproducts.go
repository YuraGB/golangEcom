package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
)

// OrderProducts holds the schema definition for the OrderProducts entity.
type OrderProducts struct {
	ent.Schema
}

// Fields of the OrderProducts.
func (OrderProducts) Fields() []ent.Field {
	return []ent.Field{
		field.Int("quantity").Positive(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the OrderProducts.
func (OrderProducts) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("order", Order.Type).
			Ref("order_products").
			Unique().
			Required(),

		edge.From("product", Product.Type).
			Ref("order_products").
			Unique().
			Required(),
	}
}
