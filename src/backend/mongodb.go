package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetonMongo(client *mongo.Client) []Sensor {
	db := client.Database("SmarTopia")
	
	coll := db.Collection("sensors")

	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	var sensors []Sensor
	if err = cursor.All(context.TODO(), &sensors); err != nil {
		log.Fatal(err)
	}

	return sensors
}

func InsertIntoMongo(client *mongo.Client, data interface{}, collection string) string {
	db := client.Database("SmarTopia")
	
	coll := db.Collection(collection)

	bsonData, err := bson.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	result, err := coll.InsertOne(context.TODO(), bsonData)
	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("Inserted document with _id: %v\n", result.InsertedID)
}

func ConnectToMongo() *mongo.Client{
	// Carregar variáveis de ambiente do arquivo .env
	err := godotenv.Load("../config/.env")

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