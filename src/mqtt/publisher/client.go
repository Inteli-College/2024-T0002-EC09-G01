package main

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var handler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received: %s on topic %s\n", msg.Payload(), msg.Topic())
	return
}

func CreateClient(broker string, id string, callback_handler mqtt.MessageHandler) mqtt.Client {

	opts := mqtt.NewClientOptions().AddBroker(broker)
	opts.SetClientID(id)
	opts.SetDefaultPublishHandler(callback_handler)

	return mqtt.NewClient(opts)
}
