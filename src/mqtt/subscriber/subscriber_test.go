package main

import (
	DefaultClient "mqtt/src/common"
	"testing"
)

func TestSubscriber(t *testing.T) {
	t.Run("Subscription to topic", func(t *testing.T) {
		client := DefaultClient.CreateClient(DefaultClient.Broker, DefaultClient.IdSubscriber, DefaultClient.Handler)

		if token := client.Connect(); token.Wait() && token.Error() != nil {
			t.Error(token.Error())
		}

		if token := client.Subscribe("sensors/SPS30", 1, nil); token.Wait() && token.Error() != nil {
			t.Error(token.Error())
			return
		}

		t.Log("Subscribed successfully to Topic")
	})
}
