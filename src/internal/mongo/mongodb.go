package mongo

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Sensors struct {
	ID string `json:"_id"`
	// Latitude   float32 `json:"latitude"`
	// Longitude  float32 `json:"longitude"`
	Name       string `json:"name"`
	SensorType string `json:"sensorType"`
}

func InsertIntoMongo(client *mongo.Client, data map[string]interface{}) {
	db := client.Database("SmarTopia")

	var coll *mongo.Collection
	// fmt.Println(data["payload"])

	payloadData := data["payload"].(map[string]interface{})

	newData := make(map[string]interface{})

	if payloadData["gases-values"] != nil {

		gasesValues := payloadData["gases-values"].(map[string]interface{})

		newData["id"] = data["packet-id"]

		newData["time"] = payloadData["current_time"]

		for key, value := range gasesValues["gases-values"].(map[string]interface{}) {
			newData[key] = value
		}

		newData["sensor"] = gasesValues["sensor"]
		newData["unit"] = gasesValues["unit"]

		coll = db.Collection("gases")

	} else {

		radiationValues := payloadData["radiation-values"].(map[string]interface{})

		newData["id"] = data["packet-id"]

		newData["time"] = payloadData["current_time"]

		for key, value := range radiationValues["radiation-values"].(map[string]interface{}) {
			newData[key] = value
		}

		newData["sensor"] = radiationValues["sensor"]
		newData["unit"] = radiationValues["unit"]

		coll = db.Collection("radiation")
	}

	bsonData, err := bson.Marshal(newData)

	result, err := coll.InsertOne(context.TODO(), bsonData)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
}

func ConnectToMongo(path string) *mongo.Client {
	// Carregar variáveis de ambiente do arquivo .env
	err := godotenv.Load(path)

	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	// Recuperar usuário e senha do arquivo .env
	mongoUser := os.Getenv("MONGO_USER")
	mongoPassword := os.Getenv("MONGO_PASSWORD")

	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@sensors.zyzjabc.mongodb.net/?retryWrites=true&w=majority&appName=sensors", mongoUser, mongoPassword)).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	return client
}

func GetAllSensors(path string) ([]Sensors, error) {

	client := ConnectToMongo(path)
	collection := client.Database("SmarTopia").Collection("sensors")

	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var sensors []Sensors
	for cursor.Next(context.TODO()) {

		var doc bson.M
		err := cursor.Decode(&doc)
		if err != nil {
			log.Fatal(err)
		}

		jsonData, err := json.MarshalIndent(doc, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		var sensor Sensors
		err = json.Unmarshal(jsonData, &sensor)
		if err != nil {
			log.Fatal(err)
		}

		sensors = append(sensors, sensor)
	}

	fmt.Print(sensors)
	return sensors, nil
}
