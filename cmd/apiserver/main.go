package main

import (
	"grape/cmd/apiserver/server"
	"grape/pkg/logger"

	"github.com/spf13/cobra"
)

var (
	log     = logger.NewLogger("apiserver")
	Version = "0.0.0"
)

func main() {
	root := cobra.Command{
		Use:     "grape-apiserver",
		Short:   "grape apiserver",
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
