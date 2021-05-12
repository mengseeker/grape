package etcdcli

import (
	"context"

	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func WatchKey(c context.Context, key string, call func(*clientv3.WatchResponse)) {
	watchChan := cli.Watch(c, key)
	for watchResp := range watchChan {
		watchBackCaller(&watchResp, key, call)
	}
}

func WatchPrefix(c context.Context, key string, call func(*clientv3.WatchResponse)) {
	watchChan := cli.Watch(c, key, clientv3.WithPrefix())
	for watchResp := range watchChan {
		watchBackCaller(&watchResp, key, call)
	}
}

func watchBackCaller(watchResp *clientv3.WatchResponse, key string, call func(*clientv3.WatchResponse)) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorf("watch %s backcall err: %v", err)
		}
	}()
	call(watchResp)
}

func PrefixAll(ctx context.Context, key string, call func([]*mvccpb.KeyValue)) error {
	resp, err := cli.Get(ctx, key,
		clientv3.WithPrefix(),
		clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend),
		// clientv3.WithLimit(1),
	)
	if err != nil {
		return err
	}
	call(resp.Kvs)
	for resp.More {
		resp, err = cli.Get(ctx, key,
			clientv3.WithRange(string(resp.Kvs[len(resp.Kvs)-1].Key)),
			clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend),
			// clientv3.WithLimit(1),
		)
		if err != nil {
			return err
		}
		call(resp.Kvs)
	}
	return nil
}
