package server

import (
	"grape/internal/mcp"
	"net"
	"time"

	"google.golang.org/grpc"
)

type mcpServer struct {
	discovery *mcp.DiscoveryServer
}

func newMCPServer() *mcpServer {
	dis, err := mcp.NewDiscoveryServer(log)
	if err != nil {
		log.Fatalf("failed to create mcpserver: %q", err.Error())
	}
	return &mcpServer{
		discovery: dis,
	}
}

func (s *mcpServer) RunListener(lis net.Listener) error {
	grpcserver := grpc.NewServer()
	s.discovery.RegisterServer(grpcserver)
	go s.discovery.HandleUpdate()
	defer s.discovery.Cancel()
	// test
	go s.RunTests()
	return grpcserver.Serve(lis)
}

func (s *mcpServer) RunTests() {
	tk := time.NewTicker(time.Second)
	for range tk.C {
		s.discovery.Update()
	}
}
