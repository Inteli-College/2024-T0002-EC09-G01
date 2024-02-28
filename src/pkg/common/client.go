package common

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// const Broker = "tcp://localhost:1891" // "broker.hivemq.com:1883"
// const Broker = "ssl://82cbaaf185b446e6a2cca175950e15ed.s1.eu.hivemq.cloud:8883/mqtt"
const Broker = "tcp://82cbaaf185b446e6a2cca175950e15ed.s1.eu.hivemq.cloud:1883/mqtt"
const IdPublisher = "go-mqtt-publisher"
const IdSubscriber = "go-mqtt-subscriber"
const Username = ""
const Password = ""

var Handler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received: %s on topic %s\n", msg.Payload(), msg.Topic())
	return
}

func CreateClient(broker string, id string, callback_handler mqtt.MessageHandler) mqtt.Client {

	opts := mqtt.NewClientOptions().AddBroker(broker)
	opts.SetClientID(id)
	opts.SetDefaultPublishHandler(callback_handler)
	opts.SetUsername(Username)
	opts.SetPassword(Password)

	return mqtt.NewClient(opts)
}
