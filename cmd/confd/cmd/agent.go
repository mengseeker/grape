package cmd

import (
	"context"
	confdv1 "grape/api/v1/confd"
	"grape/pkg/logger"

	"github.com/spf13/cobra"
)

var config struct {
	discoveryAddress string
	discovery        bool
	project          string
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
	cmd.PersistentFlags().StringVarP(&config.project, "project", "p", "", "project")
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
	discoveryChan := make(chan *confdv1.Configs)
	go runDiscovery(discoveryChan)

	app := NewApplication(config.runCmd)
	cfs := <-discoveryChan
	if config.discovery {
		// try to start application
		if err := app.TryStart(cfs.RunCmd); err != nil {
			log.Fatalf("start application err: %v", err)
		}
		handleUpdateConfigs(context.Background(), discoveryChan, app)
	} else {
		if err := app.RunExecApplication(cfs.RunCmd); err != nil {
			log.Fatal(err)
		}
	}
}

func checkfigure() {
	if config.project == "" {
		log.Fatalf("project must be set")
	}
	if config.discoveryAddress == "" {
		log.Fatal("discoveryAddress must be set")
	}
}
