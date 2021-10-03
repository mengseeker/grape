// 启动时，初始化数据
// 配置watch handle
package auth

import (
	"context"
	"grape/internal/share"
	"time"

	"go.etcd.io/etcd/api/v3/mvccpb"
)

func initConfig(clusterCode string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	var err error
	key := share.AuthAppDBKey(clusterCode)
	log.Infof("loading %s for apps", key)
	err = cli.PrefixAll(ctx, key, func(resp []*mvccpb.KeyValue) {
		for _, kv := range resp {
			SetupApp(kv)
		}
	})
	if err != nil {
		log.Fatalf("load apps err: %v", err)
	}

	key = share.AuthTokenDBKey(clusterCode)
	log.Infof("loading %s for tokens", key)
	err = cli.PrefixAll(ctx, key, func(resp []*mvccpb.KeyValue) {
		for _, kv := range resp {
			SetupToken(kv)
		}
	})
	if err != nil {
		log.Fatalf("load tokens err: %v", err)
	}
}

func watchApp(clusterCode string) {
	key := share.AuthAppDBKey(clusterCode)
	cli.WatchPrefixEvents(context.Background(), key, log, SetupApp, UpdateApp, RemoveApp)
}

func watchToken(clusterCode string) {
	key := share.AuthTokenDBKey(clusterCode)
	cli.WatchPrefixEvents(context.Background(), key, log, SetupToken, nil, RemoveToken)
}
