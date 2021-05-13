package forward

import (
	"grape/logtrans/server/logs"
	"log"
)

// 拆分数组
func AddEnvoyTraceLog(entity *logs.Trace) {
	log.Println(entity.Log)
}
