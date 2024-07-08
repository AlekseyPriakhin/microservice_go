package main

import (
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	kafkaBrokers := os.Getenv("KAFKA_BROKER")
	if kafkaBrokers == "" {
		panic("set KAFKA_BROKERS env variable")
	}

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": kafkaBrokers,
		"group.id":          "foo",
		"auto.offset.reset": "smallest"})

	if err != nil {
		panic(err)
	}

	consumer.SubscribeTopics([]string{"course"}, nil)
	run := true
	for run {
		ev := consumer.Poll(100)
		switch e := ev.(type) {
		case *kafka.Message:
			fmt.Printf("Message on %s: %s\n", e.TopicPartition, string(e.Value))
		case kafka.Error:
			fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
			run = false
		default:
		}
	}

	consumer.Close()
}
