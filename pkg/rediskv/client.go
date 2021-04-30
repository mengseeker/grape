package rediskv

import (
	"github.com/gomodule/redigo/redis"
)

var (
	cli redis.Conn
)

func Connect(address string) error {
	var err error
	cli, err = redis.Dial("tcp", address)
	if err != nil {
		return err
	}
	_, err = cli.Do("PING")
	return err
}

func GetClient() redis.Conn {
	if cli == nil {
		panic("Uninitialized redis conn!")
	}
	return cli
}
