package cmd

import (
	"grape/grapeapi/models"
	"grape/pkg/share"

	"github.com/spf13/cobra"
)

func NewMigrateCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:     "migrate",
		Aliases: []string{"s"},
		Short:   "db migrates",
		Long:    `.`,
		Run: func(cmd *cobra.Command, args []string) {
			initServer()
			autoMigrate()
		},
	}
	cmd.Flags().StringVarP(&configFile, "config", "c", share.DefaultCfgFile, "config file")

	return &cmd
}

func autoMigrate() {
	err := models.GetDB().AutoMigrate(
		&models.Namespace{},
		&models.Cluster{},
		&models.Service{},
		&models.Group{},
		&models.Node{},
		&models.Policy{},
		&models.User{},
	)
	if err != nil {
		log.Fatalf("automigrate err: %v", err)
	}
}
