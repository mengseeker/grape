package cmd

import (
	"grape/grape/models"
	"grape/grape/server"
	"grape/pkg/redispool"
	"grape/pkg/share"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configFile string
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
	cmd.Flags().StringVarP(&configFile, "config", "c", share.DefaultCfgFile, "config file")

	return &cmd
}

func Serve() {
	initServer()
	gin.SetMode(viper.GetString("grape.ginmode"))
	server.GetRouter().Run(viper.GetString("grape.address"))
}

func initServer() {
	InitConfig(configFile)
	InitDatabase()
	InitRedis()
}

func InitConfig(cfg string) {
	viper.SetDefault("grape.address", ":5000")
	viper.SetDefault("redis.address", ":6379")
	viper.SetDefault("grape.ginmode", gin.ReleaseMode)
	viper.SetDefault("grape.automigrate", false)

	share.InitConfig(configFile, log)
}

func InitDatabase() {
	err := models.Connect(viper.GetString("grape.database"))
	if err != nil {
		log.Fatalf("unable to connect db: %v", err)
	}

	if viper.GetBool("automigrate") {
		autoMigrate()
	}

	log.Infof("database connected!")
}

func InitRedis() {
	err := redispool.Connect(viper.GetString("redis.address"))
	if err != nil {
		log.Fatalf("unable to connect redis: %v", err)
	}
	log.Infof("redis connected!")
}

func autoMigrate() {
	err := models.GetDB().AutoMigrate(
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
