package confdserver

import (
	"context"
	"encoding/json"
	"fmt"
	"grape/api/v1/confd"
	"grape/internal/share"
	"grape/pkg/etcdcli"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

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

func Key(namespace, service string) string {
	return share.ServerKeyPrefix + namespace + "/" + service
}

func GetServiceConfigs(cli *etcdcli.Client, namespace, service, group string, loadVersion int64) (*confd.Configs, int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	key := Key(namespace, service)
	ops := []clientv3.OpOption{}
	if loadVersion != 0 {
		ops = append(ops, clientv3.WithRev(loadVersion))
	}
	resp, err := cli.Cli.Get(ctx, key, ops...)
	if err != nil {
		return nil, 0, fmt.Errorf("unalble to get confis form etcd: %v", err)
	}
	// if config not found, return an emptys
	if resp.Count == 0 {
		return &confd.Configs{}, 0, nil
	}
	sf, err := UnmarshalServiceConfig(resp.Kvs[0].Value)
	if err != nil {
		return nil, 0, fmt.Errorf("unmarshal configs err: %v", err)
	}
	return GetGroupConfig(sf, group), resp.Header.Revision, nil
}

func GetRevServiceConfigs(cli *etcdcli.Client, namespace, service, group string, loadVersion int64) (*confd.Configs, error) {
	cf, _, err := GetServiceConfigs(cli, namespace, service, group, loadVersion)
	return cf, err
}

func GetLatestServiceConfigs(cli *etcdcli.Client, namespace, service, group string) (*confd.Configs, error) {
	cf, _, err := GetServiceConfigs(cli, namespace, service, group, 0)
	return cf, err
}
