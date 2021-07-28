package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func IsConfigSourceKafka() bool {
	return len(viper.GetStringMap("logtrans.source.kafka")) > 0
}

func CheckSourceKafkaConfig() error {
	address := viper.GetString("logtrans.source.kafka.address")
	if len(address) == 0 {
		return fmt.Errorf("the config %q must be set", "logtrans.source.kafka.address")
	}
	if len(viper.GetString("logtrans.source.kafka.topics")) == 0 {
		return fmt.Errorf("the config %q must be set", "logtrans.source.kafka.topics")
	}
	return nil
}

func GetSourceKafkaConfig() (address, topics, group, version, assignor string) {
	address = viper.GetString("logtrans.source.kafka.address")
	group = viper.GetString("logtrans.source.kafka.group")
	topics = viper.GetString("logtrans.source.kafka.topics")
	version = viper.GetString("logtrans.source.kakfa.version")
	assignor = viper.GetString("logtrans.source.kafka.assignor")
	if group == "" {
		group = "logtrans"
	}
	if version == "" {
		version = "2.1.1"
	}
	if assignor == "" {
		assignor = "range"
	}
	return
}
