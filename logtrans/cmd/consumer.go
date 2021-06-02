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
	viper.SetDefault("logtrans.kafka.topic", "nxmc.logs")
	viper.SetDefault("logtrans.kafka.group_es", "logtrans_group_es")
	viper.SetDefault("logtrans.kafka.group_influxdb", "logtrans_group_influxdb")
	viper.SetDefault("logtrans.kakfa.version", "2.1.1")
	viper.SetDefault("lograns.kafka.verbose", false)
	viper.SetDefault("lograns.kafka.assignor", "range")

	viper.SetDefault("logtrans.es.consumer_num", 1)
	viper.SetDefault("logtrans.influxdb.consumer_num", 1)

	share.InitConfig(configFile, log)

	if viper.GetBool("lograns.kafka.verbose") {
		sarama.Logger = stdlog.New(os.Stdout, "[sarama] ", stdlog.LstdFlags)
	}
}

func RunConsumers() {
	RunEsConsumers()
	RunInfluxdbConsumers()

	wg.Wait()
}

func RunEsConsumers() {
	consumer_num := viper.GetInt("logtrans.es.consumer_num")
	kafka := viper.GetString("logtrans.kafka.address")
	es := viper.GetString("logtrans.es.address")
	esGroup := viper.GetString("logtrans.kafka.group_es")
	topic := viper.GetString("logtrans.kafka.topic")
	version := viper.GetString("logtrans.kakfa.version")
	assignor := viper.GetString("logtrans.kafka.assignor")
	env := viper.GetString("environment_code")
	cluster := viper.GetString("cluster_code")
	batchSize := viper.GetInt("logtrans.es.batch_size")
	interval := viper.GetInt("logtrans.es.interval")

	for i := 0; i < consumer_num; i++ {
		consumerName := "logtrans_es_" + strconv.FormatInt(int64(i), 10)
		esClient, err := worker.NewEsClient(es, env, cluster, log)
		if err != nil {
			log.Fatalf("connect elasticsearch err: %v", err)
		}
		runner := worker.NewRunner(esClient, batchSize, interval)
		consumer, err := worker.NewKafkaConsumer(
			kafka, assignor, esGroup, consumerName, topic, version,
			runner.NewConsume(), log,
		)
		if err != nil {
			log.Fatalf("faild to create kafka consumer, err: %v", err)
		}
		wg.Add(2)
		go func() { runner.RefreshLoop() }()
		go func() { defer wg.Done(); consumer.Run() }()
	}

}

func RunInfluxdbConsumers() {
	consumer_num := viper.GetInt("logtrans.influxdb.consumer_num")
	kafka := viper.GetString("logtrans.kafka.address")
	influx := viper.GetString("logtrans.influxdb.address")
	influxGroup := viper.GetString("logtrans.kafka.group_influxdb")
	topic := viper.GetString("logtrans.kafka.topic")
	version := viper.GetString("logtrans.kakfa_.version")
	assignor := viper.GetString("logtrans.kafka.assignor")
	env := viper.GetString("environment_code")
	cluster := viper.GetString("cluster_code")
	batchSize := viper.GetInt("logtrans.influxdb.batch_size")
	interval := viper.GetInt("logtrans.influxdb.interval")

	for i := 0; i < consumer_num; i++ {
		consumerName := "logtrans_influxdb_" + strconv.FormatInt(int64(i), 10)
		infClient, err := worker.NewInfClient(influx, env, cluster, log)
		if err != nil {
			log.Fatalf("connect influxdb err: %v", err)
		}
		runner := worker.NewRunner(infClient, batchSize, interval)
		consumer, err := worker.NewKafkaConsumer(
			kafka, assignor, influxGroup, consumerName, topic, version,
			runner.NewConsume(), log,
		)
		if err != nil {
			log.Fatalf("faild to create kafka consumer, err: %v", err)
		}
		wg.Add(2)
		go func() { runner.RefreshLoop() }()
		go func() { defer wg.Done(); consumer.Run() }()
	}
}
