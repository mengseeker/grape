package etcdcli

import (
	"context"
	"crypto/md5"
	"encoding/hex"
)

var (
	versions = map[string]string{}
)

func (cli *Client) CheckAndUpdate(ctx context.Context, k string, val []byte) error {
	newVersion := HashVersion(val)
	publishedVersion := versions[k]
	if newVersion == publishedVersion {
		return nil
	}
	versions[k] = newVersion
	_, err := cli.Cli.Put(ctx, k, string(val))
	return err
}

func HashVersion(val []byte) string {
	dst := md5.Sum(val)
	return hex.EncodeToString(dst[:])
}
