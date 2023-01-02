package data

import (
	"context"
	"time"

	opkg "order/ent/order"
	"order/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
)

type OrderRepo struct {
	data *Data
	log  *log.Helper
}

// NewOrderRepo .
func NewOrderRepo(data *Data, logger log.Logger) biz.OrderRepo {
	return &OrderRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *OrderRepo) Save(ctx context.Context, g *biz.Order) (*biz.Order, error) {
	order, err := r.data.DBClient.Order.
		Create().
		SetUserID(g.UserID).
		SetOrderID(uuid.NewString()).
		SetStatus(opkg.Status(g.Status)).
		SetTradeID(g.TradeID).
		SetOrderMount(g.OrderMount).
		SetCreateAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Order{
		ID:         order.ID,
		UserID:     order.UserID,
		OrderID:    order.OrderID,
		Status:     string(order.Status),
		TradeID:    order.TradeID,
		OrderMount: order.OrderMount,
		CreateAt:   order.CreateAt,
	}, nil
}

func (r *OrderRepo) Update(ctx context.Context, g *biz.Order) (*biz.Order, error) {
	order, err := r.data.DBClient.Order.Query().Where(opkg.OrderID(g.OrderID)).Only(ctx)
	if err != nil {
		return nil, err
	}

	order, err = order.Update().SetStatus(opkg.Status(g.Status)).Save(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Order{
		ID:         order.ID,
		UserID:     order.UserID,
		OrderID:    order.OrderID,
		Status:     string(order.Status),
		TradeID:    order.TradeID,
		OrderMount: order.OrderMount,
		CreateAt:   order.CreateAt,
	}, nil
}

func (r *OrderRepo) FindByOrderID(ctx context.Context, orderID string) (*biz.Order, error) {
	order, err := r.data.DBClient.Order.
		Query().
		Where(opkg.OrderID(orderID)).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Order{
		ID:         order.ID,
		UserID:     order.UserID,
		OrderID:    order.OrderID,
		Status:     string(order.Status),
		TradeID:    order.TradeID,
		OrderMount: order.OrderMount,
		CreateAt:   order.CreateAt,
	}, nil
}

func (r *OrderRepo) ListAll(ctx context.Context) ([]*biz.Order, error) {
	orders, err := r.data.DBClient.Order.
		Query().All(ctx)
	if err != nil {
		return nil, err
	}
	borders := make([]*biz.Order, 0)
	for _, v := range orders {
		borders = append(borders, &biz.Order{
			ID:         v.ID,
			UserID:     v.UserID,
			OrderID:    v.OrderID,
			Status:     string(v.Status),
			TradeID:    v.TradeID,
			OrderMount: v.OrderMount,
			CreateAt:   v.CreateAt,
		})
	}
	return borders, nil
}
