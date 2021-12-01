package server

import (
	"grape/api/v1/confd"
	"grape/internal/confdserver"
	"grape/internal/share"
	"grape/pkg/etcdcli"
	"grape/pkg/logger"
	"net"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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
		Short:   "grape controller",
		Long:    `grape controller`,
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

	gops := []grpc.ServerOption{}
	if viper.GetString("controller.cert") != "" {
		creds, err := credentials.NewServerTLSFromFile(viper.GetString("controller.cert"), viper.GetString("controller.cert_key"))
		if err != nil {
			log.Fatal(err)
		}
		gops = append(gops, grpc.Creds(creds))
	}

	grpcServer := grpc.NewServer(gops...)

	// config apiserver
	cas := confdserver.NewApiServer(log, ec)
	confd.RegisterApiServerServer(grpcServer, cas)

	// gw apiserver
	// mesh apiserver
	// k8s apiserver

	apiAddress := viper.GetString("controller.address")
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
	viper.SetDefault("controller.address", "0.0.0.0:15010")

	share.InitConfig(cfgFile, log)
}
