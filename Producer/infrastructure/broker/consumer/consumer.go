package consumer

import (
	"microservice_go/configuration"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func MustCreate(cfg configuration.AppConfiguration) *kafka.Consumer {
	kafkaBrokers := os.Getenv("KAFKA_BROKER")

	if kafkaBrokers == "" {
		kafkaBrokers = cfg.BrokerServers
	}

	if kafkaBrokers == "" {
		panic("set KAFKA_BROKERS env variable")
	}

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":        kafkaBrokers,
		"allow.auto.create.topics": true,
		"group.id":                 "foo"})

	if err != nil {
		panic(err)
	}

	return consumer
}
