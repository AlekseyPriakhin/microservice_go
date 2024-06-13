package infrastructure

import (
	"fmt"
	"microservice_go/configuration"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func InitProducerWithAppConfig(cfg configuration.AppConfiguration) *kafka.Producer {

	print(cfg.BrokerServers)
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":        cfg.BrokerServers,
		"client.id":                "producer",
		"allow.auto.create.topics": true,
		"acks":                     "all"})

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}
	return p
}

func InitProducerWithProducerConfig(cfg kafka.ConfigMap) *kafka.Producer {
	p, err := kafka.NewProducer(&cfg)
	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}
	return p
}
