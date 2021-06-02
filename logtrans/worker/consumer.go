package worker

import (
	"context"
	"fmt"
	"grape/pkg/logger"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/Shopify/sarama"
)

type Consumer interface {
	Run()
	Stop()
}

type Message = sarama.ConsumerMessage

type kafkaConsumer struct {
	Name    string
	topic   string
	log     logger.Logger
	consume Consume

	consumerGroup sarama.ConsumerGroup
	ctx           context.Context
	cancel        context.CancelFunc
	ready         chan bool
}

func NewKafkaConsumer(brokers, assignor, group, name, topic, version string,
	consume Consume, log logger.Logger) (*kafkaConsumer, error) {

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
	c.ready = make(chan bool)

	c.ctx, c.cancel = context.WithCancel(context.Background())
	client, err := sarama.NewConsumerGroup(strings.Split(brokers, ","), group, config)
	if err != nil {
		return nil, fmt.Errorf("creating consumer group client: %v", err)
	}
	if consume != nil {
		c.consume = consume
	}
	c.consumerGroup = client
	c.log = log

	return c, nil
}

func (c *kafkaConsumer) Run() {
	if c.consume == nil {
		return
	}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if err := c.consumerGroup.Consume(c.ctx, strings.Split(c.topic, ","), c); err != nil {
				c.log.Errorf("consumer: %v", err)
			}
			if c.ctx.Err() != nil {
				return
			}
			c.ready = make(chan bool)
		}
	}()

	<-c.ready // Await till the consumer has been set up
	c.log.Infof("Sarama consumer %s up and running!...", c.Name)

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-c.ctx.Done():
		c.log.Info("terminating: context cancelled")
	case <-sigterm:
		c.log.Info("terminating: via signal")
	}
	c.cancel()
	wg.Wait()
	if err := c.consumerGroup.Close(); err != nil {
		c.log.Fatalf("Error closing client: %v", err)
	}
}

func (c *kafkaConsumer) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(c.ready)
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

func FindHeader(hs []*sarama.RecordHeader, key string) string {
	for _, h := range hs {
		if string(h.Key) == key {
			return string(h.Value)
		}
	}
	return ""
}
