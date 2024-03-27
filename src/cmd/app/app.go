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
	"time"
	//"go.mongodb.org/mongo-driver/bson"
	//mics6814 "2024-T0002-EC09-G01/src/internal/sensors/mics6814"
)

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func main() {

	//client := mongo.ConnectToMongo()

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

	ticker := time.NewTicker(1 * time.Minute)
	startTime := time.Now()
		
	var sensorDataGases = make(map[string][]map[string]interface{})
	var sensorDataRadiation = make(map[string][]map[string]interface{})


	for msg := range msgChan {
		//log.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))

		select {
			case <-ticker.C:
				newData := make(map[string]interface{})

				for key, value := range sensorDataGases {
					newData[key] = map[string]interface{}{
						"sensor-id": key,
						"star-time": startTime,
						"end-time":  time.Now(),
						"gases-values": value,
					}
				}
				
				ticker.Stop()
				ticker = time.NewTicker(1 * time.Minute)

				fmt.Print("1 minutes have passed\n")
				//fmt.Println(string(bsonData))
				
				mongo.InsertIntoMongo(newData)
				
				

			default:
				var result map[string]interface{}
				err := json.Unmarshal(msg.Value, &result)
				if err != nil {
					log.Fatal(err)
				}
								
				packetID := fmt.Sprintf("%s", result["packet-id"])

				if sensorDataGases[packetID] == nil {
					sensorDataGases[packetID] = make([]map[string]interface{}, 0)
				}		
				if sensorDataRadiation[packetID] == nil {
					sensorDataRadiation[packetID] = make([]map[string]interface{}, 0)
				}		

				payloadData := result["payload"].(map[string]interface{})

				if gasesValues, ok := payloadData["gases-values"].(map[string]interface{}); ok {
					sensorList := sensorDataGases[packetID] 
					gasesValues["current_time"] = payloadData["current_time"]
					sensorList = append(sensorList, gasesValues)
					sensorDataGases[packetID] = sensorList
				}

				// if radiationValues, ok := payloadData["radiation-values"].(map[string]interface{}); ok {
				// 	sensorList := sensorDataRadiation[packetID] 
				// 	sensorList = append(sensorList, radiationValues)
				// 	sensorDataRadiation[packetID] = sensorList
				// }

				//fmt.Println(sensorDataGases)


		}



	}

}
