package server

import (
	"grape/internal/share"
	"grape/logtrans/server/forward"
	"grape/logtrans/server/logs"
	"log"
	"net"
	"strings"

	pb "github.com/envoyproxy/go-control-plane/envoy/service/accesslog/v2"
	"google.golang.org/grpc"
)

type accessLogServer struct {
	// pb.UnimplementedAccessLogServiceServer
}

func ServeAccessLog(address string) {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAccessLogServiceServer(s, &accessLogServer{})
	log.Println("ServeAccessLog at: ", address)
	log.Fatal(s.Serve(lis))
}

func (s *accessLogServer) StreamAccessLogs(stream pb.AccessLogService_StreamAccessLogsServer) error {
	defer stream.SendAndClose(&pb.StreamAccessLogsResponse{})
	for {
		msg, err := stream.Recv()
		if err != nil {
			return err
		}
		// envoy_access.log
		if strings.HasPrefix(msg.Identifier.LogName, share.EnvoyAccessLogPrefix) {
			for _, httpLog := range msg.GetHttpLogs().LogEntry {
				logEntity := logs.ParseEnvoyAccessLog(httpLog, msg.Identifier.LogName)
				forward.AddEnvoyAccessLog(logEntity)
			}
		}
	}
}
