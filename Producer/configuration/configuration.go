package configuration

import (
	"encoding/json"
	"fmt"
	"os"
)

type AppConfiguration struct {
	BrokerServers string `json:"brokerServers"`
}

var Configuration = readConfiguration()

func readConfiguration() AppConfiguration {
	file, _ := os.Open("configuration.json")
	print(file)
	defer file.Close()

	decoder := json.NewDecoder(file)
	configuration := AppConfiguration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}

	return configuration
}
