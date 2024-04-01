package controller

import (
	"fmt"
	"time"

	DefaultClient "2024-T0002-EC09-G01/src/pkg/common"
	pub_mics6814 "2024-T0002-EC09-G01/src/pkg/pub_mics6814"
	pub_rxwlib900 "2024-T0002-EC09-G01/src/pkg/pub_rxwlib900"
)

const (
	gasesSensorType     = 0
	radiationSensorType = 1
)

var topics = [2]string{"gases", "radiation"}

func Topic(sensorType int, id string) string {

	typeOfSensor := topics[sensorType]
	topic := fmt.Sprintf("sensor/%s/%s", typeOfSensor, id)
	return topic
}

func Payload(sensorType int, id string) string {
	if sensorType == gasesSensorType {
		return pub_mics6814.CreatePayloadGases(id)
	} else if sensorType == radiationSensorType {
		return pub_rxwlib900.CreatePayloadRadiation(id)
	}
	return "No message received"
}



func Controller(id string) {
	client := DefaultClient.CreateClient(fmt.Sprintf("publisher-%s", id), DefaultClient.Handler)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	for {
		for sensorType := 0; sensorType < len(topics); sensorType++ {

			var topic string
			var payload string

			switch sensorType {

			case gasesSensorType:
				topic = Topic(sensorType, id)
				payload = Payload(sensorType, id)
			case radiationSensorType:
				topic = Topic(sensorType, id)
				payload = Payload(sensorType, id)
			}

			token := client.Publish(topic, 1, false, payload)
			token.Wait()
			
			fmt.Printf("Published message in %s: %s\n", topic, payload)
		}

		time.Sleep(60 * time.Second)
	}
}
