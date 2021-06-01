package cmd

import (
	"grape/pkg/logger"

	"github.com/spf13/cobra"
)

var log = logger.NewLogger("logtrans")

func NewRootCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "logtrans",
		Short: ".",
		Long:  `.`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}
	cmd.AddCommand(NewConsumerCmd())
	// cmd.AddCommand(NewManagerCmd())
	// cmd.AddCommand(NewMigrateCmd())
	return &cmd
}
