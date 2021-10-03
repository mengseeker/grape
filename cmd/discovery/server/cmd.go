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
	log     = logger.NewLogger("discovery")
)

func NewCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:     "server",
		Aliases: []string{"s"},
		Short:   "discovery server",
		Long:    `discovery server`,
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

	// V3server()
	// dnsServer()
	// logagent server
	// confdagent server
	cs := confdserver.NewServer(log, ec)
	confd.RegisterConfdServerServer(grpcServer, cs)

	// gwagent server
	// // xds server
	// xdsAddress := viper.GetString("pilot.address")
	// apiv3.Serve(xdsAddress)
	discoveryAddress := viper.GetString("discovery.address")
	lis, err := net.Listen("tcp", discoveryAddress)
	if err != nil {
		log.Fatal("unable to listen %s: %v", discoveryAddress, err)
	}
	log.Infof("listen %s", discoveryAddress)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("unable to serve grpc server %v", err)
	}
}

func initConfig() {
	viper.SetDefault("discovery.address", "0.0.0.0:15010")

	share.InitConfig(cfgFile, log)
}
