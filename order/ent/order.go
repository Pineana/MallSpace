// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"order/ent/order"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Order is the model entity for the Order schema.
type Order struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// OrderID holds the value of the "order_id" field.
	OrderID string `json:"order_id,omitempty"`
	// Status holds the value of the "status" field.
	Status order.Status `json:"status,omitempty"`
	// TradeID holds the value of the "tradeID" field.
	TradeID string `json:"tradeID,omitempty"`
	// OrderMount holds the value of the "order_mount" field.
	OrderMount float32 `json:"order_mount,omitempty"`
	// PayAt holds the value of the "pay_at" field.
	PayAt time.Time `json:"pay_at,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt time.Time `json:"create_at,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	UpdateAt time.Time `json:"update_at,omitempty"`
	// DeleteAt holds the value of the "delete_at" field.
	DeleteAt time.Time `json:"delete_at,omitempty"`
	// IsDeleted holds the value of the "is_deleted" field.
	IsDeleted bool `json:"is_deleted,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Order) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case order.FieldIsDeleted:
			values[i] = new(sql.NullBool)
		case order.FieldOrderMount:
			values[i] = new(sql.NullFloat64)
		case order.FieldID, order.FieldUserID:
			values[i] = new(sql.NullInt64)
		case order.FieldOrderID, order.FieldStatus, order.FieldTradeID:
			values[i] = new(sql.NullString)
		case order.FieldPayAt, order.FieldCreateAt, order.FieldUpdateAt, order.FieldDeleteAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Order", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Order fields.
func (o *Order) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case order.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			o.ID = int(value.Int64)
		case order.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				o.UserID = int(value.Int64)
			}
		case order.FieldOrderID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field order_id", values[i])
			} else if value.Valid {
				o.OrderID = value.String
			}
		case order.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				o.Status = order.Status(value.String)
			}
		case order.FieldTradeID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tradeID", values[i])
			} else if value.Valid {
				o.TradeID = value.String
			}
		case order.FieldOrderMount:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field order_mount", values[i])
			} else if value.Valid {
				o.OrderMount = float32(value.Float64)
			}
		case order.FieldPayAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field pay_at", values[i])
			} else if value.Valid {
				o.PayAt = value.Time
			}
		case order.FieldCreateAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				o.CreateAt = value.Time
			}
		case order.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				o.UpdateAt = value.Time
			}
		case order.FieldDeleteAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field delete_at", values[i])
			} else if value.Valid {
				o.DeleteAt = value.Time
			}
		case order.FieldIsDeleted:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_deleted", values[i])
			} else if value.Valid {
				o.IsDeleted = value.Bool
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Order.
// Note that you need to call Order.Unwrap() before calling this method if this Order
// was returned from a transaction, and the transaction was committed or rolled back.
func (o *Order) Update() *OrderUpdateOne {
	return (&OrderClient{config: o.config}).UpdateOne(o)
}

// Unwrap unwraps the Order entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (o *Order) Unwrap() *Order {
	_tx, ok := o.config.driver.(*txDriver)
	if !ok {
		panic("ent: Order is not a transactional entity")
	}
	o.config.driver = _tx.drv
	return o
}

// String implements the fmt.Stringer.
func (o *Order) String() string {
	var builder strings.Builder
	builder.WriteString("Order(")
	builder.WriteString(fmt.Sprintf("id=%v, ", o.ID))
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", o.UserID))
	builder.WriteString(", ")
	builder.WriteString("order_id=")
	builder.WriteString(o.OrderID)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", o.Status))
	builder.WriteString(", ")
	builder.WriteString("tradeID=")
	builder.WriteString(o.TradeID)
	builder.WriteString(", ")
	builder.WriteString("order_mount=")
	builder.WriteString(fmt.Sprintf("%v", o.OrderMount))
	builder.WriteString(", ")
	builder.WriteString("pay_at=")
	builder.WriteString(o.PayAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("create_at=")
	builder.WriteString(o.CreateAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_at=")
	builder.WriteString(o.UpdateAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("delete_at=")
	builder.WriteString(o.DeleteAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("is_deleted=")
	builder.WriteString(fmt.Sprintf("%v", o.IsDeleted))
	builder.WriteByte(')')
	return builder.String()
}

// Orders is a parsable slice of Order.
type Orders []*Order

func (o Orders) config(cfg config) {
	for _i := range o {
		o[_i].config = cfg
	}
}