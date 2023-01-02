package data

import (
	"context"
	"mall/internal/conf"

	v1 "github.com/Pineana/user/api/user/v1"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewMallRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	userClient v1.UserServiceClient
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, discovery registry.Discovery) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		userClient: newUserServiceClient(discovery),
	}, cleanup, nil
}

func newUserServiceClient(discovery registry.Discovery) v1.UserServiceClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///user"),
		grpc.WithDiscovery(discovery),
		grpc.WithMiddleware(
			tracing.Client(),
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	c := v1.NewUserServiceClient(conn)
	return c
}
