package worker

import (
	"time"
)

const (
	unKnowIndex        = "envoy_unknow"
	esTimeout          = 10 * time.Second
	influxDatabase     = "logs"
	influxMeasurement  = "envoy_access"
)
