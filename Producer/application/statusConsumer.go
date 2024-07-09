package application

import (
	"microservice_go/types"

	"github.com/google/uuid"
)

type StatusChangeReqBrokerMsg struct {
	ReqId  string       `json:"reqId"`
	Course types.Course `json:"course"`
}

var UserRequestMap = make(map[string]uuid.UUID)

func StatusCalculatedHandler(msg StatusChangeReqBrokerMsg) {
}
