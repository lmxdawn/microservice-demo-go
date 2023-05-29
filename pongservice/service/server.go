package service

import (
	"context"

	"microservice-demo-go/pongservice/config"
	v1 "microservice-demo-go/pongservice/genproto/v1"
)

// Server Server struct
type Server struct {
	v1.UnimplementedPongServiceServer
	pongClient v1.PongServiceClient
	conf       *config.Config
}

// NewServer New service grpc server
func NewServer(conf *config.Config, pongClient v1.PongServiceClient) v1.PongServiceServer {
	return &Server{
		pongClient: pongClient,
		conf:       conf,
	}
}

func (s *Server) Pong(ctx context.Context, req *v1.PongRequest) (*v1.PongResponse, error) {
	return &v1.PongResponse{Msg: "response pong msg:" + req.Msg}, nil
}
