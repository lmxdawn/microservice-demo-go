//go:build wireinject
// +build wireinject

package server

import (
	"github.com/google/wire"

	"microservice-demo-go/pingservice/config"
	v1 "microservice-demo-go/pingservice/genproto/v1"
	"microservice-demo-go/pingservice/service"
)

// InitServer Inject service's component
func InitServer(conf *config.Config) (v1.PingServiceServer, error) {

	wire.Build(
		service.NewPongClient,
		service.NewClient,
		service.NewServer,
	)

	return &service.Server{}, nil

}
