package cmd

// import (
// 	"grape/pilot/apiv3"
// 	"grape/pkg/share"

// 	"github.com/spf13/cobra"
// 	"github.com/spf13/viper"
// )

// const ()

// var (
// 	cfgFile string
// )

// func NewServerCmd() *cobra.Command {
// 	cmd := cobra.Command{
// 		Use:     "server",
// 		Aliases: []string{"s"},
// 		Short:   "start pilot",
// 		Long:    `start pilot`,
// 		Run: func(cmd *cobra.Command, args []string) {
// 			Serve()
// 			<-make(chan int)
// 		},
// 	}
// 	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", share.DefaultCfgFile, "config file")
// 	return &cmd
// }

// func Serve() {
// 	initConfig()

// 	// xds server
// 	xdsAddress := viper.GetString("pilot.address")
// 	apiv3.Serve(xdsAddress)
// }

// func initConfig() {
// 	viper.SetDefault("pilot.address", "0.0.0.0:15010")

// 	share.InitConfig(cfgFile, log)
// }
