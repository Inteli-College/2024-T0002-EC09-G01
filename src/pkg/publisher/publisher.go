package publisher

import (
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func Publish(content string, topic string, client MQTT.Client) {
	token := client.Publish(topic, 0, false, content)
	token.Wait()
	if token.Error() != nil {
		fmt.Printf("Failed to publish to topic: %s", topic)
		panic(token.Error())
	}
}
