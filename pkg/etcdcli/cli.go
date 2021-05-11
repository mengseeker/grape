package etcdcli

import (
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
	})
	if err != nil {
		return err
	}
	return nil
}

func Cli() *clientv3.Client {
	return cli
}
