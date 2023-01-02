package biz

import (
	"context"

	v1 "user/api/user/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// User is a User model.
type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Sex      string `json:"sex"`
	Age      int    `json:"age"`
}

// UserRepo is a Greater repo.
type UserRepo interface {
	Save(context.Context, *User) (*User, error)
	Update(context.Context, *User) (*User, error)
	FindByPhone(context.Context, string) (*User, error)
	ListAll(context.Context) ([]*User, error)
	Delete(context.Context, int64) error
}

// UserUsecase is a User usecase.
type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

// NewUserUsecase new a User usecase.
func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateUser creates a User, and returns the new User.
func (uc *UserUsecase) CreateUser(ctx context.Context, g *User) (*User, error) {
	return uc.repo.Save(ctx, g)
}

func (uc *UserUsecase) UpdateUser(ctx context.Context, g *User) (*User, error) {
	return uc.repo.Update(ctx, g)
}

func (uc *UserUsecase) FindUserByPhone(ctx context.Context, phone string) (*User, error) {
	return uc.repo.FindByPhone(ctx, phone)
}

func (uc *UserUsecase) ListAll(ctx context.Context) ([]*User, error) {
	return uc.repo.ListAll(ctx)
}

func (uc *UserUsecase) DeleteUserByID(ctx context.Context, id int64) error {
	return uc.repo.Delete(ctx, id)
}
