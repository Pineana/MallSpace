// Code generated by ent, DO NOT EDIT.

package order

import (
	"fmt"
)

const (
	// Label holds the string label denoting the order type in the database.
	Label = "order"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldOrderID holds the string denoting the order_id field in the database.
	FieldOrderID = "order_id"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldTradeID holds the string denoting the tradeid field in the database.
	FieldTradeID = "trade_id"
	// FieldOrderMount holds the string denoting the order_mount field in the database.
	FieldOrderMount = "order_mount"
	// FieldPayAt holds the string denoting the pay_at field in the database.
	FieldPayAt = "pay_at"
	// FieldCreateAt holds the string denoting the create_at field in the database.
	FieldCreateAt = "create_at"
	// FieldUpdateAt holds the string denoting the update_at field in the database.
	FieldUpdateAt = "update_at"
	// FieldDeleteAt holds the string denoting the delete_at field in the database.
	FieldDeleteAt = "delete_at"
	// FieldIsDeleted holds the string denoting the is_deleted field in the database.
	FieldIsDeleted = "is_deleted"
	// Table holds the table name of the order in the database.
	Table = "orders"
)

// Columns holds all SQL columns for order fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldOrderID,
	FieldStatus,
	FieldTradeID,
	FieldOrderMount,
	FieldPayAt,
	FieldCreateAt,
	FieldUpdateAt,
	FieldDeleteAt,
	FieldIsDeleted,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultIsDeleted holds the default value on creation for the "is_deleted" field.
	DefaultIsDeleted bool
)

// Status defines the type for the "status" enum field.
type Status string

// Status values.
const (
	StatusUNPAY   Status = "UNPAY"
	StatusSUCCESS Status = "SUCCESS"
	StatusCLOSE   Status = "CLOSE"
	StatusFAIL    Status = "FAIL"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusUNPAY, StatusSUCCESS, StatusCLOSE, StatusFAIL:
		return nil
	default:
		return fmt.Errorf("order: invalid enum value for status field: %q", s)
	}
}
