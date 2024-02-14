package common

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const Broker = "broker.hivemq.com:1883"
const IdPublisher = "go-mqtt-publisher"
const IdSubscriber = "go-mqtt-subscriber"

var Handler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received: %s on topic %s\n", msg.Payload(), msg.Topic())
	return
}

func CreateClient(broker string, id string, callback_handler mqtt.MessageHandler) mqtt.Client {

	opts := mqtt.NewClientOptions().AddBroker(broker)
	opts.SetClientID(id)
	opts.SetDefaultPublishHandler(callback_handler)

	return mqtt.NewClient(opts)
}
