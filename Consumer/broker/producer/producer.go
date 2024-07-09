package producer

import (
	"consumer/configuration"
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var Producer = create(configuration.Configuration)

func create(cfg configuration.AppConfiguration) *kafka.Producer {

	kafkaBrokers := os.Getenv("KAFKA_BROKER")

	if kafkaBrokers == "" {
		kafkaBrokers = cfg.BrokerServers
	}

	if kafkaBrokers == "" {
		panic("set KAFKA_BROKERS env variable or app config")
	}

	fmt.Println("init producer with app config")
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": kafkaBrokers,
		"client.id":         "producer",
		"acks":              "all"})

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}
	return p
}

/* func InitProducerWithProducerConfig(cfg kafka.ConfigMap) *kafka.Producer {
	fmt.Println("init producer with config")
	p, err := kafka.NewProducer(&cfg)
	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}
	return p
} */
