package mongo

import (
	"context"
	"fmt"
	"log"
	"os"
	"encoding/json"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Sensors struct {
	ID        string `json:"_id"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	Name      string `json:"name"`
	SensorTyoe string `json:"sensorType"`
}

func InsertIntoMongo(data map[string]interface{}) {
	client := ConnectToMongo()
	db := client.Database("SmarTopia")
	coll := db.Collection("teste-gases")

	for key, value := range data {
		// Codifique o valor em JSON
		encodedJSON, err := json.Marshal(value)
		if err != nil {
			fmt.Printf("Erro ao codificar JSON para %s: %s\n", key, err)
			continue
		}

		// Decodifique o JSON em um mapa vazio
		var doc map[string]interface{}
		if err := json.Unmarshal(encodedJSON, &doc); err != nil {
			fmt.Printf("Erro ao decodificar JSON para %s: %s\n", key, err)
			continue
		}

		// Insira o documento no MongoDB
		if _, err := coll.InsertOne(context.TODO(), doc); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Documento inserido com sucesso para %s\n", key)
	}
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

func GetAllSensors() ([]Sensors, error) {

	client := ConnectToMongo()
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
