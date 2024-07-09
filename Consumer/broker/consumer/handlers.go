package consumer

import (
	statushandler "consumer/statusHandler"
	"consumer/types"
	"encoding/json"
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func HandleMessage(ev *kafka.Message) {
	println("handle message")

	switch *ev.TopicPartition.Topic {
	case "course":
		courseMsgHandler(extractData[types.Course](ev))
	case "status_req":
		{
			println("handle status_req message")
			statushandler.ReqHandler(extractData[statushandler.StatusChangeReqBrokerMsg](ev))
		}
	}
}

func HandleError(ev *kafka.Error, e kafka.Event) {
	fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
}

func extractData[T any](msg *kafka.Message) T {
	var data T
	json.Unmarshal(msg.Value, &data)

	err := json.Unmarshal(msg.Value, &data)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%% Unmarshal error: %v\n", err)
		return data
	}
	return data
}

func courseMsgHandler(msg types.Course) {
}
