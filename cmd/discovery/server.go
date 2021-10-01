package main

import (
	"grape/pkg/share"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const ()

var (
	cfgFile string
)

func NewServerCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:     "server",
		Aliases: []string{"s"},
		Short:   "discovery server",
		Long:    `discovery server`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return serve()
		},
	}
	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", share.DefaultCfgFile, "config file")
	return &cmd
}

func serve() error {
	initConfig()

	// V3server()
	// dnsServer()
	// logagent server
	// confdagent server
	// gwagent server
	// // xds server
	// xdsAddress := viper.GetString("pilot.address")
	// apiv3.Serve(xdsAddress)
	return nil
}

func initConfig() {
	viper.SetDefault("pilot.address", "0.0.0.0:15010")

	share.InitConfig(cfgFile, log)
}
