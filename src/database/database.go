package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	DefaultClient "2024-T0002-EC09-G01/src/pkg/common"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	_ "github.com/mattn/go-sqlite3"
)

func ConnectToDatabase(database string, source string) *sql.DB {
	db, err := sql.Open(database, source)

	catchErrors(err)

	return CreateTable(db)
}

func CreateTable(db *sql.DB) *sql.DB {

	sqlQuery := `
		CREATE TABLE IF NOT EXISTS measurements
		(
			id INTEGER PRIMARY KEY,
			sensorId INTEGER,
			latitude FLOAT,
			longitude FLOAT,
			time DATE,
			gasesData TEXT,
			radiationData TEXT
		);
	`

	command, err := db.Prepare(sqlQuery)

	catchErrors(err)

	command.Exec()

	return db
}

func InsertData(
	db *sql.DB,
	data map[string]interface{},
) {

	gasesData, err := json.Marshal(data["gases-values"])

	catchErrors(err)

	radiationData, err := json.Marshal(data["radiation-values"])

	catchErrors(err)

	query := `
			INSERT INTO measurements 
				(sensorId, latitude, longitude, time, gasesData, radiationData) 
				VALUES (?, ?, ?, ?, ?, ?);
		`

	statement, err := db.Prepare(query)

	catchErrors(err)

	_, err = statement.Exec(
		data["identifier"],
		data["latitude"],
		data["longitude"],
		data["current_time"],
		string(gasesData),
		string(radiationData))

	catchErrors(err)

	return
}

func catchErrors(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}

var DatabaseHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	// fmt.Printf("Received: %s on topic %s\n", msg.Payload(), msg.Topic())

	var data map[string]interface{}

	err := json.Unmarshal([]byte(string(msg.Payload())), &data)

	catchErrors(err)

	db, err := sql.Open("sqlite3", "./database.db")

	catchErrors(err)

	InsertData(db, data)

}

func main() {

	ConnectToDatabase("sqlite3", "./database.db")

	client := DefaultClient.CreateClient(DefaultClient.Broker, DefaultClient.IdSubscriber, DatabaseHandler)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := client.Subscribe("sensors", 1, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		return
	}

	fmt.Println("Subscriber running...")
	select {}
}
