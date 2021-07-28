package worker

import (
	"context"
	"fmt"
	"grape/logtrans/logs"
	"grape/pkg/logger"
	"strings"

	"github.com/Shopify/sarama"
)

type kafkaConsumer struct {
	Name   string
	topics string
	log    logger.Logger
	rec    logs.Receiver

	consumerGroup sarama.ConsumerGroup
}

func NewKafkaConsumer(brokers, assignor, group, name, topics, version string, log logger.Logger) (*kafkaConsumer, error) {
	c := new(kafkaConsumer)
	config := sarama.NewConfig()

	c.topics = topics
	config.ClientID = name
	c.Name = name
	var err error
	config.Version, err = sarama.ParseKafkaVersion(version)
	if err != nil {
		return nil, fmt.Errorf("parsing Kafka version: %v", err)
	}
	switch assignor {
	case "sticky":
		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategySticky
	case "roundrobin":
		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	case "range":
		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	default:
		return nil, fmt.Errorf("unrecognized consumer group partition assignor: %s", assignor)
	}
	// config.Consumer.Offsets.Initial = sarama.OffsetOldest
	client, err := sarama.NewConsumerGroup(strings.Split(brokers, ","), group, config)
	if err != nil {
		return nil, fmt.Errorf("creating consumer group client: %v", err)
	}
	c.consumerGroup = client
	c.log = log

	return c, nil
}

func (c *kafkaConsumer) Run(rec logs.Receiver, ctx context.Context) {
	if rec == nil {
		return
	}
	c.rec = rec
	for {
		if err := c.consumerGroup.Consume(ctx, strings.Split(c.topics, ","), c); err != nil {
			c.log.Errorf("consumer: %v", err)
		}
		if ctx.Err() != nil {
			break
		}
	}

	c.log.Infof("kafka consumer %q closing...", c.Name)
	if err := c.consumerGroup.Close(); err != nil {
		c.log.Fatalf("Error closing client: %v", err)
	}
}

func (c *kafkaConsumer) Setup(sarama.ConsumerGroupSession) error {
	c.log.Infof("Sarama consumer %s up and running!...", c.Name)
	return nil
}

func (c *kafkaConsumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (c *kafkaConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		c.rec.Receive(logs.Message{
			MessageType: message.Topic,
			Val:         message.Value,
		})
		session.MarkMessage(message, "")
	}
	return nil
}
