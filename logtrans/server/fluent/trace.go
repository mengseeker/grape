package fluent

import (
	"log"
	"grape/logtrans/server/logs"
)

// 拆分数组
func AddEnvoyTraceLog(entity *logs.Trace) {
	log.Println(entity.Log)
}
