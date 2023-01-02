// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// OrdersColumns holds the columns for the "orders" table.
	OrdersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_id", Type: field.TypeInt},
		{Name: "order_id", Type: field.TypeString, Unique: true},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"UNPAY", "SUCCESS", "CLOSE", "FAIL"}},
		{Name: "trade_id", Type: field.TypeString},
		{Name: "order_mount", Type: field.TypeFloat32},
		{Name: "pay_at", Type: field.TypeTime, Nullable: true},
		{Name: "create_at", Type: field.TypeTime},
		{Name: "update_at", Type: field.TypeTime, Nullable: true},
		{Name: "delete_at", Type: field.TypeTime, Nullable: true},
		{Name: "is_deleted", Type: field.TypeBool, Default: false},
	}
	// OrdersTable holds the schema information for the "orders" table.
	OrdersTable = &schema.Table{
		Name:       "orders",
		Columns:    OrdersColumns,
		PrimaryKey: []*schema.Column{OrdersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		OrdersTable,
	}
)

func init() {
}
