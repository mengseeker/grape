package confdserver

import (
	"context"
	"encoding/json"
	"fmt"
	"grape/api/confd"
	"grape/api/core"
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

func (s *apiserver) Set(ctx context.Context, req *confd.Configs) (*core.Empty, error) {
	s.log.Infof("set %s configs", req.Service)
	key := Key(req.Service)
	timeout, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	val, _ := json.Marshal(req)
	_, err := s.cli.Cli.KV.Put(timeout, key, string(val))
	return &core.Empty{}, err
}

func (s *apiserver) Get(ctx context.Context, req *confd.GetRequest) (*confd.Configs, error) {
	s.log.Infof("get %s configs", req.Service)
	key := Key(req.Service)
	timeout, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	resp, err := s.cli.Cli.Get(timeout, key)
	if err != nil {
		return nil, err
	}
	cfs := confd.Configs{}
	if resp.Count == 1 {
		err := json.Unmarshal(resp.Kvs[0].Value, &cfs)
		if err != nil {
			return nil, fmt.Errorf("unable to unmarshal configs: %v", err)
		}
	}
	return &cfs, nil
}

func (s *apiserver) Del(ctx context.Context, req *confd.DelRequest) (*core.Empty, error) {
	s.log.Infof("del %s configs", req.Service)
	key := Key(req.Service)
	timeout, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	_, err := s.cli.Cli.Delete(timeout, key)
	if err != nil {
		return nil, err
	}
	return &core.Empty{}, nil
}
