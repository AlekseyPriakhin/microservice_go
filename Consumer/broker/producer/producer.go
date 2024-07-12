package producer

import (
	"consumer/configuration"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func MustCreate(cfg configuration.AppConfiguration) *kafka.Producer {

	kafkaBrokers := os.Getenv("KAFKA_BROKER")

	if kafkaBrokers == "" {
		kafkaBrokers = cfg.BrokerServers
	}

	if kafkaBrokers == "" {
		panic("set KAFKA_BROKERS env variable or app config")
	}

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": kafkaBrokers,
		"client.id":         "producer",
		"acks":              "all"})

	if err != nil {
		panic(err)
	}
	return p
}
