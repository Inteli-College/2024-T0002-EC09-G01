package common

import (
	"fmt"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	godotenv "github.com/joho/godotenv"
)

const IdPublisher = "go-mqtt-publisher"
const IdSubscriber = "go-mqtt-subscriber"

const port = 8883

var Handler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received: %s on topic %s\n", msg.Payload(), msg.Topic())
}

func CreateClient(id string, callback_handler mqtt.MessageHandler) mqtt.Client {

	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Printf("Error loading .env file: %s", err)
	}

	opts := mqtt.NewClientOptions().AddBroker(fmt.Sprintf("tls://%s:%d", os.Getenv("BROKER_ADDR"), port))
	opts.SetClientID(id)
	opts.SetDefaultPublishHandler(callback_handler)
	opts.SetUsername(os.Getenv("HIVE_USER"))
	opts.SetPassword(os.Getenv("HIVE_PSWD"))

	return mqtt.NewClient(opts)
}
