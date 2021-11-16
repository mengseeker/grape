package etcdcli

import (
	"context"
	"grape/pkg/logger"
	"strings"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

type Client struct {
	Cli         *clientv3.Client
	ClusterCode string
}

func Connect(addrs, username, password string) (*Client, error) {
	ends := strings.Split(addrs, ",")
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   ends,
		DialTimeout: 5 * time.Second,
		LogConfig:   logger.LoggerCfg(),
		Username:    username,
		Password:    password,
	})
	if err != nil {
		return nil, err
	}
	ctx, calcel := context.WithTimeout(context.Background(), time.Second*3)
	defer calcel()
	_, err = cli.Get(ctx, "test")
	return &Client{Cli: cli}, err
}
