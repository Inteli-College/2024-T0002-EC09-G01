package main

import (
	mongo "2024-T0002-EC09-G01/src/internal/mongo"
	consumerKafka "2024-T0002-EC09-G01/src/internal/kafka"
	"encoding/json"
	"fmt"
	ckafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
	godotenv "github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	client := mongo.ConnectToMongo("../../config/.env")

	err := godotenv.Load("../../config/.env")
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
		"group.id":           "grupo1",
		"auto.offset.reset":  "latest",
	}

	kafkaRepository := consumerKafka.NewKafkaConsumer(configMap, []string{"topic_queue"})

	go func() {
		if err := kafkaRepository.Consume(msgChan); err != nil {
			log.Printf("Error consuming kafka queue: %v", err)
		}
	}()

	fmt.Printf("Kafka consumer has been started\n")

	for msg := range msgChan {
		log.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))

		/*Make sure to send the data as a JSON so no errors occur while consuming. Example of acceptable payload:

		{"CarbonMonoxide": "carbon_monoxide", "NitrogenDioxide": "nitrogen_dioxide", "Ethanol": "ethanol", "Hydrogen": "hydrogen", "Ammonia": "ammonia", "Methane": "methane", "Propane": "propane", "IsoButane": "iso_butane"}
		
		*/

		var result map[string]interface{}
		err := json.Unmarshal(msg.Value, &result)
		if err != nil {
			log.Fatal(err)
		}

		mongo.InsertIntoMongo(client, result)

	}

}
