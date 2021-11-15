package confdserver

import (
	"context"
	"fmt"
	"grape/api/v1/confd"
	"grape/pkg/etcdcli"
	"grape/pkg/logger"
	"time"
)

type apiserver struct {
	cli *etcdcli.Client
	log logger.Logger
	confd.UnimplementedApiServerServer
}

func NewApiServer(log logger.Logger, cli *etcdcli.Client) *apiserver {
	return &apiserver{cli: cli, log: log}
}

func (s *apiserver) Set(ctx context.Context, req *confd.SetRequest) (*confd.SetResponse, error) {
	s.log.Infof("set %s/%s", req.ServerConfig.Namespace, req.ServerConfig.Service)
	key := Key(req.ServerConfig.Namespace, req.ServerConfig.Service)
	timeout, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	val := MarshalServiceConfig(req.ServerConfig)
	_, err := s.cli.Cli.KV.Put(timeout, key, string(val))
	return &confd.SetResponse{}, err
}

func (s *apiserver) Get(ctx context.Context, req *confd.GetRequest) (*confd.GetResponse, error) {
	s.log.Infof("get %s/%s", req.Namespace, req.Service)
	key := Key(req.Namespace, req.Service)
	timeout, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	resp, err := s.cli.Cli.Get(timeout, key)
	if err != nil {
		return nil, err
	}
	if resp.Count == 0 {
		return &confd.GetResponse{Code: ServiceConfigNotFoundCode, Message: "resource not found"}, nil
	}
	sf, err := UnmarshalServiceConfig(resp.Kvs[0].Value)
	if err != nil {
		s.log.Errorf("unable to unmarshal serverConfigs: %v, value: %s", err, resp.Kvs[0].Value)
		return nil, fmt.Errorf("unable to unmarshal serverConfigs: %v", err)
	}
	return &confd.GetResponse{ServerConfig: sf}, nil
}

func (s *apiserver) Del(ctx context.Context, req *confd.DelRequest) (*confd.DelResponse, error) {
	s.log.Infof("del %s configs", req.Service)
	key := Key(req.Namespace, req.Service)
	timeout, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	_, err := s.cli.Cli.Delete(timeout, key)
	if err != nil {
		return nil, err
	}
	return &confd.DelResponse{}, nil
}
