package logs

import "encoding/json"

type EnvoyAccessLog struct {
	AccessTime      string `json:"access_time"`
	GatewayKind     string `json:"gateway_kind"`
	GatewayCode     string `json:"gateway_code"`
	Kind            string `json:"kind"`
	ResponseFlags   string `json:"response_flags"`
	Timestamp       int64  `json:"timestamp"`
	Remote          string `json:"remote"`
	Local           string `json:"local"`
	Method          string `json:"method"`
	Path            string `json:"path"`
	Authority       string `json:"authority"`
	Code            uint32 `json:"code"`
	ReqSize         uint64 `json:"req_size"`
	RequestTime     int64  `json:"request_time"`
	Referer         string `json:"referer"`
	Agent           string `json:"agent"`
	Forwoad         string `json:"forwoad"`
	UUID            string `json:"uuid"`
	AppId           string `json:"app_id"`
	HealthLevel     string `json:"health_level"`
	CrServiceCode   string `json:"cr_service_code"`
	SrServiceCode   string `json:"sr_service_code"`
	CrGroupCode     string `json:"cr_group_code"`
	SrGroupCode     string `json:"sr_group_code"`
	UpstreamHost    string `json:"upstream_host"`
	UpstreamCluster string `json:"upstream_cluster"`
	RouteName       string `json:"route_name"`
	// Request         string `json:"request"`
	// Response        string `json:"response"`
}

func NewEnvoyAccessLog(gatewayKind, gatewaycode, kind string) *EnvoyAccessLog {
	return &EnvoyAccessLog{
		GatewayKind: gatewayKind,
		GatewayCode: gatewaycode,
		Kind:        kind,
	}
}

func (e *EnvoyAccessLog) Marshaler() []byte {
	bs, _ := json.Marshal(e)
	return bs
}
