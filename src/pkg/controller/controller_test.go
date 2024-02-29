package controller

import (
	DefaultClient "2024-T0002-EC09-G01/src/pkg/common"
	"fmt"
	"math/rand/v2"
	"testing"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func TestController(t *testing.T) {

	client := DefaultClient.CreateClient(DefaultClient.Broker, DefaultClient.IdPublisher, DefaultClient.Handler)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	var messageChannel = make(chan mqtt.Message)
	t.Run("TestPublishFields", func(t *testing.T) {

		client.Subscribe("sensors/+/+", 1, func(client mqtt.Client, message mqtt.Message) {
			fmt.Print("Subscribed to topic: sensors/+/+")
			resultado := message.Payload()
			if messageChannel != nil {
				t.Error("Test")
			}
			fmt.Printf("%s", resultado)
		})

		for sensorType := 0; sensorType < len(topics); sensorType++ {

			var topic string
			var payload string

			var id = rand.IntN(100)

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
	})
}
