package confdserver

import (
	"context"
	confdv1 "grape/api/v1/confd"
	"grape/internal/share"
	"grape/pkg/etcdcli"
	"grape/pkg/logger"
	"sync"
	"time"

	"go.etcd.io/etcd/api/v3/mvccpb"
)

type watcher struct {
	cli   *etcdcli.Client
	l     sync.RWMutex
	chans map[string]map[chan<- *confdv1.Configs]string
}

func (w *watcher) watchLoop(ctx context.Context, log logger.Logger) {
	var handle = func(kv *mvccpb.KeyValue) {
		k, v := kv.Key, kv.Value
		project, err := UnmarshalProject(v)
		if err != nil {
			log.Errorf("Unmarshal project err: %v, value: %s", err, string(v))
		}
		go func() {
			w.l.RLock()
			defer w.l.RUnlock()

			if cs, ok := w.chans[string(k)]; ok {
				for c, group := range cs {
					if project.NodeUpdateInterval > 0 {
						<-time.After(time.Duration(project.NodeUpdateInterval) * time.Second)
					}
					c <- GetGroupConfig(project, group)
				}
			}
		}()
	}
	w.cli.WatchPrefixEvents(ctx, share.ServerKeyPrefix, log, handle, handle, nil)
}

func (w *watcher) notify(key, group string, c chan<- *confdv1.Configs) {
	w.l.Lock()
	defer w.l.Unlock()
	m := w.chans[key]
	if m == nil {
		w.chans[key] = map[chan<- *confdv1.Configs]string{c: group}
	} else {
		m[c] = group
	}
}

func (w *watcher) stop(key string, c chan<- *confdv1.Configs) {
	w.l.Lock()
	defer w.l.Unlock()
	delete(w.chans[key], c)
}
