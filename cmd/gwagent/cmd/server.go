package cmd

import (
	"grape/extauth/auth"
	"grape/pkg/share"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configFile string
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
	cmd.Flags().StringVarP(&configFile, "config", "c", share.DefaultCfgFile, "config file")
	return &cmd
}

func Serve() {
	initConfig(configFile)
	err := auth.ConnectEtcd(viper.GetString("etcd.address"))
	if err != nil {
		log.Fatalf("connect to ectd err: %v", err)
	}
	log.Info("etcd cluster connected")
	// TODO etcd
	clusterCode := viper.GetString("cluster_code")
	if clusterCode == "" {
		log.Fatal("clusterCode must setting")
	}
	auth.Serve(viper.GetString("auth.address"), clusterCode)
}

func initConfig(cfg string) {
	viper.SetDefault("auth.address", ":11001")
	viper.SetDefault("etcd.address", ":6379")

	share.InitConfig(configFile, log)
}
