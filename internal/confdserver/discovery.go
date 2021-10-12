package confdserver

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"grape/api/v1/confd"
	"grape/pkg/etcdcli"
	"grape/pkg/logger"
	"sync"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
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
		err := stream.Send(config)
		if err != nil {
			s.log.Errorf("unable to send configs: %v", err)
			return err
		}
	}
	return nil
}

func (s *server) FirstLoadConfig(discovery *confd.Discovery, configChan chan<- *confd.Configs) {
	cf, err := s.GetServiceConfigs(discovery.Service, discovery.Group, 0)
	if err != nil {
		s.log.Errorf("unable to get service configs %v", err)
		return
	}
	configChan <- cf
}

func (s *server) Download(ctx context.Context, req *confd.DownloadRequest) (*confd.DownloadResponse, error) {
	cf, err := s.GetServiceConfigs(req.Service, req.Group, req.LoadVersion)
	if err != nil {
		return nil, err
	}
	return &confd.DownloadResponse{Code: OkCode, Configs: cf}, nil
}

func (s *server) GetServiceConfigs(service, group string, loadVersion int64) (*confd.Configs, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	key := Key(service)
	ops := []clientv3.OpOption{}
	if loadVersion != 0 {
		ops = append(ops, clientv3.WithRev(loadVersion))
	}
	resp, err := s.w.cli.Cli.Get(ctx, key, ops...)
	if err != nil {
		s.log.Errorf("unalble to get confis form etcd: %v", err)
		return nil, err
	}
	if resp.Count == 0 {
		return nil, fmt.Errorf("service configs %v not found", service)
	}
	sf, err := UnmarshalServiceConfig(resp.Kvs[0].Value)
	if err != nil {
		return nil, fmt.Errorf("unmarshal configs err: %v", err)
	}
	return GetGroupConfig(sf, group), nil
}

func UnmarshalServiceConfig(raw []byte) (*confd.ServerConfig, error) {
	config := confd.ServerConfig{}
	err := json.Unmarshal(raw, &config)
	return &config, err
}

func MarshalServiceConfig(config *confd.ServerConfig) []byte {
	j, _ := json.Marshal(config)
	return j
}

func GetGroupConfig(config *confd.ServerConfig, group string) *confd.Configs {
	if cf, ok := config.GroupConfigs[group]; ok {
		return cf
	}
	return config.Default
}

func Key(service string) string {
	return ServerKeyPrefix + service
}
