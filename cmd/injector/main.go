package main

import (
	"grape/cmd/injector/server"
	"grape/pkg/logger"

	"github.com/spf13/cobra"
)

var (
	log     = logger.NewLogger("injector")
	Version = "0.0.0"
)

func main() {
	root := cobra.Command{
		Use:     "injector",
		Short:   "injector",
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
