package etcdcli

import (
	"context"
	"grape/pkg/logger"
	"strings"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

var (
	cli *clientv3.Client
)

func Connect(addrs string) error {
	var err error
	ends := strings.Split(addrs, ",")
	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   ends,
		DialTimeout: 5 * time.Second,
		LogConfig:   logger.LoggerCfg(),
	})
	if err != nil {
		return err
	}
	ctx, calcel := context.WithTimeout(context.Background(), time.Second*3)
	defer calcel()
	_, err = cli.Get(ctx, "test")
	return err
}

func Cli() *clientv3.Client {
	return cli
}
