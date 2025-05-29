package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Basket holds the schema definition for the Basket entity.
type Basket struct {
	ent.Schema
}

// Fields of the Basket.
func (Basket) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			Unique().
			Immutable().
			Positive(),

		field.Int("user_id").
			Positive(),

		field.Int("product_id").
			Positive(),

		field.Int("quantity").
			Positive(),

		field.Float("price").
			Positive(),

		field.Time("added_at").
			Default(time.Now).
			Immutable(),

		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Basket.
func (Basket) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("baskets").
			Field("user_id").
			Required().
			Unique(),
	}
}

// Indexes for the Basket.
func (Basket) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
		index.Fields("user_id", "product_id").
			Unique(),
	}
}
