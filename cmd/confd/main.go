package main

import (
	"grape/cmd/confd/agent"
	"grape/pkg/logger"

	"github.com/spf13/cobra"
)

var (
	Version = "0.0.0"
	log     = logger.NewLogger("confd")
)

func main() {
	root := cobra.Command{
		Use:     "confd",
		Short:   "confd",
		Version: Version,
		Long:    `.`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}

	root.AddCommand(agent.NewAgentCmd())
	// root.AddCommand(NewManageCmd())
	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
}
