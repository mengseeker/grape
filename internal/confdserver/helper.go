package confdserver

import (
	"context"
	"encoding/json"
	"fmt"
	confdv1 "grape/api/v1/confd"
	"grape/internal/share"
	"grape/pkg/etcdcli"
	"time"
)

func UnmarshalProject(raw []byte) (*confdv1.Project, error) {
	project := confdv1.Project{}
	err := json.Unmarshal(raw, &project)
	return &project, err
}

func MarshalProject(project *confdv1.Project) []byte {
	j, _ := json.Marshal(project)
	return j
}

func GetGroupConfig(project *confdv1.Project, group string) *confdv1.Configs {
	if cf, ok := project.GroupConfigs[group]; ok {
		return cf
	}
	return project.GroupConfigs[share.ConfdDefaultGroupName]
}

func Key(projectName string) string {
	return share.ServerKeyPrefix + projectName
}

func GetProjectConfigs(cli *etcdcli.Client, projectName, group string) (*confdv1.Configs, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	key := Key(projectName)
	resp, err := cli.Cli.Get(ctx, key)
	if err != nil {
		return nil, fmt.Errorf("unalble to get project form etcd: %v", err)
	}
	if resp.Count == 0 {
		return nil, fmt.Errorf("project %q nou found", projectName)
	}
	f, err := UnmarshalProject(resp.Kvs[0].Value)
	if err != nil {
		return nil, fmt.Errorf("unmarshal project err: %v", err)
	}
	return GetGroupConfig(f, group), nil
}
