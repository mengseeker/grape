package kafka

import (
	"context"
	"fmt"
	"grape/internal/v/worker"
	"grape/pkg/logger"
	"strings"
	"time"

	"github.com/Shopify/sarama"
)

const (
	kafkaVersion = "2.1.1"
	kafkaGroup   = "logtrans"
)

type kafkaConsumer struct {
	Name    string
	Brokers string
	Topics  string
	Log     logger.Logger

	w             worker.Worker
	saramaConfig  *sarama.Config
	consumerGroup sarama.ConsumerGroup
}

func (k *kafkaConsumer) Run(ctx context.Context, w worker.Worker) error {
	saramaConfig := sarama.NewConfig()
	k.saramaConfig = saramaConfig
	saramaConfig.ClientID = fmt.Sprintf("%s_%d", kafkaGroup, time.Now().Nanosecond())
	var err error
	saramaConfig.Version, _ = sarama.ParseKafkaVersion(kafkaVersion)
	saramaConfig.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	client, err := sarama.NewConsumerGroup(strings.Split(k.Brokers, ","), kafkaGroup, saramaConfig)
	if err != nil {
		return fmt.Errorf("creating consumer group client: %v", err)
	}
	k.consumerGroup = client

	k.w = w
	for {
		if err := k.consumerGroup.Consume(ctx, strings.Split(k.Topics, ","), k); err != nil {
			k.Log.Errorf("consumer %s: %v", k.Name, err)
		}
		if ctx.Err() != nil {
			break
		}
	}

	k.Log.Infof("kafka consumer %q closing...", k.Name)
	if err := k.consumerGroup.Close(); err != nil {
		k.Log.Fatalf("Error closing client: %v", err)
	}

	return nil
}

func (k *kafkaConsumer) Setup(sarama.ConsumerGroupSession) error {
	k.Log.Infof("Sarama consumer %s up and running!...", k.Name)
	return nil
}

func (c *kafkaConsumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (k *kafkaConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		k.w.Do(message.Value)
		session.MarkMessage(message, "")
	}
	return nil
}
