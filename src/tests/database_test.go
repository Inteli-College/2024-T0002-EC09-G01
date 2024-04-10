package testing

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestInsertIntoMongo(t *testing.T) {
	client := setupMongoClient(t)
	defer client.Disconnect(context.Background())

	// data := map[string]interface{}{
	// 	"packet-id": "123456789",
	// 	"payload": map[string]interface{}{
	// 		"gases-values": map[string]interface{}{
	// 			"CO2":   100,
	// 			"Oxygen": 20,
	// 			"sensor": "XYZ123",
	// 			"unit":   "ppm",
	// 		},
	// 		"current_time": "2024-03-28T12:00:00Z",
	// 	},
	// }

	// err := mongo.InsertIntoMongo(client, data)
	// if err != nil {
	// 	t.Errorf("Erro ao inserir no MongoDB: %v", err)
	// }
}

func setupMongoClient(t *testing.T) *mongo.Client {
	t.Helper()

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		t.Fatalf("Erro ao conectar ao mongoDB para testes: %v", err)
	}

	return client
}
