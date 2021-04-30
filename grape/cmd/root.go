package cmd

import (
	"grape/grape/cmd/server"

	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "grape",
		Short: "grape. a service mesh framework.",
		Long:  `.`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}
	cmd.AddCommand(server.NewServerCmd())
	return &cmd
}
