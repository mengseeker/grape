package agent

import (
	"context"
	"grape/pkg/logger"
	"os"

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
	cmd.PersistentFlags().StringVarP(&config.discoveryAddress, "discoveryAddress", "d", "", "discoveryAddress($DiscoveryAddress)")
	cmd.PersistentFlags().StringVarP(&config.service, "service", "s", "", "service($Namespace/$ServiceCode)")
	cmd.PersistentFlags().StringVarP(&config.run, "run", "r", "", "application run command($Run)")
	return &cmd
}

// if command args is nil, load config from env
func loadEnvConfig() {
	if config.discoveryAddress == "" {
		config.discoveryAddress = os.Getenv("DiscoveryAddress")
	}
	if config.run == "" {
		config.run = os.Getenv("Run")
	}
	if config.service == "" {
		ns := os.Getenv("Namespace")
		serviceCode := os.Getenv("ServiceCode")
		if ns != "" && serviceCode != "" {
			config.service = ns + "/" + serviceCode
		}
	}
}

func start() {
	ctx := context.Background()
	ch := DiscoveryConfig(ctx)
	handleApplication(ctx, ch)
}
