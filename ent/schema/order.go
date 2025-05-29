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
		field.Enum("paymentType").Values("CASH", "CARD"),
		field.String("address").NotEmpty(),
		field.String("city").NotEmpty(),
		field.String("state").NotEmpty(),
		field.String("zip").NotEmpty(),
		field.Int("user_id"),
		field.Float("total_price").
			Positive().
			Comment("Total price of the order, must be positive"),
		field.Enum("status").
			Values("NEW", "IN_PROGRESS", "COMPLETED", "CANCELED").
			Comment("Status of the order, can be NEW, IN_PROGRESS, COMPLETED, or CANCELED"),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("order_products", OrderProducts.Type),
		edge.From("user", User.Type).
			Ref("orders").
			Field("user_id").
			Unique().
			Required(),
	}
}
