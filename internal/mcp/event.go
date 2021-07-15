package mcp

import (
	"grape/api"

	"google.golang.org/grpc"
)

type EventServer struct {
	api.UnimplementedEventServiceServer
}

func (s *EventServer) RegisterServer(grpcServer *grpc.Server) {
	api.RegisterEventServiceServer(grpcServer, s)
}
