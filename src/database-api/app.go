package main

import (
	"fmt"
	ckafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
	godotenv "github.com/joho/godotenv"
	"log"
	"os"
	consumerKafka "2024-T0002-EC09-G01/src/internal/kafka"
	"encoding/json"
)

func main() {

	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Printf("Error loading .env file: %s", err)
	}

	msgChan := make(chan *ckafka.Message)

	configMap := &ckafka.ConfigMap{
		"bootstrap.servers":  os.Getenv("CONFLUENT_BOOTSTRAP_SERVER_SASL"),
		"sasl.mechanisms":    "PLAIN",
		"security.protocol":  "SASL_SSL",
		"sasl.username":      os.Getenv("CONFLUENT_API_KEY"),
		"sasl.password":      os.Getenv("CONFLUENT_API_SECRET"),
		"session.timeout.ms": 6000,
		"group.id":           "manu",
		"auto.offset.reset":  "latest",
	}

	kafkaRepository := consumerKafka.NewKafkaConsumer(configMap, []string{"teste"})

	go func() {
		if err := kafkaRepository.Consume(msgChan); err != nil {
			log.Printf("Error consuming kafka queue: %v", err)
		}
	}()

	fmt.Printf("Kafka consumer has been started\n")

	for msg := range msgChan {
		log.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))

		var result map[string]interface{}
		err := json.Unmarshal(msg.Value, &result)
		if err != nil {
			log.Fatal(err)
		}

	}

}

