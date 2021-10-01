package apiv3

import (
	"context"
	"net"
	"time"

	"google.golang.org/grpc"

	discoverygrpc "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	cachev3 "github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	serverv3 "github.com/envoyproxy/go-control-plane/pkg/server/v3"
)

const (
	grpcMaxConcurrentStreams = 1000000
	heartbeatInterval        = 60
)

var (
	v3server *V3Server
)

type V3Server struct {
	Address string
	Cache   cachev3.Cache
	Cb      *Callbacks
	Hash    *nodeHash
	ctx     context.Context
	srv3    serverv3.Server
}

func Serve(address string) {
	ctx := context.Background()
	hash := &nodeHash{}
	cache := cachev3.NewSnapshotCacheWithHeartbeating(ctx, true, hash, l, time.Second*heartbeatInterval)
	v3server = &V3Server{
		Address: address,
		Cache:   cache,
		Cb:      &Callbacks{},
		Hash:    hash,
	}
	v3server.serve()
}

func (s *V3Server) serve() {
	srv3 := serverv3.NewServer(s.ctx, s.Cache, s.Cb)
	s.srv3 = srv3
	var grpcOptions []grpc.ServerOption
	grpcOptions = append(grpcOptions, grpc.MaxConcurrentStreams(grpcMaxConcurrentStreams))
	grpcServer := grpc.NewServer(grpcOptions...)
	lis, err := net.Listen("tcp", s.Address)
	if err != nil {
		l.Fatal(err)
	}

	registerServer(grpcServer, srv3)

	l.Infof("server listening on %v", s.Address)
	if err = grpcServer.Serve(lis); err != nil {
		l.Fatal(err)
	}
}

func registerServer(grpcServer *grpc.Server, server serverv3.Server) {
	discoverygrpc.RegisterAggregatedDiscoveryServiceServer(grpcServer, server)
}
