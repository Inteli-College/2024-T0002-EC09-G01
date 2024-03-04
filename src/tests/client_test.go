package testing

import (
	"testing"
	"time"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	common "2024-T0002-EC09-G01/src/pkg/common"
)

func TestClient(t *testing.T) {
	t.Run("Create a Client", func(t *testing.T) {
		client := common.CreateClient("go-mqtt-client", common.Handler)

		if token := client.Connect(); token.Wait() && token.Error() != nil {
			t.Error(token.Error())
		}
	})

	t.Run("Subscribe to topic", func(t *testing.T) {
		client := common.CreateClient("go-mqtt-client-subscribe", common.Handler)

		if token := client.Connect(); token.Wait() && token.Error() != nil {
			t.Error(token.Error())
		}

		topic := "test/topic"

		token := client.Subscribe(topic, 1, func(client mqtt.Client, msg mqtt.Message) {
			t.Logf("Received message on topic %s: %s", msg.Topic(), msg.Payload())
		})

		token.Wait()
		if token.Error() != nil {
			t.Error(token.Error())
		}

		time.Sleep(2 * time.Second)

		token = client.Unsubscribe(topic)
		token.Wait()
		if token.Error() != nil {
			t.Error(token.Error())
		}
	})

	t.Run("Publish message", func(t *testing.T) {
		client := common.CreateClient("go-mqtt-client-publish", common.Handler)

		if token := client.Connect(); token.Wait() && token.Error() != nil {
			t.Error(token.Error())
		}

		topic := "test/topic"

		payload := "Publish"

		token := client.Publish(topic, 1, false, payload)

		token.Wait()
		if token.Error() != nil {
			t.Error(token.Error())
		}
	})
}
