package redispool

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

var (
	pool *redis.Pool
)

func Connect(address string) error {
	pool = &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", address) },
	}
	_, err := pool.Get().Do("PING")
	return err
}

func GetClient() redis.Conn {
	return pool.Get()
}
