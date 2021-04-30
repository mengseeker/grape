package manage

import (
	"github.com/spf13/cobra"
	_ "github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql/driver"
)

func NewManagerCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:     "manage",
		Aliases: []string{"m"},
		Short:   "debug tools",
		Long:    `debug tools`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}
	cmd.AddCommand(NewSessionCmd())
	return &cmd
}
