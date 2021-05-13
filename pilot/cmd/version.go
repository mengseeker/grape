package cmd

import "github.com/spf13/cobra"

const (
	// PILOTVERSION pilot version
	PILOTVERSION = "0.0.1"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "grape version",
	Run: func(cmd *cobra.Command, args []string) {
		println(PILOTVERSION)
	},
}
