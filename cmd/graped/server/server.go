package server

import (
	"grape/api/v1/confd"
	"grape/api/v1/view"
	"grape/internal/confdserver"
	"grape/internal/injector"
	"grape/internal/share"
	"grape/internal/viewserver"
	"grape/pkg/etcdcli"
	"grape/pkg/logger"
	"net"
	"net/http"

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

	ec, err := etcdcli.Connect(viper.GetString("etcd.address"))
	if err != nil {
		log.Fatalf("unalble to connect to etcd: %v", err)
	}

	if viper.GetBool("injector.enable") {
		go runInjectorServer(ec)
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

	apiAddress := viper.GetString("graped.address")
	lis, err := net.Listen("tcp", apiAddress)
	if err != nil {
		log.Fatal("unable to listen %s: %v", apiAddress, err)
	}
	log.Infof("listenning api server at %s", apiAddress)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("unable to serve grpc server %v", err)
	}
}

func runInjectorServer(cli *etcdcli.Client) {
	// set up webhook server
	cert_file := viper.GetString("injector.cert")
	key_file := viper.GetString("injector.cert_key")
	address := viper.GetString("injector.address")
	discoveryAddress := viper.GetString("injector.discovery_address")
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("unable to listen %s: %v", address, err)
	}
	injectorConfig := &injector.InjectorConfig{
		Cli:              cli,
		Log:              log,
		DiscoveryAddress: discoveryAddress,
		EnableConfd:      viper.GetBool("injector.enable_confd"),
		EnableMesh:       viper.GetBool("injector.enable_mesh"),
		EnableView:       viper.GetBool("injector.enable_view"),
	}
	log.Infof("listenning injector server at %s", address)
	mux := http.NewServeMux()
	mux.HandleFunc("/inject", injectorConfig.NewjectHandler())
	if err := http.ServeTLS(lis, mux, cert_file, key_file); err != nil {
		log.Fatalf("unable to start inject server: %v", err)
	}
}

func initConfig() {
	viper.SetDefault("graped.address", "0.0.0.0:15010")

	viper.SetDefault("injector.enable", true)
	viper.SetDefault("injector.address", "0.0.0.0:8082")
	viper.SetDefault("injector.enable_confd", true)
	viper.SetDefault("injector.enable_mesh", true)
	viper.SetDefault("injector.enable_view", true)
	viper.SetDefault("injector.discovery_address", "grape-discovery:15020")

	share.InitConfig(cfgFile, log)
}
