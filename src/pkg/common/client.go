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

	broker := os.Getenv("BROKER_ADDR")
	username := os.Getenv("HIVE_USER")
	password := os.Getenv("HIVE_PSWD")

	if username == "" || password == "" {
		err := godotenv.Load("../.env")
		if err != nil {
			fmt.Printf("Error loading .env file. error: %s", err)
		}
		broker = os.Getenv("BROKER_ADDR")
		username = os.Getenv("HIVE_USER")
		password = os.Getenv("HIVE_PSWD")
	}

	opts := mqtt.NewClientOptions().AddBroker(fmt.Sprintf("tls://%s:%d", broker, port))
	opts.SetClientID(id)
	opts.SetDefaultPublishHandler(callback_handler)
	opts.SetUsername(username)
	opts.SetPassword(password)

	return mqtt.NewClient(opts)
}
