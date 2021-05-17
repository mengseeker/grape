package resr

import (
	"context"
	"encoding/json"
	"grape/pkg/etcdcli"
	"time"
)

const (
	timeout = 3
)

var Marshal = json.Marshal
var Unmarshal = json.Unmarshal

type Res interface {
	Marshal() []byte
	Key(string) string
}

func Update(cli *etcdcli.Client, r Res) error {
	key := r.Key(cli.ClusterCode)
	val := r.Marshal()
	ctx, calcel := context.WithTimeout(context.Background(), time.Second*timeout)
	defer calcel()
	return cli.CheckAndUpdate(ctx, key, val)
}

func Delete(cli *etcdcli.Client, r Res) error {
	key := r.Key(cli.ClusterCode)
	ctx, calcel := context.WithTimeout(context.Background(), time.Second*timeout)
	defer calcel()
	_, err := cli.Cli.Delete(ctx, key)
	return err
}
