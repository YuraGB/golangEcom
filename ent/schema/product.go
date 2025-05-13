package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty(),
		field.String("description").NotEmpty(),
		field.String("category").NotEmpty(),
		field.Float("price"),
		field.Float("discount_percentage"),
		field.Float("rating"),
		field.Int("stock"),
		field.String("brand").NotEmpty(),
		field.String("sku").NotEmpty(),
		field.Float("weight"),
		field.String("warranty_information").NotEmpty(),
		field.String("shipping_information").NotEmpty(),
		field.String("availability_status").NotEmpty(),
		field.String("return_policy").NotEmpty(),
		field.Int("minimum_order_quantity"),
		field.Time("created_at"),
		field.Time("updated_at"),
		field.String("barcode").NotEmpty(),
		field.String("qr_code").NotEmpty(),
		field.String("thumbnail").NotEmpty(),
		field.JSON("tags", []string{}),
		field.JSON("images", []string{}),
		field.JSON("dimensions", map[string]float64{}),
		field.JSON("reviews", []map[string]interface{}{}),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("order_products", OrderProducts.Type),
	}
}
