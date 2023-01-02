package data

import (
	"context"
	"fmt"
	"user/ent"
	"user/internal/conf"

	v1 "github.com/Pineana/order/api/order/v1"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	entClient   *ent.Client
	orderClient v1.OrderServiceClient
}

// NewData .
func NewData(confData *conf.Data, logger log.Logger, discovery registry.Discovery) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		entClient:   newEntClient(confData),
		orderClient: newOrderServiceClient(discovery),
	}, cleanup, nil
}

func newEntClient(c *conf.Data) *ent.Client {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", c.Database.Username, c.Database.Password, c.Database.Host, c.Database.Port, c.Database.DatabaseName)
	fmt.Println(url)
	client, err := ent.Open(c.Database.Driver, url)
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}

func newOrderServiceClient(discovery registry.Discovery) v1.OrderServiceClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///order"),
		grpc.WithDiscovery(discovery),
		grpc.WithMiddleware(
			tracing.Client(),
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	c := v1.NewOrderServiceClient(conn)
	return c
}
