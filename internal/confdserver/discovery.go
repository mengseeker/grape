package confdserver

import (
	"context"
	"grape/api/confd"
	"grape/pkg/etcdcli"
	"grape/pkg/logger"
	"sync"
	"time"

	"google.golang.org/protobuf/proto"
)

const (
	ServerKeyPrefix = "configs/"
)

type server struct {
	confd.UnimplementedConfdServerServer
	log logger.Logger
	w   *watcher
}

func NewServer(log logger.Logger, cli *etcdcli.Client) *server {
	ctx := context.Background()
	watch := &watcher{
		cli:   cli,
		l:     sync.Mutex{},
		chans: map[string]map[chan<- *confd.Configs]bool{},
	}
	go watch.watchLoop(ctx, log)
	return &server{
		log: log,
		w:   watch,
	}
}

func (s *server) StreamResources(discovery *confd.Discovery, stream confd.ConfdServer_StreamResourcesServer) error {
	s.log.Infof("discovery from %s(%s)", discovery.Service, discovery.Node.String())
	configChan := make(chan *confd.Configs, 1)
	defer close(configChan)
	key := Key(discovery.Service)
	// first send configs
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	resp, err := s.w.cli.Cli.Get(ctx, key)
	if err != nil {
		s.log.Errorf("unalble to get confis form etcd: %v", err)
		return err
	} else {
		config := confd.Configs{}
		if resp.Count == 1 {
			err := proto.Unmarshal(resp.Kvs[0].Value, &config)
			if err != nil {
				s.log.Errorf("Unmarshal configs err: %v", err)
			} else {
				configChan <- &config
			}
		} else {
			// no service configs
			configChan <- &config
		}
	}
	s.w.notify(key, configChan)
	defer s.w.stop(key, configChan)
	for config := range configChan {
		err := stream.Send(config)
		if err != nil {
			s.log.Errorf("unable to send configs: %v", err)
			return err
		}
	}
	return nil
}

func Key(service string) string {
	return ServerKeyPrefix + service
}
