//go:build wireinject
// +build wireinject

package server

import (
	"github.com/google/wire"

	"microservice-demo-go/pongservice/config"
	v1 "microservice-demo-go/pongservice/genproto/v1"
	"microservice-demo-go/pongservice/service"
)

// InitServer Inject service's component
func InitServer(conf *config.Config) (v1.PongServiceServer, error) {

	wire.Build(
		service.NewClient,
		service.NewServer,
	)

	return &service.Server{}, nil

}
