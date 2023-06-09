// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package server

import (
	"microservice-demo-go/pingservice/config"
	"microservice-demo-go/pingservice/genproto/v1"
	"microservice-demo-go/pingservice/service"
)

// Injectors from wire.go:

// InitServer Inject service's component
func InitServer(conf *config.Config) (v1.PingServiceServer, error) {
	pingServiceClient, err := service.NewClient(conf)
	if err != nil {
		return nil, err
	}
	pongServiceClient, err := service.NewPongClient(conf)
	if err != nil {
		return nil, err
	}
	pingServiceServer := service.NewServer(conf, pingServiceClient, pongServiceClient)
	return pingServiceServer, nil
}
