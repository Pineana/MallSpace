package schema

import (
	"entgo.io/ent"
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
		field.String("order_id").Unique(),
		field.Enum("status").Values(
			"UNPAY", "SUCCESS", "CLOSE", "FAIL",
		),
		field.String("tradeID"),
		field.Float32("order_mount"),
		field.Time("pay_at").Optional(),
		field.Time("create_at"),
		field.Time("update_at").Optional(),
		field.Time("delete_at").Optional(),
		field.Bool("is_deleted").Default(false),
	}
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return nil
}
