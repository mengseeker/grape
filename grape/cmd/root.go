package cmd

import (

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
	cmd.AddCommand(NewServerCmd())
	cmd.AddCommand(NewManagerCmd())
	return &cmd
}
