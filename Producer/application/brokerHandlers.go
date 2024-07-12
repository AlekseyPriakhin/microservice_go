package application

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func HandleMessage(ev *kafka.Message) {

	switch *ev.TopicPartition.Topic {

	case "status_res":
		{
			msg := extractData[StatusChangeResBrokerMsg](ev)
			StatusCalculatedHandler(msg)
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
