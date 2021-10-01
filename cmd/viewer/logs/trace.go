package logs

const (
// IndexTraceMapping = ``
)

type Trace struct {
	TraceId       string                 `json:"traceId"`
	Id            string                 `json:"id"`
	ParentId      string                 `json:"parent_id"`
	Kind          string                 `json:"kind"`
	Duration      int64                  `json:"duration"`
	Timestamp     int64                  `json:"timestamp"`
	LocalEndpoint map[string]interface{} `json:"localEndpoint"`
	Tags          map[string]interface{} `json:"tags"`
	Name          string                 `json:"name"`

	Tenant          string `json:"tenant"`
	EnvironmentCode string `json:"environment_code"`
	// 兼容方案，namespace_code也是environment_code
	NamespaceCode string `json:"namespace_code"`
	ClusterCode   string `json:"cluster_code"`
}
