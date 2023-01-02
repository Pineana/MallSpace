package service

import (
	"context"

	v1 "mall/api/mall/v1"
	"mall/internal/biz"
)

// MallService is a Mall service.
type MallService struct {
	v1.UnimplementedMallServer

	uc *biz.MallUsecase
}

// NewMallService new a Mall service.
func NewMallService(uc *biz.MallUsecase) *MallService {
	return &MallService{uc: uc}
}

// SayHello implements helloworld.MallServer.
func (s *MallService) Register(ctx context.Context, in *v1.RegisterRequest) (*v1.RegisterReply, error) {
	bo, err := s.uc.Register(ctx, &biz.RegisterRequest{
		Name:     in.Name,
		Password: in.Password,
		Sex:      in.Sex.String(),
		Year:     int(in.Year),
	})
	if err != nil {
		return nil, err
	}
	return &v1.RegisterReply{
		Id: bo.Id,
	}, nil
}

func (s *MallService) Login(ctx context.Context, in *v1.LoginRequest) (*v1.LoginReply, error) {
	bo, err := s.uc.Login(ctx, &biz.LoginRequest{
		Phone:    in.Phone,
		Password: in.Password,
	})
	if err != nil {
		return nil, err
	}
	return &v1.LoginReply{
		Token: bo.Token,
	}, nil
}
