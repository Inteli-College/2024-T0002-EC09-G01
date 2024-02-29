package subscriber

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func Subscribe(topic string, client mqtt.Client, handler mqtt.MessageHandler) {
	token := client.Subscribe(topic, 1, handler)
	token.Wait()
	if token.Error() != nil {
		fmt.Printf("Failed to subscribe to topic: %v", token.Error())
		panic(token.Error())
	}
	fmt.Printf("\nSubscribed to topic: %s\n", topic)
}
