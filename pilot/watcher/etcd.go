package watcher

import "grape/pkg/etcdcli"

var (
	cli *etcdcli.Client
)

func ConnectEtcd(addr string) error {
	var err error
	cli, err = etcdcli.Connect(addr)
	return err
}
