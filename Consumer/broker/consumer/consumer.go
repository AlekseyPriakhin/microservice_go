package consumer

import (
	"consumer/configuration"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var Consumer = create(configuration.Configuration)

func create(cfg configuration.AppConfiguration) *kafka.Consumer {
	kafkaBrokers := os.Getenv("KAFKA_BROKER")

	if kafkaBrokers == "" {
		kafkaBrokers = cfg.BrokerServers
	}

	if kafkaBrokers == "" {
		panic("set KAFKA_BROKERS env variable")
	}

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": kafkaBrokers,
		"group.id":          "foo"})

	if err != nil {
		panic(err)
	}

	return consumer
}
