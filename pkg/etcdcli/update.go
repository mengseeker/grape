package etcdcli

import (
	"context"
	"reflect"
)

func (cli *Client) CheckAndUpdate(ctx context.Context, k string, val []byte) error {
	old, err := cli.Cli.Get(ctx, k)
	if err != nil {
		return err
	}
	if old.Count == 1 {
		if reflect.DeepEqual(old.Kvs[0], val) {
			return nil
		}
	}
	_, err = cli.Cli.Put(ctx, k, string(val))
	return err
}
