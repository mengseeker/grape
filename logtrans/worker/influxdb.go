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
	allKeys            = map[string]bool{
		"access_time": true, "kind": true, "gateway_kind": true, "gateway_code": true, "response_flags": true, "timestamp": true, "remote": true, "local": true, "method": true, "path": true, "code": true, "status": true, "req_size": true, "res_size": true, "traffic": true, "request_time": true, "referer": true, "agent": true, "forwoad": true, "request": true, "response": true, "uuid": true, "app_id": true, "health_level": true, "cr_service_code": true, "sr_service_code": true, "cr_group_code": true, "sr_group_code": true, "upstream_host": true, "upstream_cluster": true,
	}
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
	count := 0
	for _, m := range ms {
		if GetLogType(m) == logTypeEnvoyAccess {
			point := e.BuildPoint(m)
			if point != nil {
				bs.AddPoint(point)
				count++
			}
		}
	}
	err = e.infCli.Write(bs)
	if err != nil {
		e.l.Errorf("faild to write to influxdb: %v", err)
	} else {
		e.l.Debugf("write logs to influxdb count: %d", count)
	}
}

func (e *InfClient) BuildPoint(m *Message) *influxdb.Point {
	var tags = make(map[string]string, 8)
	var fields = make(map[string]interface{}, 8)
	err := json.Unmarshal(m.Value, &fields)
	if err != nil {
		e.l.Error(string(m.Value))
		e.l.Errorf("unmarshal envoyAccess log err: %v", err)
		return nil
	}
	// 存在一些异常日志数据，导致influxdb fields非常多
	// 这里将只保留正常的字段
	for key := range fields {
		if !allKeys[key] {
			delete(fields, key)
		}
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
