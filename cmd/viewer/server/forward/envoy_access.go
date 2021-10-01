package forward

import (
	"grape/logtrans/server/logs"
	"log"
)

func AddEnvoyAccessLog(entity *logs.EnvoyAccessLog) {
	log.Println(string(entity.Marshaler()))
}
