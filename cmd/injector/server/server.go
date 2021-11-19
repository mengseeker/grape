package server

import (
	"grape/internal/injector"
	"grape/internal/share"
	"grape/pkg/etcdcli"
	"grape/pkg/logger"
	"net"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const ()

var (
	cfgFile string
	log     = logger.NewLogger("injector")
)

func NewCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:     "server",
		Aliases: []string{"s"},
		Short:   "grape injector",
		Long:    `grape injector`,
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

	runInjectorServer(ec)

}

func runInjectorServer(cli *etcdcli.Client) {
	// set up webhook server
	cert_file := viper.GetString("injector.cert")
	key_file := viper.GetString("injector.cert_key")
	address := viper.GetString("injector.address")
	injectDiscoveryAddress := viper.GetString("injector.inject_discovery_address")
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("unable to listen %s: %v", address, err)
	}
	injectorConfig := &injector.InjectorConfig{
		Cli:                    cli,
		Log:                    log,
		InjectDiscoveryAddress: injectDiscoveryAddress,
		EnableConfd:            viper.GetBool("injector.enable_confd"),
		EnableMesh:             viper.GetBool("injector.enable_mesh"),
		MeshSidecarImage:       viper.GetString("injector.mesh_sidecar_image"),
	}
	log.Infof("listenning injector server at %s", address)
	mux := http.NewServeMux()
	mux.HandleFunc("/inject", injectorConfig.NewjectHandler())
	if err := http.ServeTLS(lis, mux, cert_file, key_file); err != nil {
		log.Fatalf("unable to start inject server: %v", err)
	}
}

func initConfig() {
	viper.SetDefault("injector.address", "0.0.0.0:8443")
	viper.SetDefault("injector.enable_confd", true)
	viper.SetDefault("injector.enable_mesh", true)
	viper.SetDefault("injector.inject_discovery_address", "grape-discovery.grape-system:15020")

	share.InitConfig(cfgFile, log)
}
