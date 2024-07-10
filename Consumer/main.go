package main

import (
	"consumer/broker/consumer"
	"consumer/broker/producer"
	statushandler "consumer/statusHandler"
	"encoding/json"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	brokerCmr := consumer.Consumer

	brokerCmr.SubscribeTopics([]string{"status_req"}, nil)

	go statushandler.QueueHandler()
	run := true
	for run {

		ev := brokerCmr.Poll(100)
		switch e := ev.(type) {
		case *kafka.Message:
			consumer.HandleMessage(e)
		case kafka.Error:
			consumer.HandleError(&e, ev)
			run = false
		}

		time.Sleep(2 * time.Second)

		//Обработка очереди обработанных сообщений
		if !statushandler.ResQueue.IsEmpty() {
			data := statushandler.ResQueue.Dequeue()

			topic := "status_res"
			msg, _ := json.Marshal(data)

			producer.Producer.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &topic},
				Value:          []byte(msg),
			}, nil)
		}

	}

	brokerCmr.Close()
}
