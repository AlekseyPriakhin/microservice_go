package main

import (
	"consumer/broker/consumer"
	"consumer/broker/producer"
	"consumer/configuration"
	statushandler "consumer/statusHandler"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	cfg := configuration.MustCreate()
	prd := producer.MustCreate(cfg)
	cmr := consumer.MustCreate(cfg)

	cmr.SubscribeTopics([]string{"status_req"}, nil)

	go statushandler.ReqQueueHandler()
	go statushandler.ResQueueHandler(prd)
	run := true
	for run {

		ev := cmr.Poll(100)
		switch e := ev.(type) {
		case *kafka.Message:
			consumer.HandleMessage(e)
		case kafka.Error:
			consumer.HandleError(&e, ev)
			run = false
		}

		time.Sleep(2 * time.Second)

	}

	cmr.Close()
}
