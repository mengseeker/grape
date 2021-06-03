package worker

import (
	"encoding/json"
	"grape/pkg/logger"
	"strconv"
	"time"

	_ "github.com/influxdata/influxdb1-client"
	influxdb "github.com/influxdata/influxdb1-client/v2"
)

var (
	InfluxRemoveFields = [...]string{"agent", "access_time", "code"}
	stringTags         = [...]string{"path", "gateway_kind", "method", "health_level", "health_level", "cr_service_code", "sr_service_code", "app_id"}
)

type InfClient struct {
	ClusterCode     string
	EnvironmentCode string

	infCli influxdb.Client
	l      logger.Logger
	nextID int64
}

func NewInfClient(addr, env, cluster string, l logger.Logger) (*InfClient, error) {
	l.Debugf("connecting to influxdb(%s)....", addr)
	c, err := influxdb.NewHTTPClient(influxdb.HTTPConfig{
		Addr: addr,
	})
	if err != nil {
		return nil, err
	}
	_, _, err = c.Ping(3 * time.Second)
	if err != nil {
		return nil, err
	}
	l.Debugf("influxdb connected")
	cli := new(InfClient)
	cli.ClusterCode = cluster
	cli.EnvironmentCode = env
	cli.l = l
	cli.infCli = c
	return cli, nil
}

func (e *InfClient) Write(ms []*Message) {
	defer func() {
		if rerr := recover(); rerr != nil {
			e.l.Errorf("write message to influxdb panic: %v", rerr)
		}
	}()
	if len(ms) == 0 {
		return
	}
	if len(ms) == 0 {
		return
	}
	bc := influxdb.BatchPointsConfig{
		Precision: "ns",
		Database:  influxDatabase,
	}
	bs, err := influxdb.NewBatchPoints(bc)
	if err != nil {
		e.l.Errorf("can not create influxdb batchPoints: %v", err)
		return
	}
	for _, m := range ms {
		if GetLogType(m) == logTypeEnvoyAccess {
			bs.AddPoint(e.BuildPoint(m))
		}
	}
	err = e.infCli.Write(bs)
	if err != nil {
		e.l.Errorf("faild to write to influxdb: %v", err)
	} else {
		e.l.Debugf("write logs to influxdb %d", len(ms))
	}
}

func (e *InfClient) BuildPoint(m *Message) *influxdb.Point {
	var tags = make(map[string]string, 8)
	var fields = make(map[string]interface{}, 8)
	err := json.Unmarshal(m.Value, &fields)
	if err != nil {
		e.l.Error(string(m.Value))
		e.l.Errorf("unmarshal envoyAccess log err: %v", err)
	}
	accessTime := int64(fields["timestamp"].(float64))
	timestamp := time.Unix(accessTime/1000_000, e.GetNextID(accessTime))
	tags["tenant"] = e.EnvironmentCode
	tags["enviroment_code"] = e.EnvironmentCode
	tags["cluster_code"] = e.ClusterCode

	fields["namespace_code"] = e.EnvironmentCode

	tags["code"] = strconv.FormatInt(int64(fields["code"].(float64)), 10)

	for _, t := range stringTags {
		if v, ok := fields[t]; ok {
			tags[t] = v.(string)
			delete(fields, t)
		}
	}

	for _, f := range InfluxRemoveFields {
		delete(fields, f)
	}

	point, err := influxdb.NewPoint(influxMeasurement, tags, fields, timestamp)
	if err != nil {
		e.l.Errorf("build Point err: %v", err)
	}
	return point
}

// 只在一个线程内跑，不用加锁
// 0 < id < 1000_000_000
func (e *InfClient) GetNextID(timestamp16 int64) int64 {
	e.nextID = (e.nextID + 1) % 1000
	return (timestamp16*1000 + e.nextID) % 1000_000_000
}
