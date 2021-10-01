package main

import "github.com/spf13/cobra"

func NewManageCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:     "manage",
		Aliases: []string{"m"},
		Short:   "pilot manager tools",
		Long:    ".",
		Run: func(cmd *cobra.Command, args []string) {
			println("todo")
		},
	}

	return &cmd
}
