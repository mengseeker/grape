package server

import (
	"grape/api/v1/confd"
	"grape/api/v1/view"
	"grape/internal/confdserver"
	"grape/internal/share"
	"grape/internal/viewserver"
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
	log     = logger.NewLogger("graped")
)

func NewCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:     "server",
		Aliases: []string{"s"},
		Short:   "grape apiserver",
		Long:    `grape apiserver`,
		Run: func(cmd *cobra.Command, args []string) {
			serve()
		},
	}
	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", share.DefaultCfgFile, "config file")
	return &cmd
}

func serve() {
	initConfig()

	ec, err := etcdcli.Connect(
		viper.GetString("etcd.address"),
		viper.GetString("etcd.username"),
		viper.GetString("etcd.password"),
	)
	if err != nil {
		log.Fatalf("unalble to connect to etcd: %v", err)
	}

	grpcServer := grpc.NewServer()

	// config apiserver
	cas := confdserver.NewApiServer(log, ec)
	confd.RegisterApiServerServer(grpcServer, cas)

	// log apiserver
	las := viewserver.NewApiServer(log, ec)
	view.RegisterApiServerServer(grpcServer, las)

	// gw apiserver
	// mesh apiserver
	// k8s apiserver

	apiAddress := viper.GetString("apiserver.address")
	lis, err := net.Listen("tcp", apiAddress)
	if err != nil {
		log.Fatal("unable to listen %s: %v", apiAddress, err)
	}
	log.Infof("listenning api server at %s", apiAddress)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("unable to serve grpc server %v", err)
	}
}

func initConfig() {
	viper.SetDefault("apiserver.address", "0.0.0.0:15010")

	share.InitConfig(cfgFile, log)
}
