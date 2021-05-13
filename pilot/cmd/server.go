package cmd

import (
	"grape/pilot/apiv3"
	"grape/pkg/share"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const ()

var (
	defaultCfgFile = "pilot.yaml"
	cfgFile        string
)

func NewServerCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:     "server",
		Aliases: []string{"s"},
		Short:   "start pilot",
		Long:    `start pilot`,
		Run: func(cmd *cobra.Command, args []string) {
			Serve()
			<-make(chan int)
		},
	}
	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", defaultCfgFile, "config file")
	return &cmd
}

func Serve() {
	initConfig()

	// xds server
	xdsAddress := viper.GetString("pilot.address")
	apiv3.Serve(xdsAddress, xdsAddress)
}

func initConfig() {
	// xds
	viper.SetDefault("pilot.address", "0.0.0.0:15010")

	viper.SetConfigFile(cfgFile)
	viper.SetEnvPrefix(share.ViperEnvPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		log.Infof("Using config file: %s", viper.ConfigFileUsed())
	} else {
		log.Fatalf("unable to load config: %v", err)
	}
}
