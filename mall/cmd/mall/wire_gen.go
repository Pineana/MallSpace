// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"mall/internal/biz"
	"mall/internal/conf"
	"mall/internal/data"
	"mall/internal/server"
	"mall/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	discovery := server.NewDiscovery(confServer)
	dataData, cleanup, err := data.NewData(confData, logger, discovery)
	if err != nil {
		return nil, nil, err
	}
	mallRepo := data.NewMallRepo(dataData, logger)
	mallUsecase := biz.NewMallUsecase(mallRepo, logger)
	mallService := service.NewMallService(mallUsecase)
	grpcServer := server.NewGRPCServer(confServer, mallService, logger)
	httpServer := server.NewHTTPServer(confServer, mallService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
