package logs

const (
// IndexEnvoyAccessMapping = ``
)

type EnvoyAccess struct {
	UUID        string `json:"uuid"`
	Timestamp   int64  `json:"timestamp"`
	Timestamp1  int64  `json:"@timestamp"`
	Kind        string `json:"kind"`
	GatewayKind string `json:"gateway_kind"`
	GatewayCode string `json:"gateway_code"`

	Method        string `json:"method"`
	Path          string `json:"path"`
	Code          int    `json:"code"`
	Status        string `json:"status"`
	ReqSize       uint64 `json:"req_size"`
	ResSize       int64  `json:"res_size"`
	Traffic       int    `json:"traffic"`
	RequestTime   int64  `json:"request_time"`
	Agent         string `json:"agent"`
	ResponseFlags string `json:"response_flags"`

	Remote      string `json:"remote"`
	Local       string `json:"local"`
	AppId       string `json:"app_id"`
	HealthLevel string `json:"health_level"`

	SourceService   string `json:"cr_service_code"`
	TargetService   string `json:"sr_service_code"`
	SourceGroup     string `json:"cr_group_code"`
	TargetGroup     string `json:"sr_group_code"`
	Upstream        string `json:"upstream_host"`
	UpstreamCluster string `json:"upstream_cluster"`
	RouteName       string `json:"route_name"`
	AggHost         string `json:"agg_host"`

	Tenant          string `json:"tenant"`
	EnvironmentCode string `json:"environment_code"`
	// 兼容方案，namespace_code也是environment_code
	NamespaceCode string `json:"namespace_code"`
	ClusterCode   string `json:"cluster_code"`
}
