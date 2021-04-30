package server

import (
	"database/sql"
	"grape/grape/pkg/rediskv"
	"grape/grape/server"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/volatiletech/sqlboiler/boil"
	_ "github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql/driver"
)

var (
	configFile        string
	defaultConfigFile = "grape.yaml"
	envPrefix         = "GRAPE"
)

func NewServerCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:     "server",
		Aliases: []string{"s"},
		Short:   "start grape",
		Long:    `start grape api server`,
		Run: func(cmd *cobra.Command, args []string) {
			Serve()
		},
	}
	cmd.Flags().StringVarP(&configFile, "config", "c", defaultConfigFile, "config file (default: "+defaultConfigFile+")")
	return &cmd
}

func Serve() {
	initServer()
	gin.SetMode(viper.GetString("ginmode"))
	server.GetRouter().Run(viper.GetString("address"))
}

func initServer() {
	initConfig()
	initDatabase()
	initRedis()
}

func initConfig() {
	viper.SetDefault("address", ":5000")
	viper.SetDefault("redis", ":6379")
	viper.SetDefault("ginmode", gin.ReleaseMode)

	viper.SetConfigFile(configFile)
	viper.SetEnvPrefix(envPrefix)
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		log.Infof("Using config file: %s", viper.ConfigFileUsed())
	} else {
		log.Fatalf("unable to load config: %v", err)
	}
}

func initDatabase() {
	db, err := sql.Open("postgres", "dbname=fun user=abc")
	if err != nil {
		log.Fatalf("unable to connect db: %v", err)
	}
	log.Infof("database connected!")
	boil.SetDB(db)
}

func initRedis() {
	err := rediskv.Connect(viper.GetString("redis"))
	if err != nil {
		log.Fatalf("unable to connect redis: %v", err)
	}
	log.Infof("redis connected!")
}
