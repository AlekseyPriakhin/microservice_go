package statushandler

import (
	"consumer/types"
	"encoding/json"
	"sync"
	"time"

	"github.com/AlekseyPriakhin/queue"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type StatusChangeReqBrokerMsg struct {
	ReqId  string       `json:"reqId"`
	Course types.Course `json:"course"`
}

type StatusChangeResBrokerMsg struct {
	ReqId  string       `json:"reqId"`
	Course types.Course `json:"course"`
	Result bool         `json:"result"`
}

var ReqQueue = queue.CreateQueue[StatusChangeReqBrokerMsg]()
var ResQueue = queue.CreateQueue[StatusChangeResBrokerMsg]()
var resQueueMtx = sync.Mutex{}

func ReqHandler(msg StatusChangeReqBrokerMsg) {
	ReqQueue.Enqueue(msg)
}

func ReqQueueHandler() {
	run := true
	ch := make(chan struct{}, 2)
	for run {

		if ReqQueue.IsEmpty() {
			time.Sleep(5 * time.Second)
			continue
		}

		go handleMsg(ReqQueue.Dequeue(), &resQueueMtx, ch)
	}
}

func handleMsg(msg StatusChangeReqBrokerMsg, mutex *sync.Mutex, ch chan struct{}) {
	ch <- struct{}{}

	time.Sleep(15 * time.Second)
	res := types.ValidateStages(msg.Course.Stages)

	mutex.Lock()
	ResQueue.Enqueue(StatusChangeResBrokerMsg{ReqId: msg.ReqId, Course: msg.Course, Result: res})
	mutex.Unlock()
	<-ch
}

func ResQueueHandler(producer *kafka.Producer) {
	run := true
	for run {
		time.Sleep(2 * time.Second)

		var data StatusChangeResBrokerMsg = StatusChangeResBrokerMsg{}
		dequeueItem(ResQueue, &resQueueMtx, &data)

		if data.Course.Id == 0 {
			continue
		}
		println("dequeue item", data.ReqId, data.Course.Id)

		msg, _ := json.Marshal(data)
		sendToTopic("status_res", msg, producer)
	}
}

func sendToTopic(topic string, msg []byte, producer *kafka.Producer) {
	producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic},
		Value:          []byte(msg),
	}, nil)
}

func dequeueItem[T any](q queue.Queue[T], mtx *sync.Mutex, d *T) {
	mtx.Lock()
	if !ResQueue.IsEmpty() {
		*d = q.Dequeue()
	}
	mtx.Unlock()
}
