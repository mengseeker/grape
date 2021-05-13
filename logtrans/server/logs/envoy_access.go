package logs

import (
	"encoding/json"
	"fmt"
	"strings"

	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	v2 "github.com/envoyproxy/go-control-plane/envoy/data/accesslog/v2"
)

type EnvoyAccessLog struct {
	UUID      string `json:"uuid"`
	Timestamp int64  `json:"timestamp"`
	LogName   string `json:"kind"`

	Method string `json:"method"`
	Path   string `json:"path"`

	RequestSize uint64 `json:"req_size"`
	RequestTime int64  `json:"request_time"`
	Agent       string `json:"agent"`

	ResponseFlags string `json:"response_flags"`
	ResponseCode  uint32 `json:"code"`

	Remote string `json:"remote"`
	Local  string `json:"local"`
	AppId  string `json:"app_id"`
	// HealthLevel string `json:"health_level"`

	SourceService   string `json:"cr_service_code"`
	TargetService   string `json:"sr_service_code"`
	SourceGroup     string `json:"cr_group_code"`
	TargetGroup     string `json:"sr_group_code"`
	Upstream        string `json:"upstream_host"`
	UpstreamCluster string `json:"upstream_cluster"`
	RouteName       string `json:"route_name"`
}

func (e *EnvoyAccessLog) Marshaler() []byte {
	bs, _ := json.Marshal(e)
	return bs
}

func ParseEnvoyAccessLog(httpLog *v2.HTTPAccessLogEntry, LogName string) *EnvoyAccessLog {
	commonProperties := httpLog.GetCommonProperties()
	request := httpLog.GetRequest()
	response := httpLog.GetResponse()
	logEntity := EnvoyAccessLog{
		LogName:   LogName,
		UUID:      request.RequestId,
		Path:      request.Path,
		Method:    request.RequestMethod.String(),
		Timestamp: commonProperties.StartTime.AsTime().UnixNano(),

		RequestSize: request.RequestBodyBytes + request.RequestHeadersBytes,
		RequestTime: commonProperties.TimeToLastUpstreamRxByte.AsDuration().Milliseconds(),
		Agent:       request.UserAgent,

		ResponseCode:  response.ResponseCode.Value,
		ResponseFlags: responseFlagsToString(commonProperties.ResponseFlags),
		Remote:        addressToString(commonProperties.DownstreamRemoteAddress),
		Local:         addressToString(commonProperties.DownstreamLocalAddress),

		Upstream:        addressToString(commonProperties.UpstreamLocalAddress),
		UpstreamCluster: commonProperties.UpstreamCluster,
		RouteName:       commonProperties.RouteName,
	}

	requestHeaders := request.GetRequestHeaders()
	if requestHeaders != nil {
		logEntity.SourceService = ""
		logEntity.TargetService = ""
		logEntity.SourceGroup = ""
		logEntity.TargetGroup = ""
		logEntity.AppId = requestHeaders["appid"]
	}

	return &logEntity
}

func addressToString(addr *core.Address) string {
	if addr != nil {
		saddr := addr.GetSocketAddress()
		if saddr != nil {
			return fmt.Sprintf("%s:%d", saddr.GetAddress(), saddr.GetPortValue())
		}
	}
	return ""
}

// func ppfTimestampAsTime(t *timestamp.Timestamp) time.Time {
// 	return time.Unix(t.GetSeconds(), t.GetSeconds())
// }

// func ppfDurationAsDuration(d *duration.Duration) time.Duration {
// 	return time.Duration(time.Unix(d.GetSeconds(), d.GetSeconds()).UnixNano())
// }

// TODO 优化if
func responseFlagsToString(f *v2.ResponseFlags) string {
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
