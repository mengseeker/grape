package etlinks

import (
	"fmt"
	"grape/grape/models"
	"grape/pkg/etcdcli"
)

var (
	links = map[int64]*etcdcli.Client{}
)

func AddLink(link *models.EtcdLink) error {
	cli, err := etcdcli.Connect(link.Address)
	if err != nil {
		return err
	}
	links[link.ID] = cli
	return nil
}

func GetCli(clu *models.Cluster) *etcdcli.Client {
	cli, exists := links[int64(clu.EtcdID)]
	if exists {
		return cli
	}
	err := AddLink(clu.EctdLink())
	if err != nil {
		panic(fmt.Errorf("can not get link :%v", err))
	}
	return links[int64(clu.EtcdID)]
}
