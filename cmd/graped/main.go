package main

import (
	"grape/cmd/graped/server"
	"grape/pkg/logger"
	"time"

	"github.com/spf13/cobra"
)

var (
	log     = logger.NewLogger("apiserver")
	Version = "0.0.0"
)

func main() {
	root := cobra.Command{
		Use:     "graped",
		Short:   "graped",
		Version: Version,
		Long:    `.`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}

	root.AddCommand(server.NewCmd())
	if err := root.Execute(); err != nil {
		log.Fatal(err)
		time.Now().Format()
	}
}
