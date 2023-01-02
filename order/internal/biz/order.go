package biz

import (
	"context"
	"fmt"
	v1 "order/api/order/v1"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Order is a Order model.
type Order struct {
	ID int `json:"id,omitempty"`

	UserID int `json:"user_id,omitempty"`

	OrderID string `json:"order_id,omitempty"`

	Status string `json:"status,omitempty"`

	TradeID string `json:"tradeID,omitempty"`

	OrderMount float32 `json:"order_mount,omitempty"`

	PayAt time.Time `json:"pay_at,omitempty"`

	CreateAt time.Time `json:"created_at,omitempty"`

	UpdateAt time.Time `json:"updated_at,omitempty"`

	DeleteAt time.Time `json:"deleted_at,omitempty"`

	IsDeleted bool `json:"is_deleted,omitempty"`
}

// OrderRepo is a Greater repo.
type OrderRepo interface {
	Save(context.Context, *Order) (*Order, error)
	Update(context.Context, *Order) (*Order, error)
	FindByOrderID(context.Context, string) (*Order, error)
	ListAll(context.Context) ([]*Order, error)
}

// OrderUsecase is a Order usecase.
type OrderUsecase struct {
	repo OrderRepo
	log  *log.Helper
}

// NewOrderUsecase new a Order usecase.
func NewOrderUsecase(repo OrderRepo, logger log.Logger) *OrderUsecase {
	return &OrderUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateOrder creates a Order, and returns the new Order.
func (uc *OrderUsecase) CreateOrder(ctx context.Context, o *Order) (*Order, error) {
	o.CreateAt = time.Now()
	o.OrderID = uuid.NewString()
	o.Status = v1.Status_UNPAY.String()
	fmt.Println(o)
	return uc.repo.Save(ctx, o)
}

func (uc *OrderUsecase) UpdateOrderStatus(ctx context.Context, orderID string) (*Order, error) {
	return uc.repo.Update(ctx, &Order{OrderID: orderID})
}

func (uc *OrderUsecase) ListAll(ctx context.Context, o *Order) ([]*Order, error) {
	return uc.repo.ListAll(ctx)
}

func (uc *OrderUsecase) GetOrderByOrderID(ctx context.Context, orderID string) (*Order, error) {
	return uc.repo.FindByOrderID(ctx, orderID)
}
