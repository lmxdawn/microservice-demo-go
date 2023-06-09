package server

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"microservice-demo-go/pongservice/config"
	v1 "microservice-demo-go/pongservice/genproto/v1"
)

// RunGrpcServer Run grpc server
func RunGrpcServer(server v1.PongServiceServer, conf *config.Config) {

	grpcServer := grpc.NewServer()
	v1.RegisterPongServiceServer(grpcServer, server)

	fmt.Println("Listening grpc server on port" + conf.Grpc.Port)
	listen, err := net.Listen("tcp", conf.Grpc.Port)
	if err != nil {
		panic("listen grpc tcp failed.[ERROR]=>" + err.Error())
	}

	go func() {
		if err = grpcServer.Serve(listen); err != nil {
			log.Fatal("grpc serve failed", err)
		}
	}()

	conf.Grpc.Server = grpcServer

}
