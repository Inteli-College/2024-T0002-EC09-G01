package mongo

import (
	"context"
	"fmt"
	"log"
	"os"
	// "encoding/json"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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

func ConnectToMongo() *mongo.Client{
	// Carregar variáveis de ambiente do arquivo .env
	err := godotenv.Load("../../config/.env")

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
