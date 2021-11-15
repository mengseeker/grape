package cmd

import (
	"context"
	confdv1 "grape/api/v1/confd"
	"grape/internal/confdserver"
	"grape/pkg/logger"
	"time"

	"github.com/spf13/cobra"
)

var config struct {
	discoveryAddress string
	discovery        bool
	namespace        string
	service          string
	group            string
	// loadVersion      int64
	runCmd string
}

var (
	log = logger.NewLogger("confd")
)

func NewCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "confd",
		Short: "confd",
		Long:  `.`,
		Run: func(cmd *cobra.Command, args []string) {
			// loadEnvConfig()
			checkfigure()
			startAgent()
		},
	}
	cmd.PersistentFlags().StringVarP(&config.discoveryAddress, "discoveryAddress", "a", "", "discoveryAddress")
	cmd.PersistentFlags().BoolVarP(&config.discovery, "discovery", "d", false, "discovery")
	cmd.PersistentFlags().StringVarP(&config.namespace, "namespace", "n", "", "namespace")
	cmd.PersistentFlags().StringVarP(&config.service, "service", "s", "", "service")
	cmd.PersistentFlags().StringVarP(&config.group, "group", "g", "", "group")
	cmd.PersistentFlags().StringVarP(&config.runCmd, "runCmd", "r", "", "runCmd")
	return &cmd
}

// if command args is nil, load config from env
// func loadEnvConfig() {
// 	if config.namespace == "" {
// 		config.namespace = share.GetNamespace()
// 	}
// 	if config.service == "" {
// 		config.service = share.GetServiceCode()
// 	}
// 	if config.group == "" {
// 		config.group = share.GetGroupCode()
// 	}
// }

func startAgent() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	cfs := lcadConfigs(ctx)
	app := NewApplication(config.runCmd)
	if config.discovery {
		discoveryChan := make(chan *confdv1.Configs, 1)
		discoveryChan <- cfs
		go runDiscovery(discoveryChan)
		handleUpdateConfigs(context.Background(), discoveryChan, app)
	} else {
		app.UpdateEnv(cfs.EnvConfigs)
		if err := app.RunExecApplication(cfs.RunCmd); err != nil {
			log.Fatal(err)
		}
	}
}

func checkfigure() {
	if config.namespace == "" {
		log.Fatalf("namespace must be set")
	}
	if config.service == "" {
		log.Fatalf("service must be set")
	}
	if config.discoveryAddress == "" {
		log.Fatal("discoveryAddress must be set")
	}
}

func lcadConfigs(ctx context.Context) *confdv1.Configs {
	dialDiscoveryServer(ctx)
	req := confdv1.DownloadRequest{
		Service: config.service,
		Group:   config.group,
		// LoadVersion: config.loadVersion,
	}
	resp, err := disconveryClient.Download(ctx, &req)
	if err != nil {
		log.Fatalf("failed to download configs: %v", err)
	}
	if resp.Code != confdserver.OkCode {
		log.Fatalf("failed to download configs: %v", resp.Message)
	}
	err = WriteConfigFiles(resp.Configs)
	if err != nil {
		log.Fatalf("failed to download configs: %v", err)
	}
	return resp.Configs
}
