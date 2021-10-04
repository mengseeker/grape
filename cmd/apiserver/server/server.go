package server

import (
	"grape/api/confd"
	"grape/internal/confdserver"
	"grape/internal/share"
	"grape/pkg/etcdcli"
	"grape/pkg/logger"
	"net"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

const ()

var (
	cfgFile string
	log     = logger.NewLogger("apiserver")
)

func NewCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:     "server",
		Aliases: []string{"s"},
		Short:   "apiserver",
		Long:    `apiserver`,
		Run: func(cmd *cobra.Command, args []string) {
			serve()
		},
	}
	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", share.DefaultCfgFile, "config file")
	return &cmd
}

func serve() {
	initConfig()

	ec, err := etcdcli.Connect(viper.GetString("etcd.address"))
	if err != nil {
		log.Fatalf("unalble to connect to etcd: %v", err)
	}
	grpcServer := grpc.NewServer()

	cas := confdserver.NewApiServer(log, ec)
	confd.RegisterApiServerServer(grpcServer, cas)

	apiAddress := viper.GetString("api.address")
	lis, err := net.Listen("tcp", apiAddress)
	if err != nil {
		log.Fatal("unable to listen %s: %v", apiAddress, err)
	}
	log.Infof("listen %s", apiAddress)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("unable to serve grpc server %v", err)
	}
}

func initConfig() {
	viper.SetDefault("api.address", "0.0.0.0:15010")

	share.InitConfig(cfgFile, log)
}
