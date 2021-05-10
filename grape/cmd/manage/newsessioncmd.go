package manage

import (
	"fmt"
	"grape/grape/cmd/server"
	"grape/pkg/session"

	"github.com/spf13/cobra"
	_ "github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql/driver"
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
			server.InitConfig(configFile)
			server.InitRedis()
			s := session.NewSession(userID)
			s.Set("debug", "1")
			fmt.Println(s.ID)
		},
	}
	cmd.Flags().StringVarP(&configFile, "config", "c", defaultConfigFile, "config file (default: "+defaultConfigFile+")")
	cmd.Flags().IntVarP(&userID, "userid", "u", 0, "userID")
	return &cmd
}
