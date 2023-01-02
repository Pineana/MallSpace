package biz

import (
	"context"

	v1 "mall/api/mall/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Mall is a Mall model.
type RegisterRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Sex      string `json:"sex"`
	Year     int    `json:"year"`
}

type RegisterResponse struct {
	Id int64 `json:"id"`
}

type LoginRequest struct {
	Phone    string `json:"name"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

// MallRepo is a Greater repo.
type MallRepo interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
}

// MallUsecase is a Mall usecase.
type MallUsecase struct {
	repo MallRepo
	log  *log.Helper
}

// NewMallUsecase new a Mall usecase.
func NewMallUsecase(repo MallRepo, logger log.Logger) *MallUsecase {
	return &MallUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *MallUsecase) Register(ctx context.Context, r *RegisterRequest) (*RegisterResponse, error) {
	return uc.repo.Register(ctx, r)
}

func (uc *MallUsecase) Login(ctx context.Context, r *LoginRequest) (*LoginResponse, error) {
	return uc.repo.Login(ctx, r)
}
