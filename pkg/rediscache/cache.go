package rediscache

import (
	"encoding/json"
	"grape/pkg/redispool"
)

const (
	TTL = 60 * 5
)

func AutoGet(key string, v func() interface{}) interface{} {
	rep, err := redispool.GetClient().Do("GET", key)
	if err != nil {
		panic(err)
	}
	if rep == nil {
		val := v()
		raw, err := json.Marshal(val)
		if err != nil {
			panic(err)
		}
		_, err = redispool.GetClient().Do("SETEX", key, TTL, raw)
		if err != nil {
			panic(err)
		}
		return val
	}
	var r interface{}
	err = json.Unmarshal([]byte(rep.(string)), &r)
	if err != nil {
		panic(err)
	}
	return r
}

func UnSet(key string) {
	_, err := redispool.GetClient().Do("DEL", key)
	if err != nil {
		panic(err)
	}
}
