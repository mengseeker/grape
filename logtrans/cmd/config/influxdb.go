package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func IsConfigOutInfluxDB() bool {
	return len(viper.GetStringMap("logtrans.out.influxdb")) > 0
}

func CheckOutInfluxDBConfig() error {
	address := viper.GetString("logtrans.out.influxdb.address")
	if len(address) == 0 {
		return fmt.Errorf("the config %q must be set", "logtrans.out.influxdb.address")
	}
	if len(viper.GetStringSlice("logtrans.out.influxdb.types")) == 0 {
		return fmt.Errorf("the config %q must be set", "logtrans.out.influxdb.types")
	}
	return nil
}

func GetOutInfluxDBConfig() (address string, batch_size, interval int, types []string) {
	address = viper.GetString("logtrans.out.influxdb.address")
	batch_size = viper.GetInt("logtrans.out.influxdb.batch_size")
	interval = viper.GetInt("logtrans.out.influxdb.interval")
	types = viper.GetStringSlice("logtrans.out.influxdb.types")

	// default values
	if batch_size == 0 {
		batch_size = 100
	}
	if interval == 0 {
		interval = 3
	}

	return
}
