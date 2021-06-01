package cmd

import (
	"grape/logtrans/worker"
	"grape/pkg/share"
	stdlog "log"
	"os"
	"strconv"
	"sync"

	"github.com/Shopify/sarama"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
			RunConsumers()
		},
	}
	cmd.Flags().StringVarP(&configFile, "config", "c", share.DefaultCfgFile, "config file")
	return &cmd
}

func initConfig(configFile string) {
	// viper.SetDefault("logtrans.kafka", ":5000")
	// viper.SetDefault("logtrans.es", ":5000")
	// viper.SetDefault("logtrans.influxdb", ":5000")
	viper.SetDefault("logtrans.kafka_topic", "nxmc.logs")
	viper.SetDefault("logtrans.kafka_group_es", "logtrans_group_es")
	viper.SetDefault("logtrans.kafka_group_influxdb", "logtrans_group_influxdb")
	viper.SetDefault("logtrans.consumer_num", 1)
	viper.SetDefault("logtrans.kakfa_version", "2.1.1")
	viper.SetDefault("lograns.kafka_verbose", false)
	viper.SetDefault("lograns.kafka_assignor", "range")
	share.InitConfig(configFile, log)

	if viper.GetBool("lograns.kafka_verbose") {
		sarama.Logger = stdlog.New(os.Stdout, "[sarama] ", stdlog.LstdFlags)
	}
}

func RunConsumers() {
	consumer_num := viper.GetInt("logtrans.consumer_num")
	kafka := viper.GetString("logtrans.kafka")
	es := viper.GetString("logtrans.es")
	influx := viper.GetString("logtrans.influxdb")
	esGroup := viper.GetString("logtrans.kafka_group_es")
	influxGroup := viper.GetString("logtrans.kafka_group_influxdb")
	topic := viper.GetString("logtrans.kafka_topic")
	version := viper.GetString("logtrans.kakfa_version")
	assignor := viper.GetString("logtrans.kafka_assignor")

	for i := 0; i < consumer_num; i++ {
		wg.Add(2)
		go func(id int) {
			consumerName := "logtrans_es_" + strconv.FormatInt(int64(id), 10)
			esClient, err := worker.NewEsClient(es)
			if err != nil {
				log.Fatalf("connect elasticsearch err: %v", err)
			}
			consumer, err := worker.NewKafkaConsumer(
				kafka, assignor, esGroup, consumerName, topic, version,
				esClient.NewRunner(), log,
			)
			if err != nil {
				log.Fatalf("faild to create kafka consumer, err: %v", err)
			}
			defer wg.Done()
			consumer.Run()
		}(i)

		go func(id int) {
			consumerName := "logtrans_influxdb_" + strconv.FormatInt(int64(id), 10)
			esClient, err := worker.NewInfClient(influx)
			if err != nil {
				log.Fatalf("connect influxdb err: %v", err)
			}
			consumer, err := worker.NewKafkaConsumer(
				kafka, assignor, influxGroup, consumerName, topic, version,
				esClient.NewRunner(), log,
			)
			if err != nil {
				log.Fatalf("faild to create kafka consumer, err: %v", err)
			}
			defer wg.Done()
			consumer.Run()
		}(i)
	}

	wg.Wait()
}
