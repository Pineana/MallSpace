package data

import (
	"fmt"
	"order/ent"
	"order/internal/conf"

	"context"

	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewOrderRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	DBClient *ent.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		DBClient: newEntClient(c),
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
