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
	token := req.Attributes.Request.Http.GetHeaders()["Token"]
	path := CutHttpPath(req.Attributes.Request.Http.Path)
	method := req.Attributes.Request.Http.Method
	if !NeedAuth(method, path) {
		return &OkResp, nil
	}
	return Auth(method, path, token)
}

func Serve(address string) {
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
