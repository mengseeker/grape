package redislocker

import (
	"fmt"
	"grape/pkg/redispool"
)

const (
	defaultTTL = 60
)

func LockP(key string, ttl int, f func()) {
	if ttl <= 0 {
		ttl = defaultTTL
	}
	lock, err := redispool.GetClient().Do("SET", key, 1, "NX", "PX", ttl*1000)
	if err != nil {
		panic(err)
	}
	if lock == nil {
		panic(fmt.Errorf("unable to get resource %s: resource has locked", key))
	}
	defer func() {
		redispool.GetClient().Do("DEL", key)
	}()

	f()
}

// TODO
func LockE(key string, ttl int, f func()) error {

	return nil
}
