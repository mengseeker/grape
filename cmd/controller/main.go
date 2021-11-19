package main

import (
	"grape/cmd/controller/server"
	"grape/pkg/logger"

	"github.com/spf13/cobra"
)

var (
	log     = logger.NewLogger("controller")
	Version = "0.0.0"
)

func main() {
	root := cobra.Command{
		Use:     "controller",
		Short:   "controller",
		Version: Version,
		Long:    `.`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}

	root.AddCommand(server.NewCmd())
	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
}
