package controller

import (
	DefaultClient "2024-T0002-EC09-G01/src/pkg/common"
	"testing"
)

func TestController(t *testing.T) {

	client := DefaultClient.CreateClient(DefaultClient.Broker, DefaultClient.IdPublisher, DefaultClient.Handler)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	var messageChannel = make(chan mqtt.message)
	t.Run("TestPublishFields", func(t *testing.T) {
		client.Subscribe("sensors", 1, func(client mqtt.Client, message mqtt.message) {
			messageChannel <- message
		})
	})
}
