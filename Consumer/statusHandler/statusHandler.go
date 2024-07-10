package statushandler

import (
	"consumer/types"
	"consumer/utils"
	"time"
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

var ReqQueue = utils.CreateQueue[StatusChangeReqBrokerMsg]()
var ResQueue = utils.CreateQueue[StatusChangeResBrokerMsg]()

func ReqHandler(msg StatusChangeReqBrokerMsg) {
	println("Добавляю в очередь новое сообщение", msg.ReqId)
	ReqQueue.Enqueue(msg)
}

func QueueHandler() {
	run := true
	for run {

		if ReqQueue.IsEmpty() {
			//println("Сообщений нет")
			time.Sleep(5 * time.Second)
			continue
		}

		msg := ReqQueue.Dequeue()
		println("Извлеченное сообщение: ", msg.ReqId)
		time.Sleep(2 * time.Second)
		println("Обрабатываю...")
		time.Sleep(12 * time.Second)

		res := types.ValidateStages(msg.Course.Stages)
		println("Обработанное сообщение: ", msg.ReqId)
		ResQueue.Enqueue(StatusChangeResBrokerMsg{ReqId: msg.ReqId, Course: msg.Course, Result: res})
	}
}
