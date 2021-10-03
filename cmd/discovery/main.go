package main

import (
	"grape/cmd/discovery/manage"
	"grape/cmd/discovery/server"
	"grape/pkg/logger"

	"github.com/spf13/cobra"
)

var (
	log     = logger.NewLogger("discovery")
	Version = "0.0.0"
)

func main() {
	root := cobra.Command{
		Use:     "grape-discovery",
		Short:   "grape discovery service",
		Version: Version,
		Long:    `.`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}

	root.AddCommand(server.NewCmd())
	root.AddCommand(manage.NewCmd())
	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
}
