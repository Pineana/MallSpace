package data

import (
	"context"

	"mall/internal/biz"

	v1 "github.com/Pineana/user/api/user/v1"
	"github.com/go-kratos/kratos/v2/log"
)

type MallRepo struct {
	data *Data
	log  *log.Helper
}

// NewMallRepo .
func NewMallRepo(data *Data, logger log.Logger) biz.MallRepo {
	return &MallRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *MallRepo) FindUserIsExist(ctx context.Context, phone string) (bool, error) {
	_, err := r.data.userClient.FindUserByPhone(ctx, &v1.FindUserByPhoneRequest{
		Phone: phone,
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *MallRepo) VerifyPassword(ctx context.Context, password string) (bool, error) {
	_, err := r.data.userClient.VerifyPassword(ctx, &v1.VerifyPasswordRequest{
		Password: password,
	})
	if err != nil {
		return nil, err
	}
	return true, nil
}

func (r *MallRepo) CreateUser(ctx context.Context, rr *biz.RegisterRequest) (*biz.RegisterResponse, error) {
	user, err := r.data.userClient.CreateUser(ctx, &v1.CreateUserRequest{
		Name:     rr.Name,
		Phone:    rr.Phone,
		Password: rr.Password,
		Sex:      v1.Sex(v1.Sex_value[rr.Sex]),
		Year:     int32(rr.Year),
	})
	if err != nil {
		return nil, err
	}
	return &biz.RegisterResponse{
		Id: user.Id,
	}, nil
}

func (r *MallRepo) Login(ctx context.Context, lr *biz.LoginRequest) (*biz.LoginResponse, error) {

	return lr, nil
}
