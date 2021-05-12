// 启动时，初始化数据
// 配置watch handle
package auth

import (
	"context"
	"grape/pkg/etcdcli"
	"grape/pkg/share"
	"time"

	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func initConfig() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	var err error
	err = etcdcli.PrefixAll(ctx, share.AppDB, func(resp []*mvccpb.KeyValue) {
		for _, kv := range resp {
			err = SetupApp(kv)
			if err != nil {
				log.Fatalf("setup apps err: %v", err)
			}
		}
	})
	if err != nil {
		log.Fatalf("load apps err: %v", err)
	}

	err = etcdcli.PrefixAll(ctx, share.TokenDB, func(resp []*mvccpb.KeyValue) {
		for _, kv := range resp {
			err = SetupToken(kv)
			if err != nil {
				log.Fatalf("setup token err: %v", err)
			}
		}
	})
	if err != nil {
		log.Fatalf("load tokens err: %v", err)
	}
}

func watchApp() {
	etcdcli.WatchPrefix(context.Background(), share.AppDB, func(resp *clientv3.WatchResponse) {
		err := resp.Err()
		if err != nil {
			log.Errorf("watch apps err: %v", err)
		} else {
			for _, event := range resp.Events {
				var err error
				switch event.Type {
				case clientv3.EventTypeDelete:
					err = RemoveApp(event.Kv)
				case clientv3.EventTypePut:
					if event.IsCreate() {
						err = SetupApp(event.Kv)
					} else {
						err = UpdateApp(event.Kv)
					}
				}
				if err != nil {
					log.Errorf("apply app event %s err: %v", event.Type.String(), err)
				}
			}
		}
	})
}

func watchToken() {
	etcdcli.WatchPrefix(context.Background(), share.TokenDB, func(resp *clientv3.WatchResponse) {
		err := resp.Err()
		if err != nil {
			log.Errorf("watch token err: %v", err)
		} else {
			for _, event := range resp.Events {
				var err error
				switch event.Type {
				case clientv3.EventTypeDelete:
					err = RemoveToken(event.Kv)
				case clientv3.EventTypePut:
					if event.IsCreate() {
						err = SetupToken(event.Kv)
					} else {
						err = UpdateToken(event.Kv)
					}
				}
				if err != nil {
					log.Errorf("apply token event %s err: %v", event.Type.String(), err)
				}
			}
		}
	})
}
