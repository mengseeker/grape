package cmd

import (
	"context"
	"fmt"
	"grape/logtrans/cmd/config"
	"grape/logtrans/logs"
	"grape/logtrans/worker"
	"grape/pkg/share"
	"grape/pkg/util"
	stdlog "log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/Shopify/sarama"
	"github.com/spf13/cobra"
)

var configFile string
var wg sync.WaitGroup

func NewConsumerCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "consumer",
		Short: ".",
		Long:  `.`,
		Run: func(cmd *cobra.Command, args []string) {
			initConfig(configFile)
			Run()
		},
	}
	cmd.Flags().StringVarP(&configFile, "config", "c", share.DefaultCfgFile, "config file")
	return &cmd
}

func initConfig(configFile string) {
	share.InitConfig(configFile, log)
}

func Run() {
	ctx, cancel := context.WithCancel(context.Background())
	var tran logs.Transmitter
	HandleSource(tran, ctx)
	RunConsumer(tran, ctx)

	go func() {
		sigterm := make(chan os.Signal, 1)
		signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
		<-sigterm
		log.Warn("terminating: via signal")
		cancel()
		<-time.After(10 * time.Second)
		log.Warnf("timeout for quit, force exit...")
		syscall.Exit(1)
	}()

	wg.Wait()
}

func HandleSource(tran logs.Transmitter, ctx context.Context) {
	// kafka
	if config.IsConfigSourceKafka() {
		err := config.CheckSourceKafkaConfig()
		if err != nil {
			log.Fatal(err)
		}
		if share.IsDebug() {
			sarama.Logger = stdlog.New(os.Stdout, "[sarama] ", stdlog.LstdFlags)
		}
		address, topics, group, version, assignor := config.GetSourceKafkaConfig()
		consumerName := fmt.Sprintf("%s_%s", strings.NewReplacer(",", "_").Replace(topics), util.CreateRandomString(5))
		consumer, err := worker.NewKafkaConsumer(address, assignor, group, consumerName, topics, version, log)
		if err != nil {
			log.Fatalf("faild to create kafka consumer: %v", err)
		}
		wg.Add(1)
		go func() { defer wg.Done(); consumer.Run(tran, ctx) }()
	}
}

func RunConsumer(tran logs.Transmitter, ctx context.Context) {
	env := share.GetEnvironment()
	cluster := share.GetCluster()
	// elasticsearch
	if config.IsConfigOutEs() {
		err := config.CheckOutEsConfig()
		if err != nil {
			log.Fatal(err)
		}
		address, batch_size, interval, types := config.GetOutEsConfig()
		esClient, err := worker.NewEsClient(address, env, cluster, log)
		if err != nil {
			log.Fatalf("faild to connecting elasticsearch: %v", err)
		}
		esRunner := worker.NewRunner(esClient, batch_size, interval, log)
		tran.Distribute(types, esRunner)
		wg.Add(1)
		go func() { defer wg.Done(); esRunner.RefreshLoop(ctx) }()
	}

	// influxdb
	if config.IsConfigOutInfluxDB() {
		err := config.CheckOutInfluxDBConfig()
		if err != nil {
			log.Fatal(err)
		}
		address, batch_size, interval, types := config.GetOutInfluxDBConfig()
		influxClient, err := worker.NewInfClient(address, env, cluster, log)
		if err != nil {
			log.Fatalf("faild to connecting influxdb: %v", err)
		}
		esRunner := worker.NewRunner(influxClient, batch_size, interval, log)
		tran.Distribute(types, esRunner)
		wg.Add(1)
		go func() { defer wg.Done(); esRunner.RefreshLoop(ctx) }()
	}
}
