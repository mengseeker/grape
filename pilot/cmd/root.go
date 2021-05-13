package cmd

import (
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "pilot",
		Short: "grape xDS service",
		Long:  `.`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}

	cmd.AddCommand(NewServerCmd())
	cmd.AddCommand(NewManageCmd())
	cmd.AddCommand(versionCmd)
	return &cmd
}
