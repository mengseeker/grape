package worker

import (
	"grape/pkg/logger"
	"time"

	_ "github.com/influxdata/influxdb1-client"
	influxdb "github.com/influxdata/influxdb1-client/v2"
)

type InfClient struct {
	ClusterCode     string
	EnvironmentCode string

	infCli influxdb.Client
	l      logger.Logger
}

func NewInfClient(addr, env, cluster string, l logger.Logger) (*InfClient, error) {
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
	cli := new(InfClient)
	cli.ClusterCode = cluster
	cli.EnvironmentCode = env
	cli.l = l
	cli.infCli = c
	return cli, nil
}

func (e *InfClient) Write([]*Message) {

}
