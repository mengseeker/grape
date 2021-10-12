package viewserver

import (
	"context"
	"encoding/json"
	"grape/api/v1/view"
	"grape/internal/iutils"
	"grape/pkg/etcdcli"
	"grape/pkg/logger"
	"sync"

	clientv3 "go.etcd.io/etcd/client/v3"
)

type server struct {
	log   logger.Logger
	cli   *etcdcli.Client
	w     map[chan<- *view.AllLogs]bool
	last  *view.AllLogs
	wlock sync.Mutex
	view.UnimplementedDiscoveryServerServer
}

func NewServer(log logger.Logger, cli *etcdcli.Client) *server {
	s := &server{
		log:   log,
		cli:   cli,
		w:     map[chan<- *view.AllLogs]bool{},
		wlock: sync.Mutex{},
		last:  &view.AllLogs{},
	}
	go watchLoop(s, cli)
	return s
}

func (s *server) StreamLogs(discovery *view.DiscoveryRequest, stream view.DiscoveryServer_StreamLogsServer) error {
	s.log.Infof("discovery servicelogs")
	logsChan := make(chan *view.AllLogs, 1)
	defer close(logsChan)
	s.watch(logsChan)
	defer s.stop(logsChan)
	logsChan <- s.last
	for slogs := range logsChan {
		err := stream.Send(slogs)
		if err != nil {
			s.log.Errorf("unable to send servicelogs: %v", err)
			return err
		}
	}
	return nil
}

func (s *server) watch(c chan<- *view.AllLogs) {
	s.wlock.Lock()
	defer s.wlock.Unlock()
	s.w[c] = true
}

func (s *server) stop(c chan<- *view.AllLogs) {
	s.wlock.Lock()
	defer s.wlock.Unlock()
	delete(s.w, c)
}

func watchLoop(s *server, cli *etcdcli.Client) {
	logs := map[string]*view.ServiceLogs{}
	version := iutils.NewVersion()
	var handle = func(resp *clientv3.WatchResponse) {
		err := resp.Err()
		if err != nil {
			s.log.Errorf("watch %s err: %v", ServerKeyPrefix, err)
			return
		}
		for _, event := range resp.Events {
			switch event.Type {
			case clientv3.EventTypeDelete:
				delete(logs, string(event.Kv.Key))
				version = iutils.NewVersion()
			case clientv3.EventTypePut:
				val := event.Kv.Value
				slogs := view.ServiceLogs{}
				err := json.Unmarshal(val, &slogs)
				if err != nil {
					s.log.Errorf("unable to unmarshal serverlogs: %v", err)
					continue
				}
				logs[string(event.Kv.Key)] = &slogs
				if slogs.Version > version {
					version = slogs.Version
				}
			}
		}
		s.wlock.Lock()
		defer s.wlock.Unlock()
		// update
		s.last.Logs = []*view.ServiceLogs{}
		for _, slogs := range logs {
			s.last.Logs = append(s.last.Logs, slogs)
		}
		for lc := range s.w {
			lc <- s.last
		}
	}
	cli.WatchPrefix(context.Background(), ServerKeyPrefix, handle)
}
