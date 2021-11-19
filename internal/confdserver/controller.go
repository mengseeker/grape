package confdserver

import (
	"context"
	"fmt"
	"grape/api/v1/confd"
	"grape/pkg/etcdcli"
	"grape/pkg/logger"
	"time"
)

const (
	ServiceConfigNotFoundCode = 404
	OkCode                    = 0
)

type apiserver struct {
	cli *etcdcli.Client
	log logger.Logger
	confd.UnimplementedApiServerServer
}

func NewApiServer(log logger.Logger, cli *etcdcli.Client) *apiserver {
	return &apiserver{cli: cli, log: log}
}

func (s *apiserver) Set(ctx context.Context, req *confd.ApiRequest) (*confd.ApiResponse, error) {
	s.log.Infof("set %s", req.ProjectName)
	key := Key(req.ProjectName)

	timeout, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	val := MarshalProject(req.Project)
	_, err := s.cli.Cli.KV.Put(timeout, key, string(val))
	return &confd.ApiResponse{Code: OkCode}, err
}

func (s *apiserver) Get(ctx context.Context, req *confd.ApiRequest) (*confd.ApiResponse, error) {
	s.log.Infof("get %s", req.ProjectName)
	key := Key(req.ProjectName)

	timeout, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	resp, err := s.cli.Cli.Get(timeout, key)
	if err != nil {
		return nil, err
	}
	if resp.Count == 0 {
		return &confd.ApiResponse{Code: ServiceConfigNotFoundCode, Message: "project not found"}, nil
	}

	project, err := UnmarshalProject(resp.Kvs[0].Value)
	if err != nil {
		s.log.Errorf("unable to unmarshal project: %v, value: %s", err, resp.Kvs[0].Value)
		return nil, fmt.Errorf("unable to unmarshal project: %v", err)
	}
	return &confd.ApiResponse{Code: OkCode, Project: project}, nil
}

func (s *apiserver) Del(ctx context.Context, req *confd.ApiRequest) (*confd.ApiResponse, error) {
	s.log.Infof("del %s project", req.ProjectName)
	key := Key(req.ProjectName)
	timeout, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	_, err := s.cli.Cli.Delete(timeout, key)
	if err != nil {
		return nil, err
	}
	return &confd.ApiResponse{Code: OkCode}, nil
}
