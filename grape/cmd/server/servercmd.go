package server

import (
	"database/sql"
	"grape/grape/server"
	"grape/pkg/redispool"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/volatiletech/sqlboiler/v4/boil"
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
	viper.SetDefault("boildebug", false)

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
	db, err := sql.Open("postgres", viper.GetString("database"))
	if err != nil {
		log.Fatalf("unable to connect db: %v", err)
	}
	boildebugmode := viper.GetString("boildebug")
	if boildebugmode != "" {
		boil.DebugMode = true
		if boildebugmode != "/dev/stdout" {
			fh, err := os.OpenFile(boildebugmode, os.O_CREATE | os.O_RDWR, 0644)
			if err != nil {
				log.Fatalf("init boil err: %v", err)
			}
			boil.DebugWriter = fh
		}
	}

	log.Infof("database connected!")
	boil.SetDB(db)
	boil.SetLocation(time.Local)
}

func InitRedis() {
	err := redispool.Connect(viper.GetString("redis"))
	if err != nil {
		log.Fatalf("unable to connect redis: %v", err)
	}
	log.Infof("redis connected!")
}
