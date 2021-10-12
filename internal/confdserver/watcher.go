package confdserver

import (
	"context"
	"grape/api/v1/confd"
	"grape/pkg/etcdcli"
	"grape/pkg/logger"
	"sync"

	"go.etcd.io/etcd/api/v3/mvccpb"
)

type watcher struct {
	cli   *etcdcli.Client
	l     sync.Mutex
	chans map[string]map[chan<- *confd.Configs]string
}

func (w *watcher) watchLoop(ctx context.Context, log logger.Logger) {
	var handle = func(kv *mvccpb.KeyValue) {
		k, v := kv.Key, kv.Value
		sf, err := UnmarshalServiceConfig(v)
		if err != nil {
			log.Errorf("Unmarshal serverConfigs err: %v, value: %s", err, string(v))
		}
		if cs, ok := w.chans[string(k)]; ok {
			for c, group := range cs {
				c <- GetGroupConfig(sf, group)
			}
		}
	}
	w.cli.WatchPrefixEvents(ctx, ServerKeyPrefix, log, handle, handle, nil)
}

func (w *watcher) notify(key, group string, c chan<- *confd.Configs) {
	w.l.Lock()
	defer w.l.Unlock()
	m := w.chans[key]
	if m == nil {
		w.chans[key] = map[chan<- *confd.Configs]string{c: group}
	} else {
		m[c] = group
	}
}

func (w *watcher) stop(key string, c chan<- *confd.Configs) {
	w.l.Lock()
	defer w.l.Unlock()
	delete(w.chans[key], c)
}
