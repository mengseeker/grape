package worker

import (
	"grape/pkg/logger"
	"strings"

	"github.com/olivere/elastic/v7"
)

const (
	EnvoyAccessTag = "envoy_access"
	TraceTag       = "trace"
	esBatchMaxSize = 100
)

type EsClient struct {
	// Addr  string
	esCli           *elastic.Client
	ClusterCode     string
	EnvironmentCode string

	l logger.Logger
}

func NewEsClient(addr, env, cluster string, l logger.Logger) (*EsClient, error) {
	urls := strings.Split(addr, ",")
	cli, err := elastic.NewClient(elastic.SetURL(urls...))
	if err != nil {
		return nil, err
	}
	_, err = cli.ElasticsearchVersion(urls[0])
	if err != nil {
		return nil, err
	}
	es := new(EsClient)
	es.esCli = cli
	es.ClusterCode = cluster
	es.EnvironmentCode = env
	es.l = l
	return es, nil
}

func (e *EsClient) Write([]*Message) {

}

// func (e *EsClient) dealEnvoyAccessLog(m *Message) error {
// 	data := new(logs.EnvoyAccess)
// 	err := json.Unmarshal(m.Value, data)
// 	if err != nil {
// 		return err
// 	}
// 	data.Tenant = e.EnvironmentCode
// 	data.EnvironmentCode = e.EnvironmentCode
// 	data.NamespaceCode = e.EnvironmentCode
// 	data.ClusterCode = e.ClusterCode
// 	if data.Kind == "client" {
// 		if data.GatewayKind == "mesh_gateway" {
// 			data.AggHost = data.Remote
// 		} else {
// 			data.AggHost = data.Local
// 		}
// 	}
// 	// 162,262,498,785,5722
// 	data.Timestamp /= 1000
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// 	defer cancel()
// 	_, err = e.esCli.Index().Index(e.indexEnvoyAccess).BodyJson(data).Do(ctx)
// 	return err
// }

// func (e *EsClient) dealTraceLog(m *Message) error {
// 	data := new(logs.Trace)
// 	err := json.Unmarshal(m.Value, data)
// 	if err != nil {
// 		return err
// 	}
// 	data.Tenant = e.EnvironmentCode
// 	data.EnvironmentCode = e.EnvironmentCode
// 	data.NamespaceCode = e.EnvironmentCode
// 	data.ClusterCode = e.ClusterCode
// 	// 162,262,594,083,1145
// 	data.Timestamp /= 1000
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// 	defer cancel()
// 	_, err = e.esCli.Index().Index(e.indexTrace).BodyJson(data).Do(ctx)
// 	return err
// }

// func (e *EsClient) BuildMessage(m *Message) (index string, data interface{}) {
// 	indexEnvoyAccess := fmt.Sprintf("envoy_access-%s", time.Now().Format("2006-01-02"))
// 	indexTrace := fmt.Sprintf("trace-%s", time.Now().Format("2006-01-02"))
// 	if e.indexEnvoyAccess != indexEnvoyAccess {
// 		var err error
// 		err = e.CreateIndex(indexEnvoyAccess)
// 		if err == nil {
// 			err = e.CreateIndex(indexTrace)
// 			if err == nil {
// 				e.indexEnvoyAccess = indexEnvoyAccess
// 				e.indexTrace = indexTrace
// 				e.l.Infof("envoy_access log index has changed: %s", e.indexEnvoyAccess)
// 				e.l.Infof("trace log index has changed: %s", e.indexTrace)
// 				continue
// 			}
// 		}
// 		e.l.Errorf("failed to create index %s: %v", indexEnvoyAccess, err)
// 	return "", nil
// }

// func (e *EsClient) CreateIndex(index string) error {
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// 	defer cancel()
// 	_, err := e.esCli.CreateIndex(index).Do(ctx)
// 	return err
// }
