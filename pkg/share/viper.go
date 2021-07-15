package share

import (
	"grape/pkg/logger"
	"strings"

	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
)

const DefaultCfgFile = "bootstrap.yaml"

func InitConfig(cfgFile string, log logger.Logger) {
	viper.SetConfigFile(cfgFile)
	viper.SetEnvPrefix(ViperEnvPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// viper.SetDefault("logging_level", "info")

	if err := viper.ReadInConfig(); err == nil {
		log.Infof("Using config file: %s", viper.ConfigFileUsed())
	} else {
		log.Fatalf("unable to load config: %v", err)
	}

	loggingLevel := viper.GetString("logging_level")
	if loggingLevel != "" {
		var l zapcore.Level
		err := l.Set(loggingLevel)
		if err != nil {
			log.Fatalf("unable to set logging level: %v", err)
		}
		logger.SetLevel(l)
	}

}
