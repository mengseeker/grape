package agent

import (
	"context"
	"grape/api/v1/confd"
	"grape/internal/confdserver"
	"grape/internal/share"
	"grape/pkg/logger"
	"time"

	"github.com/spf13/cobra"
)

var config struct {
	discoveryAddress string
	discovery        bool
	service          string
	group            string
	loadVersion      int64
}

var (
	log = logger.NewLogger("confd")
)

func NewAgentCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "confd",
		Short: "confd",
		Long:  `.`,
		Run: func(cmd *cobra.Command, args []string) {
			loadEnvConfig()
			checkfigure()
			if config.discovery {
				runDiscovery()
			} else {
				runLocadConfigs()
			}
		},
	}
	cmd.PersistentFlags().StringVarP(&config.discoveryAddress, "discoveryAddress", "a", "discovery:15020", "discoveryAddress")
	cmd.PersistentFlags().StringVarP(&config.service, "service", "s", "", "service")
	cmd.PersistentFlags().StringVarP(&config.group, "group", "g", "", "group")
	cmd.PersistentFlags().Int64VarP(&config.loadVersion, "loadVersion", "l", 0, "loadVersion")
	cmd.PersistentFlags().BoolVarP(&config.discovery, "discovery", "d", false, "enable discovery always")
	return &cmd
}

// if command args is nil, load config from env
func loadEnvConfig() {
	if config.service == "" {
		config.service = share.GetService()
	}
	if config.group == "" {
		config.group = share.GetGroupCode()
	}
}

func checkfigure() {
	if config.service == "" {
		log.Fatalf("service not configued")
	}
	if config.discoveryAddress == "" {
		log.Fatal("discoveryAddress must be set")
	}
}

func runLocadConfigs() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	DialDiscoveryServer(ctx)
	req := confd.DownloadRequest{
		Service: config.service,
		Group: config.group,
		LoadVersion: config.loadVersion,
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
}

func runDiscovery() {
	ctx := context.Background()
	DialDiscoveryServer(ctx)
	cfChan := make(chan *confd.Configs)
	go handleDiscovery(ctx, cfChan)
	handleUpdateConfig(ctx, cfChan)
}
