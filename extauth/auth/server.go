package auth

import (
	"context"
	"net"
	"strings"

	pb "github.com/envoyproxy/go-control-plane/envoy/service/auth/v3"
	"google.golang.org/grpc"
)

type authServer struct {
	// pb.UnimplementedAuthorizationServer
}

func (s *authServer) Check(c context.Context, req *pb.CheckRequest) (*pb.CheckResponse, error) {
	// token := req.Attributes.Request.Http.GetHeaders()["Token"]
	// log.Debug(req.Attributes.Request.Http.GetHeaders())
	token := req.Attributes.Request.Http.GetHeaders()["token"]
	// reqID := req.Attributes.Request.Http.Id
	reqID := req.Attributes.Request.Http.GetHeaders()["x-request-id"]
	path := CutHttpPath(req.Attributes.Request.Http.Path)
	method := req.Attributes.Request.Http.Method
	endpoint := GetEndpoint(method, path)
	if !NeedAuth(endpoint) {
		log.Debugf("%s auth ignored", endpoint)
		return &OkResp, nil
	}
	return Auth(endpoint, reqID, token)
}

func Serve(address, clusterCode string) {
	initConfig(clusterCode)
	go watchApp(clusterCode)
	go watchToken(clusterCode)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAuthorizationServer(s, &authServer{})
	log.Info("auth server run at: ", address)
	log.Fatal(s.Serve(lis))
}

func CutHttpPath(rawPath string) string {
	idx := strings.Index(rawPath, "?")
	if idx >= 0 {
		return rawPath[0:idx]
	}
	return rawPath
}
