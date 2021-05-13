package apiv3

import (
	"context"
	"grape/pilot/apiv3/cache"
	"net"

	"google.golang.org/grpc"

	discoverygrpc "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	serverv3 "github.com/envoyproxy/go-control-plane/pkg/server/v3"
)

const (
	grpcMaxConcurrentStreams = 1000000
)

// Start start xds server
func Serve(address string) {
	// Run the xDS server
	ctx := context.Background()
	cb := &Callbacks{}
	srv3 := serverv3.NewServer(ctx, cache.New(l), cb)
	var grpcOptions []grpc.ServerOption
	grpcOptions = append(grpcOptions, grpc.MaxConcurrentStreams(grpcMaxConcurrentStreams))
	grpcServer := grpc.NewServer(grpcOptions...)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		l.Fatal(err)
	}

	registerServer(grpcServer, srv3)

	l.Infof("server listening on %v", address)
	if err = grpcServer.Serve(lis); err != nil {
		l.Fatal(err)
	}
}

func registerServer(grpcServer *grpc.Server, server serverv3.Server) {
	discoverygrpc.RegisterAggregatedDiscoveryServiceServer(grpcServer, server)
}
