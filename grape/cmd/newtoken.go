package cmd

import (
	"fmt"
	"grape/pkg/session"

	"github.com/spf13/cobra"
)

func NewSessionCmd() *cobra.Command {
	var configFile string
	var defaultConfigFile = "grape.yaml"
	var userID int
	cmd := cobra.Command{
		Use:   "newtoken",
		Short: "create a login token",
		Long:  `create a login token`,
		Run: func(cmd *cobra.Command, args []string) {
			InitConfig(configFile)
			InitRedis()
			s := session.NewSession(userID)
			s.Set("debug", "1")
			fmt.Println(s.ID)
		},
	}
	cmd.Flags().StringVarP(&configFile, "config", "c", defaultConfigFile, "config file (default: "+defaultConfigFile+")")
	cmd.Flags().IntVarP(&userID, "userid", "u", 0, "userID")
	return &cmd
}
