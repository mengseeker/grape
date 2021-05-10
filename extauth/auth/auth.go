package auth

import (
	"time"

	v31 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	pb "github.com/envoyproxy/go-control-plane/envoy/service/auth/v3"
	"google.golang.org/genproto/googleapis/rpc/status"
)

func NeedAuth(path string) bool {
	return true
}

func Auth(token, path string) (*pb.CheckResponse, error) {
	log.Debugf("auth path: %s, token: %s", path, token)
	resp := pb.CheckResponse{
		Status: &status.Status{Code: 200},
		HttpResponse: &pb.CheckResponse_OkResponse{
			OkResponse: &pb.OkHttpResponse{
				Headers: []*v31.HeaderValueOption{
					MakeHeader("auth-at", time.Now().Format(time.RFC3339)),
				},
			},
		},
	}
	return &resp, nil
}

func MakeHeader(key, val string) *v31.HeaderValueOption {
	headerVal := &v31.HeaderValue{
		Key:   key,
		Value: val,
	}
	return &v31.HeaderValueOption{Header: headerVal}
}
