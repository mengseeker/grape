package confdserver

import (
	"context"
	"encoding/json"
	"grape/api/confd"
	"grape/pkg/etcdcli"
	"grape/pkg/logger"
	"sync"

	"go.etcd.io/etcd/api/v3/mvccpb"
)

type watcher struct {
	cli   *etcdcli.Client
	l     sync.Mutex
	chans map[string]map[chan<- *confd.Configs]bool
}

func (w *watcher) watchLoop(ctx context.Context, log logger.Logger) {
	var handle = func(k, v []byte) {
		if cs, ok := w.chans[string(k)]; ok {
			config := confd.Configs{}
			err := json.Unmarshal(v, &config)
			if err != nil {
				log.Errorf("Unmarshal configs err: %v", err)
			} else {
				for c := range cs {
					c <- &config
				}
			}
		}
	}
	w.cli.WatchPrefixEvents(ctx, ServerKeyPrefix, log,
		func(kv *mvccpb.KeyValue) { handle(kv.Key, kv.Value) },
		func(kv *mvccpb.KeyValue) { handle(kv.Key, kv.Value) },
		nil,
	)
}

func (w *watcher) notify(key string, c chan<- *confd.Configs) {
	w.l.Lock()
	defer w.l.Unlock()
	m := w.chans[key]
	if m == nil {
		w.chans[key] = map[chan<- *confd.Configs]bool{c: true}
	} else {
		m[c] = true
	}
}

func (w *watcher) stop(key string, c chan<- *confd.Configs) {
	w.l.Lock()
	defer w.l.Unlock()
	delete(w.chans[key], c)
}
