package viewserver

import (
	"context"
	"encoding/json"
	"fmt"
	"grape/api/v1/core"
	"grape/api/v1/view"
	"grape/pkg/etcdcli"
	"grape/pkg/logger"
	"time"
)

const (
	ServerKeyPrefix = "serverlogs/"
)

type apiserver struct {
	cli *etcdcli.Client
	log logger.Logger
	view.UnimplementedApiServerServer
}

func NewApiServer(log logger.Logger, cli *etcdcli.Client) *apiserver {
	return &apiserver{cli: cli, log: log}
}

func (s *apiserver) Set(ctx context.Context, req *view.ServiceLogs) (*core.Empty, error) {
	s.log.Infof("set %s logs", req.Service)
	key := Key(req.Service)
	timeout, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	val, _ := json.Marshal(req)
	_, err := s.cli.Cli.KV.Put(timeout, key, string(val))
	return &core.Empty{}, err
}

func (s *apiserver) Get(ctx context.Context, req *view.GetRequest) (*view.ServiceLogs, error) {
	s.log.Infof("get %s logs", req.Service)
	key := Key(req.Service)
	timeout, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	resp, err := s.cli.Cli.Get(timeout, key)
	if err != nil {
		return nil, err
	}
	logs := view.ServiceLogs{}
	if resp.Count == 1 {
		err := json.Unmarshal(resp.Kvs[0].Value, &logs)
		if err != nil {
			return nil, fmt.Errorf("unable to unmarshal logs: %v", err)
		}
	}
	return &logs, nil
}

func (s *apiserver) Del(ctx context.Context, req *view.DelRequest) (*core.Empty, error) {
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

func Key(service string) string {
	return ServerKeyPrefix + service
}
