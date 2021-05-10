package fluent

import (
	"log"
	"grape/logtrans/server/logs"
)

func AddEnvoyAccessLog(entity *logs.EnvoyAccessLog) {
	log.Println(string(entity.Marshaler()))
}
