package worker

import (
	"context"
	"encoding/json"
	"grape/logtrans/logs"
	"grape/pkg/logger"
	"strings"
	"time"

	"github.com/olivere/elastic/v7"
)

type EsClient struct {
	// Addr  string
	esCli           *elastic.Client
	ClusterCode     string
	EnvironmentCode string
	l               logger.Logger
	lastDate        string
}

func NewEsClient(addr, env, cluster string, l logger.Logger) (*EsClient, error) {
	urls := strings.Split(addr, ",")
	l.Debugf("connecting to es(%s)....", addr)
	cli, err := elastic.NewClient(elastic.SetURL(urls...), elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}
	version, err := cli.ElasticsearchVersion(urls[0])
	l.Debugf("es(%s) connected", version)
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

func (e *EsClient) Write(ms []*Message) {
	defer func() {
		if rerr := recover(); rerr != nil {
			e.l.Errorf("write message to es panic: %v", rerr)
		}
	}()
	if len(ms) == 0 {
		return
	}
	envoyIndex, traceIndex := e.GetEsIndex()
	bulk := e.esCli.Bulk()
	var data interface{}
	var index = unKnowIndex
	count := 0
	for _, m := range ms {
		logType := GetLogType(m)
		if logType == logTypeEnvoyAccess {
			data = e.dealEnvoyAccessLog(m)
			index = envoyIndex
			count++
		} else if logType == logTypeTrace {
			data = e.dealTraceLog(m)
			index = traceIndex
			count++
		} else {
			e.l.Warnf("logType %s undefined", logType)
			data = map[string]string{"raw": string(m.Value)}
		}
		bulk.Add(
			elastic.NewBulkCreateRequest().
				UseEasyJSON(true).
				Index(index).
				Doc(data),
		)
	}
	ctx, cancel := context.WithTimeout(context.Background(), esTimeout)
	defer cancel()
	_, err := bulk.Do(ctx)
	if err != nil {
		e.l.Errorf("write logs to es err: %v", err)
	} else {
		e.l.Debugf("write logs to elasticsearch count: %d", count)
	}
}

func (e *EsClient) GetEsIndex() (string, string) {
	date := time.Now().Format("2006-01-02")
	indexEnvoyAccess := "envoy_access-" + date
	indexTrace := "trace-" + date
	if e.lastDate != date {
		err := e.CreateIndex(indexEnvoyAccess)
		if err != nil {
			e.l.Errorf("create index %s err: %v", indexEnvoyAccess, err)
		}
		err = e.CreateIndex(indexTrace)
		if err != nil {
			e.l.Errorf("create index %s err: %v", indexTrace, err)
		}
		e.lastDate = date
	}
	return indexEnvoyAccess, indexTrace
}

func (e *EsClient) dealEnvoyAccessLog(m *Message) *logs.EnvoyAccess {
	data := new(logs.EnvoyAccess)
	err := json.Unmarshal(m.Value, data)
	if err != nil {
		e.l.Error(string(m.Value))
		e.l.Errorf("unmarshal envoyAccess log err: %v", err)
	}
	data.Tenant = e.EnvironmentCode
	data.EnvironmentCode = e.EnvironmentCode
	data.NamespaceCode = e.EnvironmentCode
	data.ClusterCode = e.ClusterCode
	if data.Kind == "client" {
		if data.GatewayKind == "mesh_gateway" {
			data.AggHost = data.Remote
		} else {
			data.AggHost = data.Local
		}
	}
	data.Timestamp /= 1000
	return data
}

func (e *EsClient) dealTraceLog(m *Message) *logs.Trace {
	data := new(logs.Trace)
	err := json.Unmarshal(m.Value, data)
	if err != nil {
		e.l.Error(string(m.Value))
		e.l.Errorf("unmarshal Trace log err: %v", err)
	}
	data.Tenant = e.EnvironmentCode
	data.EnvironmentCode = e.EnvironmentCode
	data.NamespaceCode = e.EnvironmentCode
	data.ClusterCode = e.ClusterCode
	// 162,262,594,083,1145
	data.Timestamp /= 1000
	return data
}

func (e *EsClient) CreateIndex(index string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := e.esCli.CreateIndex(index).Do(ctx)
	return err
}
