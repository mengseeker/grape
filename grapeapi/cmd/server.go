package cmd

import (
	"grape/grapeapi/models"
	"grape/grapeapi/server"
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
	serverConfig := server.ServerConfig{
		GinMode: viper.GetString("grape.ginmode"),
		ApiAddr: viper.GetString("grape.api_address"),
		MCPAddr: viper.GetString("grape.mcp_address"),
	}
	server.Serve(serverConfig)
}

func initServer() {
	InitConfig(configFile)
	InitDatabase()
	InitRedis()
}

func InitConfig(cfg string) {
	viper.SetDefault("grape.api_address", ":5000")
	viper.SetDefault("grape.mcp_address", ":15010")
	viper.SetDefault("redis.address", ":6379")
	viper.SetDefault("grape.ginmode", gin.ReleaseMode)
	// viper.SetDefault("grape.automigrate", false)

	share.InitConfig(configFile, log)
}

func InitDatabase() {
	err := models.Connect(viper.GetString("grape.database"))
	if err != nil {
		log.Fatalf("unable to connect db: %v", err)
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
