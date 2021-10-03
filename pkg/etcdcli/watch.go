package etcdcli

import (
	"context"
	"grape/pkg/logger"

	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type EventHandle func(kv *mvccpb.KeyValue)

func (cli *Client) WatchKey(c context.Context, key string, call func(*clientv3.WatchResponse)) {
	watchChan := cli.Cli.Watch(c, key)
	for watchResp := range watchChan {
		watchBackCaller(&watchResp, key, call)
	}
}

func (cli *Client) WatchPrefix(c context.Context, key string, call func(*clientv3.WatchResponse)) {
	watchChan := cli.Cli.Watch(c, key, clientv3.WithPrefix())
	for watchResp := range watchChan {
		watchBackCaller(&watchResp, key, call)
	}
}

func (cli *Client) WatchPrefixEvents(
	ctx context.Context, key string,
	log logger.Logger,
	CreateHandle, UpdateHandle, RemoveHandle EventHandle,
) {
	log.Infof("watching prefix %s", key)
	cli.WatchPrefix(ctx, key, func(resp *clientv3.WatchResponse) {
		err := resp.Err()
		if err != nil {
			log.Errorf("watch %s err: %v", key, err)
		} else {
			for _, event := range resp.Events {
				func() {
					defer func() {
						if err := recover(); err != nil {
							log.Errorf("watch %s backcall panic: %v", key, err)
						}
					}()
					var bk EventHandle
					switch event.Type {
					case clientv3.EventTypeDelete:
						bk = RemoveHandle
					case clientv3.EventTypePut:
						if event.IsCreate() {
							bk = CreateHandle
						} else {
							bk = UpdateHandle
						}
					}
					if bk != nil {
						bk(event.Kv)
					}
				}()
			}
		}
	})
}

func watchBackCaller(watchResp *clientv3.WatchResponse, key string, call func(*clientv3.WatchResponse)) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorf("watch %s backcall panic: %v", key, err)
		}
	}()
	call(watchResp)
}

func (cli *Client) PrefixAll(ctx context.Context, key string, call func([]*mvccpb.KeyValue)) error {
	resp, err := cli.Cli.Get(ctx, key,
		clientv3.WithPrefix(),
		clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend),
		// clientv3.WithLimit(1),
	)
	if err != nil {
		return err
	}
	call(resp.Kvs)
	for resp.More {
		resp, err = cli.Cli.Get(ctx, key,
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
