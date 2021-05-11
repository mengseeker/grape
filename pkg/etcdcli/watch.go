package etcdcli

import (
	"context"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func WatchKey(c context.Context, key string, call func(*clientv3.WatchResponse)) {
	watchChan := cli.Watch(c, key)
	for watchResp := range watchChan {
		call(&watchResp)
	}
}

func WatchPrefix(c context.Context, key string, call func(*clientv3.WatchResponse)) {
	watchChan := cli.Watch(c, key, clientv3.WithPrefix())
	for watchResp := range watchChan {
		call(&watchResp)
	}
}
