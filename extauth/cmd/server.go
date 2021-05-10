package cmd

import (
	"grape/extauth/auth"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configFile        string
	defaultConfigFile = "grape.yaml"
	envPrefix         = "GRAPE"
)

func NewServerCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:     "server",
		Aliases: []string{"s"},
		Short:   "start extauth server",
		Long:    `.`,
		Run: func(cmd *cobra.Command, args []string) {
			Serve()
		},
	}
	cmd.Flags().StringVarP(&configFile, "config", "c", defaultConfigFile, "config file (default: "+defaultConfigFile+")")
	return &cmd
}

func Serve() {
	initConfig(configFile)
	// TODO etcd
	auth.Serve(viper.GetString("auth.address"))
}

func initConfig(cfg string) {
	viper.SetDefault("auth.address", ":11001")
	viper.SetDefault("etcd.address", ":6379")

	viper.SetConfigFile(cfg)
	viper.SetEnvPrefix(envPrefix)
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		log.Infof("Using config file: %s", viper.ConfigFileUsed())
	} else {
		log.Fatalf("unable to load config: %v", err)
	}
}
