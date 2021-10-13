package confdserver

import (
	"context"
	"errors"
	"grape/api/v1/confd"
	"grape/pkg/etcdcli"
	"grape/pkg/logger"
	"sync"
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
		chans: map[string]map[chan<- *confd.Configs]string{},
	}
	go watch.watchLoop(ctx, log)
	return &server{
		log: log,
		w:   watch,
	}
}

func (s *server) StreamResources(discovery *confd.Discovery, stream confd.ConfdServer_StreamResourcesServer) error {
	s.log.Infof("discovery from %s:%s", discovery.Service, discovery.Group)
	if discovery.Service == "" {
		return errors.New("bad discovery service")
	}
	configChan := make(chan *confd.Configs, 1)
	defer close(configChan)
	s.FirstLoadConfig(discovery, configChan)
	key := Key(discovery.Service)
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

func (s *server) FirstLoadConfig(discovery *confd.Discovery, configChan chan<- *confd.Configs) {
	cf, err := GetLatestServiceConfigs(s.w.cli, discovery.Service, discovery.Group)
	if err != nil {
		s.log.Errorf("unable to get service configs %v", err)
		return
	}
	configChan <- cf
}

func (s *server) Download(ctx context.Context, req *confd.DownloadRequest) (*confd.DownloadResponse, error) {
	cf, err := GetRevServiceConfigs(s.w.cli, req.Service, req.Group, req.LoadVersion)
	if err != nil {
		return nil, err
	}
	return &confd.DownloadResponse{Code: OkCode, Configs: cf}, nil
}
