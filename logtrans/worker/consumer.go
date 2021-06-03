package worker

import (
	"context"
	"fmt"
	"grape/pkg/logger"
	"strings"
	"sync"

	"github.com/Shopify/sarama"
)

type Message = sarama.ConsumerMessage

type kafkaConsumer struct {
	Name    string
	topic   string
	log     logger.Logger
	consume Consume

	consumerGroup sarama.ConsumerGroup
}

func NewKafkaConsumer(brokers, assignor, group, name, topic, version string, log logger.Logger) (*kafkaConsumer, error) {

	c := new(kafkaConsumer)
	config := sarama.NewConfig()

	c.topic = topic
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

func (c *kafkaConsumer) Run(r Runner, ctx context.Context) {
	if r == nil {
		return
	}
	c.consume = r.NewConsume()
	// 启动runner
	runCtx, runCancel := context.WithCancel(context.Background())
	runWg := sync.WaitGroup{}
	runWg.Add(1)
	go func() {
		defer runWg.Done()
		r.RefreshLoop(runCtx)
	}()
	for {
		if err := c.consumerGroup.Consume(ctx, strings.Split(c.topic, ","), c); err != nil {
			c.log.Errorf("consumer: %v", err)
		}
		if ctx.Err() != nil {
			break
		}
	}

	if err := c.consumerGroup.Close(); err != nil {
		c.log.Fatalf("Error closing client: %v", err)
	}
	runCancel()
	runWg.Wait()
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
		c.consume(message)
		session.MarkMessage(message, "")
	}
	return nil
}
