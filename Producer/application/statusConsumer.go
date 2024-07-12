package application

import (
	"microservice_go/types"
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

var userReqMap map[string]string

func init() {
	userReqMap = make(map[string]string)
}

func GetRequestMap() *map[string]string {
	return &userReqMap
}

func AddRequestToMap(reqId string, userId string) {
	userReqMap[reqId] = userId
}

func DeleteRequestFromMap(reqId string) {
	delete(userReqMap, reqId)
}

func StatusCalculatedHandler(msg StatusChangeResBrokerMsg) {
	reqId := msg.ReqId
	if userReqMap[reqId] == "" {
		return
	}
	println("Calculated", reqId)
	DeleteRequestFromMap(reqId)
}
