package watcher

// import (
// 	"context"
// 	"grape/pkg/share"
// 	"time"

// 	"go.etcd.io/etcd/api/v3/mvccpb"
// )

// func initConfig(clusterCode string) {
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
// 	defer cancel()
// 	var err error
// 	key := share.ResourceServicePrefix(clusterCode)
// 	log.Infof("loading %s for services", key)
// 	err = cli.PrefixAll(ctx, key, func(resp []*mvccpb.KeyValue) {
// 		for _, kv := range resp {
// 			SetupService(kv)
// 		}
// 	})
// 	if err != nil {
// 		log.Fatalf("load apps err: %v", err)
// 	}
// }

// func watchService(clusterCode string) {
// 	key := share.ResourceServicePrefix(clusterCode)
// 	ctx := context.Background()
// 	cli.WatchPrefixEvents(ctx, key, log, SetupService, UpdateService, RemoveService)
// }
