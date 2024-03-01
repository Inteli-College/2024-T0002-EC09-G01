package common

import (
	"testing"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func TestClient(t *testing.T) {
	t.Run("Create a Client", func(t *testing.T) {
		client := CreateClient(Broker, "go-mqtt-client", Handler)

		if token := client.Connect(); token.Wait() && token.Error() != nil {
			t.Error(token.Error())
		}
	})

	t.Run("Subscribe to topic", func(t *testing.T) {
		client := CreateClient(Broker, "go-mqtt-client-subscribe", Handler)

		if token := client.Connect(); token.Wait() && token.Error() != nil {
			t.Error(token.Error())
		}

		topic := "test/topic"

		token := client.Subscribe(topic, 1, func(client mqtt.Client, msg mqtt.Message){
			t.Logf("Received message on topic %s: %s", msg.Topic(), msg.Payload())
		})

		token.Wait()
		if token.Error() != nil{
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
		client := CreateClient(Broker, "go-mqtt-client-publish", Handler)

		if token := client.Connect(); token.Wait() && token.Error() != nil {
			t.Error(token.Error())
		}

		topic := "test/topic"

		payload := "Publish"

		token := client.Publish(topic, 1, false, payload)

		token.Wait()
		if token.Error() != nil{
			t.Error(token.Error())
		}
	})

	t.Run("Authorization test", func(t *testing.T){
		publishTopic := "test/publish"
		subscribeTopic := "test/subscribe"

		client := CreateClient(Broker, "go-mqtt-client-auth", Handler)

		if token := client.Connect(); token.Wait() && token.Error() != nil {
			t.Error(token.Error())
		}

		token := client.Publish(publishTopic, 1, false, "Authorization Test: Publish")
		token.Wait()
		if token.Error() != nil {
			t.Error(token.Error())
		}

		token := client.Subscribe(subscribeTopic, 1, nil)
		token.Wait()
		if token.Error() != nil {
			t.Error(token.Error())
		}
	})
}
