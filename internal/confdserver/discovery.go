package confdserver

import (
	"context"
	"errors"
	confdv1 "grape/api/v1/confd"
	"grape/pkg/etcdcli"
	"grape/pkg/logger"
	"sync"
)

type server struct {
	confdv1.UnimplementedConfdServerServer
	log logger.Logger
	w   *watcher
}

func NewServer(log logger.Logger, cli *etcdcli.Client) *server {
	ctx := context.Background()
	watch := &watcher{
		cli:   cli,
		l:     sync.RWMutex{},
		chans: map[string]map[chan<- *confdv1.Configs]string{},
	}
	go watch.watchLoop(ctx, log)
	return &server{
		log: log,
		w:   watch,
	}
}

func (s *server) StreamDiscovery(discovery *confdv1.Discovery, stream confdv1.ConfdServer_StreamDiscoveryServer) error {
	s.log.Infof("discovery from %s:%s", discovery.ProjectName, discovery.Group)
	key := Key(discovery.ProjectName)

	if discovery.ProjectName == "" {
		return errors.New("empty projectName")
	}

	configChan := make(chan *confdv1.Configs)
	defer close(configChan)

	config, err := GetProjectConfigs(s.w.cli, discovery.ProjectName, discovery.Group)
	if err != nil {
		return err
	}
	go func() { configChan <- config }()

	s.w.notify(key, discovery.Group, configChan)
	defer s.w.stop(key, configChan)

	for config := range configChan {
		if config != nil {
			err := stream.Send(config)
			if err != nil {
				s.log.Errorf("unable to send configs: %v", err)
				return err
			}
		}
	}
	return nil
}
