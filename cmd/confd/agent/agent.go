package agent

import (
	"context"
	"grape/internal/share"
	"grape/pkg/logger"

	"github.com/spf13/cobra"
)

var config struct {
	discoveryAddress string
	service          string
	run              string
}

var (
	log = logger.NewLogger("confd")
)

func NewAgentCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "agent",
		Short: "agent",
		Long:  `.`,
		Run: func(cmd *cobra.Command, args []string) {
			loadEnvConfig()
			start()
		},
	}
	cmd.PersistentFlags().StringVarP(&config.discoveryAddress, "discoveryAddress", "d", "", "discoveryAddress")
	cmd.PersistentFlags().StringVarP(&config.service, "service", "s", "", "service")
	cmd.PersistentFlags().StringVarP(&config.run, "run", "r", "", "application run command")
	return &cmd
}

// if command args is nil, load config from env
func loadEnvConfig() {
	if config.discoveryAddress == "" {
		config.discoveryAddress = share.GetDiscoveryAddress()
	}
	if config.run == "" {
		config.run = share.GetRun()
	}
	if config.service == "" {
		config.service = share.GetService()
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

func start() {
	checkfigure()
	ctx := context.Background()
	ch := DiscoveryConfig(ctx)
	handleApplication(ctx, ch)
}
