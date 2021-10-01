package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func IsConfigOutEs() bool {
	return len(viper.GetStringMap("logtrans.out.es")) > 0
}

func CheckOutEsConfig() error {
	address := viper.GetString("logtrans.out.es.address")
	if len(address) == 0 {
		return fmt.Errorf("the config %q must be set", "logtrans.out.es.address")
	}
	if len(viper.GetStringSlice("logtrans.out.es.types")) == 0 {
		return fmt.Errorf("the config %q must be set", "logtrans.out.es.types")
	}
	return nil
}

func GetOutEsConfig() (address string, batch_size, interval int, types []string) {
	address = viper.GetString("logtrans.out.es.address")
	batch_size = viper.GetInt("logtrans.out.es.batch_size")
	interval = viper.GetInt("logtrans.out.es.interval")
	types = viper.GetStringSlice("logtrans.out.es.types")

	// default values
	if batch_size == 0 {
		batch_size = 100
	}
	if interval == 0 {
		interval = 3
	}

	return
}
