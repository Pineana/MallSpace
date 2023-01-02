package service

import (
	"context"

	v1 "user/api/user/v1"
	"user/internal/biz"

	"google.golang.org/protobuf/types/known/emptypb"
)

// UserService is a User service.
type UserService struct {
	v1.UnimplementedUserServiceServer

	uc *biz.UserUsecase
}

// NewUserService new a User service.
func NewUserService(uc *biz.UserUsecase) *UserService {
	return &UserService{uc: uc}
}

// SayHello implements helloworld.UserServer.
func (s *UserService) CreateUser(ctx context.Context, req *v1.CreateUserRequest) (*v1.User, error) {
	bi := &biz.User{
		Name:     req.Name,
		Password: req.Password,
		Phone:    req.Phone,
		Sex:      req.Sex.String(),
		Age:      int(req.Year),
	}
	bo, err := s.uc.CreateUser(ctx, bi)
	if err != nil {
		return nil, err
	}
	return &v1.User{
		Id:       bo.ID,
		Name:     bo.Name,
		Password: bo.Password,
		Phone:    bo.Phone,
		Sex:      v1.Sex(v1.Sex_value[bo.Sex]),
		Age:      int32(bo.Age),
	}, nil
}
func (s *UserService) DeleteUser(ctx context.Context, req *v1.DeleteUserRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.uc.DeleteUserByID(ctx, req.Id)
}

func (s *UserService) FindUserByPhone(ctx context.Context, req *v1.FindUserByPhoneRequest) (*v1.User, error) {
	bo, err := s.uc.FindUserByPhone(ctx, req.Phone)
	if err != nil {
		return nil, err
	}
	return &v1.User{
		Id:       bo.ID,
		Name:     bo.Name,
		Password: bo.Password,
		Phone:    bo.Phone,
		Sex:      v1.Sex(v1.Sex_value[bo.Sex]),
		Age:      int32(bo.Age),
	}, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (*v1.User, error) {
	bi := &biz.User{
		ID:       req.Id,
		Name:     req.User.Name,
		Password: req.User.Password,
		Phone:    req.User.Phone,
		Sex:      req.User.Sex.String(),
		Age:      int(req.User.Age),
	}
	bo, err := s.uc.UpdateUser(ctx, bi)
	if err != nil {
		return nil, err
	}
	return &v1.User{
		Id:       bo.ID,
		Name:     bo.Name,
		Password: bo.Password,
		Phone:    bo.Phone,
		Sex:      v1.Sex(v1.Sex_value[bo.Sex]),
		Age:      int32(bo.Age),
	}, nil
}
