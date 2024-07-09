package broker

type BrokerMessage[T any] struct {
	UUID      string `json:"uuid"`
	Timestamp int    `json:"timestamp"`
	Data      T      `json:"data"`
}
