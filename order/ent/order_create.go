// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"order/ent/order"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OrderCreate is the builder for creating a Order entity.
type OrderCreate struct {
	config
	mutation *OrderMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (oc *OrderCreate) SetUserID(i int) *OrderCreate {
	oc.mutation.SetUserID(i)
	return oc
}

// SetOrderID sets the "order_id" field.
func (oc *OrderCreate) SetOrderID(s string) *OrderCreate {
	oc.mutation.SetOrderID(s)
	return oc
}

// SetStatus sets the "status" field.
func (oc *OrderCreate) SetStatus(o order.Status) *OrderCreate {
	oc.mutation.SetStatus(o)
	return oc
}

// SetTradeID sets the "tradeID" field.
func (oc *OrderCreate) SetTradeID(s string) *OrderCreate {
	oc.mutation.SetTradeID(s)
	return oc
}

// SetOrderMount sets the "order_mount" field.
func (oc *OrderCreate) SetOrderMount(f float32) *OrderCreate {
	oc.mutation.SetOrderMount(f)
	return oc
}

// SetPayAt sets the "pay_at" field.
func (oc *OrderCreate) SetPayAt(t time.Time) *OrderCreate {
	oc.mutation.SetPayAt(t)
	return oc
}

// SetNillablePayAt sets the "pay_at" field if the given value is not nil.
func (oc *OrderCreate) SetNillablePayAt(t *time.Time) *OrderCreate {
	if t != nil {
		oc.SetPayAt(*t)
	}
	return oc
}

// SetCreateAt sets the "create_at" field.
func (oc *OrderCreate) SetCreateAt(t time.Time) *OrderCreate {
	oc.mutation.SetCreateAt(t)
	return oc
}

// SetUpdateAt sets the "update_at" field.
func (oc *OrderCreate) SetUpdateAt(t time.Time) *OrderCreate {
	oc.mutation.SetUpdateAt(t)
	return oc
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (oc *OrderCreate) SetNillableUpdateAt(t *time.Time) *OrderCreate {
	if t != nil {
		oc.SetUpdateAt(*t)
	}
	return oc
}

// SetDeleteAt sets the "delete_at" field.
func (oc *OrderCreate) SetDeleteAt(t time.Time) *OrderCreate {
	oc.mutation.SetDeleteAt(t)
	return oc
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (oc *OrderCreate) SetNillableDeleteAt(t *time.Time) *OrderCreate {
	if t != nil {
		oc.SetDeleteAt(*t)
	}
	return oc
}

// SetIsDeleted sets the "is_deleted" field.
func (oc *OrderCreate) SetIsDeleted(b bool) *OrderCreate {
	oc.mutation.SetIsDeleted(b)
	return oc
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (oc *OrderCreate) SetNillableIsDeleted(b *bool) *OrderCreate {
	if b != nil {
		oc.SetIsDeleted(*b)
	}
	return oc
}

// Mutation returns the OrderMutation object of the builder.
func (oc *OrderCreate) Mutation() *OrderMutation {
	return oc.mutation
}

// Save creates the Order in the database.
func (oc *OrderCreate) Save(ctx context.Context) (*Order, error) {
	var (
		err  error
		node *Order
	)
	oc.defaults()
	if len(oc.hooks) == 0 {
		if err = oc.check(); err != nil {
			return nil, err
		}
		node, err = oc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = oc.check(); err != nil {
				return nil, err
			}
			oc.mutation = mutation
			if node, err = oc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(oc.hooks) - 1; i >= 0; i-- {
			if oc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = oc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, oc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Order)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from OrderMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (oc *OrderCreate) SaveX(ctx context.Context) *Order {
	v, err := oc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oc *OrderCreate) Exec(ctx context.Context) error {
	_, err := oc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oc *OrderCreate) ExecX(ctx context.Context) {
	if err := oc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oc *OrderCreate) defaults() {
	if _, ok := oc.mutation.IsDeleted(); !ok {
		v := order.DefaultIsDeleted
		oc.mutation.SetIsDeleted(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oc *OrderCreate) check() error {
	if _, ok := oc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Order.user_id"`)}
	}
	if _, ok := oc.mutation.OrderID(); !ok {
		return &ValidationError{Name: "order_id", err: errors.New(`ent: missing required field "Order.order_id"`)}
	}
	if _, ok := oc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Order.status"`)}
	}
	if v, ok := oc.mutation.Status(); ok {
		if err := order.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Order.status": %w`, err)}
		}
	}
	if _, ok := oc.mutation.TradeID(); !ok {
		return &ValidationError{Name: "tradeID", err: errors.New(`ent: missing required field "Order.tradeID"`)}
	}
	if _, ok := oc.mutation.OrderMount(); !ok {
		return &ValidationError{Name: "order_mount", err: errors.New(`ent: missing required field "Order.order_mount"`)}
	}
	if _, ok := oc.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "Order.create_at"`)}
	}
	if _, ok := oc.mutation.IsDeleted(); !ok {
		return &ValidationError{Name: "is_deleted", err: errors.New(`ent: missing required field "Order.is_deleted"`)}
	}
	return nil
}

func (oc *OrderCreate) sqlSave(ctx context.Context) (*Order, error) {
	_node, _spec := oc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (oc *OrderCreate) createSpec() (*Order, *sqlgraph.CreateSpec) {
	var (
		_node = &Order{config: oc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: order.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: order.FieldID,
			},
		}
	)
	if value, ok := oc.mutation.UserID(); ok {
		_spec.SetField(order.FieldUserID, field.TypeInt, value)
		_node.UserID = value
	}
	if value, ok := oc.mutation.OrderID(); ok {
		_spec.SetField(order.FieldOrderID, field.TypeString, value)
		_node.OrderID = value
	}
	if value, ok := oc.mutation.Status(); ok {
		_spec.SetField(order.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := oc.mutation.TradeID(); ok {
		_spec.SetField(order.FieldTradeID, field.TypeString, value)
		_node.TradeID = value
	}
	if value, ok := oc.mutation.OrderMount(); ok {
		_spec.SetField(order.FieldOrderMount, field.TypeFloat32, value)
		_node.OrderMount = value
	}
	if value, ok := oc.mutation.PayAt(); ok {
		_spec.SetField(order.FieldPayAt, field.TypeTime, value)
		_node.PayAt = value
	}
	if value, ok := oc.mutation.CreateAt(); ok {
		_spec.SetField(order.FieldCreateAt, field.TypeTime, value)
		_node.CreateAt = value
	}
	if value, ok := oc.mutation.UpdateAt(); ok {
		_spec.SetField(order.FieldUpdateAt, field.TypeTime, value)
		_node.UpdateAt = value
	}
	if value, ok := oc.mutation.DeleteAt(); ok {
		_spec.SetField(order.FieldDeleteAt, field.TypeTime, value)
		_node.DeleteAt = value
	}
	if value, ok := oc.mutation.IsDeleted(); ok {
		_spec.SetField(order.FieldIsDeleted, field.TypeBool, value)
		_node.IsDeleted = value
	}
	return _node, _spec
}

// OrderCreateBulk is the builder for creating many Order entities in bulk.
type OrderCreateBulk struct {
	config
	builders []*OrderCreate
}

// Save creates the Order entities in the database.
func (ocb *OrderCreateBulk) Save(ctx context.Context) ([]*Order, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ocb.builders))
	nodes := make([]*Order, len(ocb.builders))
	mutators := make([]Mutator, len(ocb.builders))
	for i := range ocb.builders {
		func(i int, root context.Context) {
			builder := ocb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrderMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ocb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ocb *OrderCreateBulk) SaveX(ctx context.Context) []*Order {
	v, err := ocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ocb *OrderCreateBulk) Exec(ctx context.Context) error {
	_, err := ocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocb *OrderCreateBulk) ExecX(ctx context.Context) {
	if err := ocb.Exec(ctx); err != nil {
		panic(err)
	}
}
