package server

import (
	"fmt"
	"grape/logtrans/server/fluent"
	"grape/logtrans/server/logs"
	"log"
	"net"
	"strings"
	"time"

	envoy_api_v2_core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	envoy_data_accesslog_v2 "github.com/envoyproxy/go-control-plane/envoy/data/accesslog/v2"
	pb "github.com/envoyproxy/go-control-plane/envoy/service/accesslog/v2"
	"github.com/golang/protobuf/ptypes/duration"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
)

const (
	EnvoyAccessLogPrefix = "envoy_access."
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

func splitEnvoyAccess(logName string) (gatewayKind, gatewaycode, kind string) {
	info := strings.Split(logName, ".")
	return info[1], info[2], info[3]
}

// envoy_access.log
func (s *accessLogServer) StreamAccessLogs(stream pb.AccessLogService_StreamAccessLogsServer) error {
	defer stream.SendAndClose(&pb.StreamAccessLogsResponse{})
	msg, err := stream.Recv()
	if err != nil {
		return err
	}
	if !strings.HasPrefix(msg.Identifier.LogName, EnvoyAccessLogPrefix) {
		return nil
	}
	gatewayKind, gatewaycode, kind := splitEnvoyAccess(msg.Identifier.LogName)
	for {
		if err != nil {
			return err
		}
		for _, httpLog := range msg.GetHttpLogs().LogEntry {
			logEntity := logs.NewEnvoyAccessLog(gatewayKind, gatewaycode, kind)
			parseHttpLog(httpLog, logEntity)
			fluent.AddEnvoyAccessLog(logEntity)
		}
		msg, err = stream.Recv()
	}
}

func parseHttpLog(httpLog *envoy_data_accesslog_v2.HTTPAccessLogEntry, logEntity *logs.EnvoyAccessLog) {
	// commonProperties and request and response must exist!
	commonProperties := httpLog.GetCommonProperties()
	request := httpLog.GetRequest()
	response := httpLog.GetResponse()

	logEntity.AccessTime = ppfTimestampAsTime(commonProperties.StartTime).Format(time.RFC3339)
	logEntity.ResponseFlags = responseFlagsToString(commonProperties.ResponseFlags)
	logEntity.Timestamp = ppfTimestampAsTime(commonProperties.StartTime).UnixNano() / 1000
	logEntity.Remote = addressToString(commonProperties.DownstreamRemoteAddress)
	logEntity.Local = addressToString(commonProperties.DownstreamLocalAddress)
	logEntity.Method = request.RequestMethod.String()
	logEntity.Authority = request.Authority
	logEntity.Path = request.Path
	logEntity.UUID = request.RequestId
	logEntity.Code = response.ResponseCode.Value
	logEntity.ReqSize = request.RequestBodyBytes + request.RequestHeadersBytes
	logEntity.RequestTime = ppfDurationAsDuration(commonProperties.TimeToLastUpstreamRxByte).Milliseconds()
	logEntity.Referer = request.Referer
	logEntity.Agent = request.UserAgent
	logEntity.Forwoad = request.ForwardedFor
	logEntity.UpstreamHost = addressToString(commonProperties.UpstreamLocalAddress)
	logEntity.UpstreamCluster = commonProperties.UpstreamCluster
	logEntity.RouteName = commonProperties.RouteName

	requestHeaders := request.GetRequestHeaders()
	if requestHeaders != nil {
		logEntity.CrServiceCode = requestHeaders["X-Source-Service"]
		logEntity.SrServiceCode = requestHeaders["X-Target-Service"]
		logEntity.CrGroupCode = requestHeaders["X-Source-Group"]
		logEntity.SrGroupCode = requestHeaders["X-Target-Group"]
		logEntity.AppId = requestHeaders["appid"]
	}

	responseHeaders := response.GetResponseHeaders()
	if responseHeaders != nil {
		logEntity.HealthLevel = responseHeaders["Health-Level"]
	}

}

func addressToString(addr *envoy_api_v2_core.Address) string {
	if addr != nil {
		saddr := addr.GetSocketAddress()
		if saddr != nil {
			return fmt.Sprintf("%s:%d", saddr.GetAddress(), saddr.GetPortValue())
		}
	}
	return ""
}

func ppfTimestampAsTime(t *timestamp.Timestamp) time.Time {
	return time.Unix(t.GetSeconds(), t.GetSeconds())
}

func ppfDurationAsDuration(d *duration.Duration) time.Duration {
	return time.Duration(time.Unix(d.GetSeconds(), d.GetSeconds()).UnixNano())
}

// TODO 优化if
func responseFlagsToString(f *envoy_data_accesslog_v2.ResponseFlags) string {
	flags := []string{}
	if f != nil {
		if f.FailedLocalHealthcheck {
			flags = append(flags, "failed_local_healthcheck")
		}
		if f.NoHealthyUpstream {
			flags = append(flags, "no_healthy_upstream")
		}
		if f.UpstreamRequestTimeout {
			flags = append(flags, "upstream_request_timeout")
		}
		if f.LocalReset {
			flags = append(flags, "local_reset")
		}
		if f.UpstreamRemoteReset {
			flags = append(flags, "upstream_remote_reset")
		}
		if f.UpstreamConnectionFailure {
			flags = append(flags, "upstream_connection_failure")
		}
		if f.UpstreamConnectionTermination {
			flags = append(flags, "upstream_connection_termination")
		}
		if f.UpstreamOverflow {
			flags = append(flags, "upstream_overflow")
		}
		if f.NoRouteFound {
			flags = append(flags, "no_route_found")
		}
		if f.DelayInjected {
			flags = append(flags, "delay_injected")
		}
		if f.FaultInjected {
			flags = append(flags, "fault_injected")
		}
		if f.RateLimited {
			flags = append(flags, "rate_limited")
		}
		if f.UnauthorizedDetails != nil {
			if len(f.UnauthorizedDetails.Reason.String()) > 0 {
				flags = append(flags, "unauthorized_details:"+f.UnauthorizedDetails.Reason.String())
			}
		}
		if f.RateLimitServiceError {
			flags = append(flags, "rate_limit_service_error")
		}
		if f.DownstreamConnectionTermination {
			flags = append(flags, "downstream_connection_termination")
		}
		if f.UpstreamRetryLimitExceeded {
			flags = append(flags, "upstream_retry_limit_exceeded")
		}
		if f.StreamIdleTimeout {
			flags = append(flags, "stream_idle_timeout")
		}
		if f.InvalidEnvoyRequestHeaders {
			flags = append(flags, "invalid_envoy_request_headers")
		}
		if f.DownstreamProtocolError {
			flags = append(flags, "downstream_protocol_error")
		}
	}
	return strings.Join(flags, ",")
}
