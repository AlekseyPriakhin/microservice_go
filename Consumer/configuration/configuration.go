package configuration

import (
	"encoding/json"
	"os"
)

type AppConfiguration struct {
	BrokerServers string `json:"brokerServers"`
}

func MustCreate() AppConfiguration {
	file, e := os.Open("configuration.json")

	if e != nil {
		panic(e)
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	configuration := AppConfiguration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		panic(err)
	}

	return configuration
}
