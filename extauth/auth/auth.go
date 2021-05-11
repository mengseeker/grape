package auth

import (
	v31 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	pb "github.com/envoyproxy/go-control-plane/envoy/service/auth/v3"
	t3 "github.com/envoyproxy/go-control-plane/envoy/type/v3"
	"google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
)

var (
	OkResp = pb.CheckResponse{
		Status: &status.Status{Code: int32(codes.OK)},
	}
)

func NeedAuth(method, path string) bool {
	return GetEndpoint(method, path) > 0
	// return rand.Int()%2 == 0
}

func Auth(method, path, token string) (*pb.CheckResponse, error) {
	endpoint := GetEndpoint(method, path)
	app, err := GetAppByToken(token)
	if err != nil {
		log.Warnf("auth fail: %v(path: %s, token: %s)", err, path, token)
		return FailResponse(err.Error()), nil
	}
	if err = app.Auth(endpoint); err != nil {
		log.Warnf("auth fail: %v(path: %s, token: %s)", err, path, token)
		return FailResponse(err.Error()), nil
	}
	log.Debugf("auth ok(path: %s, token: %s)", path, token)
	authHeaders := app.Headers()
	resp := pb.CheckResponse{
		Status: &status.Status{Code: int32(codes.OK)},
		HttpResponse: &pb.CheckResponse_OkResponse{
			OkResponse: &pb.OkHttpResponse{
				Headers: MakeHeaders(authHeaders),
			},
		},
	}
	return &resp, nil
}

// TODO add fail message
func FailResponse(message string) *pb.CheckResponse {
	return &pb.CheckResponse{
		Status: &status.Status{Code: int32(codes.PermissionDenied)},
		HttpResponse: &pb.CheckResponse_DeniedResponse{
			DeniedResponse: &pb.DeniedHttpResponse{
				Status: &t3.HttpStatus{
					Code: t3.StatusCode_Forbidden,
				},
				Headers: []*v31.HeaderValueOption{
					MakeHeader("X-Authentication", message),
				},
			},
		},
	}
}

func MakeHeaders(hs map[string]string) []*v31.HeaderValueOption {
	hv := []*v31.HeaderValueOption{}
	for k, v := range hs {
		hv = append(hv, MakeHeader(k, v))
	}
	return hv
}

func MakeHeader(key, val string) *v31.HeaderValueOption {
	headerVal := &v31.HeaderValue{
		Key:   key,
		Value: val,
	}
	return &v31.HeaderValueOption{Header: headerVal}
}
