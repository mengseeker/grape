package share

import (
	"grape/pkg/logger"
	"strings"

	"github.com/spf13/viper"
)

const DefaultCfgFile = "grape.yaml"

func InitConfig(cfgFile string, log logger.Logger) {
	viper.SetConfigFile(cfgFile)
	viper.SetEnvPrefix(ViperEnvPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		log.Infof("Using config file: %s", viper.ConfigFileUsed())
	} else {
		log.Fatalf("unable to load config: %v", err)
	}
}
