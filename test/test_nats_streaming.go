package main

import (
	"fmt"
	"io/ioutil"

	"github.com/Gunga-D/taskL0/internal/config"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
)

func formTestingData(pathToTestingData string) ([][]byte, error) {
	var result [][]byte

	testingDataDir, err := ioutil.ReadDir(pathToTestingData)
	if err != nil {
		return nil, err
	}

	for _, file := range testingDataDir {
		content, err := ioutil.ReadFile(pathToTestingData + "/" + file.Name())
		if err != nil {
			return nil, err
		}

		result = append(result, content)
	}
	return result, nil
}

func main() {
	config, err := config.LoadNATStreamingConfig("./configs", "natstreaming_config")
	if err != nil {
		logrus.Fatalf(err.Error())
	}

	testingData, err := formTestingData("./test/data")
	if err != nil {
		logrus.Fatalf(err.Error())
	}

	natsUrl := fmt.Sprintf("nats://%s:%s", config.Host, config.Port)
	stream, err := stan.Connect(config.ClusterID, "client-321", stan.NatsURL(natsUrl), stan.MaxPubAcksInflight(1000))
	if err != nil {
		logrus.Fatalf(err.Error())
	}

	for _, portion := range testingData {
		stream.Publish("orders", portion)
	}
}
