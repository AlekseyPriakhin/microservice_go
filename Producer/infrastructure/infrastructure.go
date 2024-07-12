package infrastructure

import (
	"microservice_go/application"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func HandleBrokerMsg(cmr *kafka.Consumer) {
	run := true

	for run {
		ev := cmr.Poll(100)
		switch e := ev.(type) {
		case *kafka.Message:
			application.HandleMessage(e)
		case kafka.Error:
			println(&e, ev)
		}

		time.Sleep(2 * time.Second)
	}
}
