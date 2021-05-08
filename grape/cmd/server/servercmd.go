package server

import (
	"grape/grape/models"
	"grape/grape/pkg/postgresdb"
	"grape/grape/server"
	"grape/pkg/redispool"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	InitConfig(configFile)
	InitDatabase()
	InitRedis()
}

func InitConfig(cfg string) {
	viper.SetDefault("address", ":5000")
	viper.SetDefault("redis", ":6379")
	viper.SetDefault("ginmode", gin.ReleaseMode)
	viper.SetDefault("automigrate", false)

	viper.SetConfigFile(cfg)
	viper.SetEnvPrefix(envPrefix)
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		log.Infof("Using config file: %s", viper.ConfigFileUsed())
	} else {
		log.Fatalf("unable to load config: %v", err)
	}
}

func InitDatabase() {
	err := postgresdb.Connect(viper.GetString("database"))
	if err != nil {
		log.Fatalf("unable to connect db: %v", err)
	}

	if viper.GetBool("automigrate") {
		autoMigrate()
	}

	log.Infof("database connected!")
}

func InitRedis() {
	err := redispool.Connect(viper.GetString("redis"))
	if err != nil {
		log.Fatalf("unable to connect redis: %v", err)
	}
	log.Infof("redis connected!")
}

func autoMigrate() {
	err := postgresdb.GetDB().AutoMigrate(
		&models.Namespace{},
		&models.EtcdLink{},
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
