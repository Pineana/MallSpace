package service

import (
	"context"

	v1 "order/api/order/v1"
	"order/internal/biz"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// OrderService is a Order service.
type OrderService struct {
	v1.UnimplementedOrderServiceServer

	uc *biz.OrderUsecase
}

// NewOrderService new a Order service.
func NewOrderService(uc *biz.OrderUsecase) *OrderService {
	return &OrderService{uc: uc}
}

// SayHello implements helloworld.OrderServer.
func (s *OrderService) UpdateOrderStatus(ctx context.Context, req *v1.UpdateOrderStatusRequest) (*emptypb.Empty, error) {
	_, err := s.uc.UpdateOrderStatus(ctx, req.OrderId)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, err
}

func (s *OrderService) CreateOrder(ctx context.Context, req *v1.CreateOrderRequest) (*v1.Order, error) {
	bi := &biz.Order{
		UserID:     int(req.Order.UserId),
		OrderMount: req.Order.OrderMount,
	}
	order, err := s.uc.CreateOrder(ctx, bi)
	if err != nil {
		return nil, err
	}
	return &v1.Order{
		OrderId:    order.OrderID,
		UserId:     int64(order.UserID),
		Status:     v1.Status(v1.Status_value[order.Status]),
		OrderMount: order.OrderMount,
		CreateAt:   timestamppb.New(order.CreateAt),
	}, nil
}

func (s *OrderService) DeleteOrder(_ context.Context, _ *v1.DeleteOrderRequest) (*emptypb.Empty, error) {
	panic("not implemented") // TODO: Implement
}

func (s *OrderService) GetOrder(ctx context.Context, req *v1.GetOrderRequest) (*v1.Order, error) {
	order, err := s.uc.GetOrderByOrderID(ctx, req.OrderId)
	if err != nil {
		return nil, err
	}
	return &v1.Order{
		OrderId:    order.OrderID,
		UserId:     int64(order.UserID),
		Status:     v1.Status(v1.Status_value[order.Status]),
		OrderMount: order.OrderMount,
		CreateAt:   timestamppb.New(order.CreateAt),
	}, nil
}
