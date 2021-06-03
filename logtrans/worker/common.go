package worker

import (
	"time"

	"github.com/Shopify/sarama"
)

const (
	logTypeEnvoyAccess = "envoy_access"
	logTypeTrace       = "trace"
	TagKey             = "tag"
	unKnowIndex        = "envoy_unknow"
	esTimeout          = 10 * time.Second
	influxDatabase     = "logs"
	influxMeasurement  = "envoy_access"
)

func FindHeader(hs []*sarama.RecordHeader, key string) string {
	for _, h := range hs {
		if string(h.Key) == key {
			return string(h.Value)
		}
	}
	return ""
}

func GetLogType(m *Message) string {
	return FindHeader(m.Headers, TagKey)
}
