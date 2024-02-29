package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

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
		CREATE TABLE IF NOT EXISTS gases
		(
			id INTEGER PRIMARY KEY,
			sensorId INTEGER,
			sensorName TEXT,
			unit VARCHAR(7),
			time DATETIME,
			NH3 FLOAT,        -- Ammonia
			CO FLOAT,         -- Carbon-Monoxide
			C2H5OH FLOAT,     -- Ethanol
			H2 FLOAT,         -- Hydrogen
			iC4H10 FLOAT,     -- Iso-butane
			CH4 FLOAT,        -- Methane
			NO2 FLOAT,        -- Nitrogen-Dioxide
			C3H8 FLOAT       -- Propane
		);
	`

	anotherQuery := `
		CREATE TABLE IF NOT EXISTS radiation
		(
			id INTEGER PRIMARY KEY,
			sensorId INTEGER,
			sensorName TEXT,
			unit VARCHAR(7),
			time DATETIME,
			radiation FLOAT
		);
	`

	oneMoreQuery := `
		CREATE TABLE IF NOT EXISTS sensors
		(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			latitude FLOAT,
			longitude FLOAT,
			sensor_name VARCHAR(100),
			Sensor_code VARCHAR(100),
			manufacturer VARCHAR(100),
			registered_by VARCHAR(100),
			registration_date DATETIME
		)
	`

	command, err := db.Prepare(sqlQuery)

	catchErrors(err)

	command2, err := db.Prepare(anotherQuery)

	catchErrors(err)

	command3, err := db.Prepare(oneMoreQuery)

	catchErrors(err)

	command.Exec()
	command2.Exec()
	command3.Exec()

	return db
}

func InsertIntoSensors(
	db *sql.DB,
	data map[string]interface{},
) {

	query := `
			INSERT INTO sensors
				(latitude, longitude, sensor_name, sensor_code, manufacturer, registered_by, registration_date) 
				VALUES (?, ?, ?, ?, ?, ?, ?);
		`

	statement, err := db.Prepare(query)

	catchErrors(err)

	_, err = statement.Exec(
		data["latitude"],
		data["longitude"],
		data["sensor"],
		data["code"],
		data["manufacturer"],
		data["author"],
		data["date"],
	)

	catchErrors(err)

	return
}

func insertIntoGases(
	db *sql.DB,
	data map[string]interface{},
) {
	query := `
		INSERT INTO gases
		(sensorId, sensorName, unit, time, NH3, CO, C2H5OH, H2, iC4H10, CH4, NO2, C3H8)
		VALUES (?,?,?,?,?,?,?,?,?,?,?,?)
	`
	statement, err := db.Prepare(query)

	catchErrors(err)

	_, err = statement.Exec(
		data["id"],
		data["sensor"],
		data["unit"],
		data["current_time"],
		data["ammonia"],
		data["carbon_monoxide"],
		data["ethanol"],
		data["hydrogen"],
		data["iso_butane"],
		data["methane"],
		data["nitrogen_dioxide"],
		data["propane"],
	)

	catchErrors(err)

	return

}

func insertIntoRadiation(
	db *sql.DB,
	data map[string]interface{},
) {
	query := `
		INSERT INTO radiation
		(sensorId, sensorName, unit, time, radiation)
		VALUES (?,?,?,?,?)
	`
	statement, err := db.Prepare(query)

	catchErrors(err)

	_, err = statement.Exec(
		data["id"],
		data["sensor"],
		data["unit"],
		data["current_time"],
		data["radiation"],
	)

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

	// fmt.Println(data["payload"])

	payloadData := data["payload"].(map[string]interface{})

	db, err := sql.Open("sqlite3", "./database.db")

	catchErrors(err)

	if payloadData["gases-values"] != nil {
		
		gasesValues := payloadData["gases-values"].(map[string]interface{})

		result := make(map[string]interface{})

		result["id"] = data["packet-id"]

		result["time"] = payloadData["current_time"]

		for key, value := range gasesValues["gases-values"].(map[string]interface{}) {
			result[key] = value
		}

		result["sensor"] = gasesValues["sensor"]
		result["unit"] = gasesValues["unit"]

		insertIntoGases(db, result)

	} else {
		
		radiationValues := payloadData["radiation-values"].(map[string]interface{})

		result := make(map[string]interface{})

		result["id"] = data["packet-id"]

		result["time"] = payloadData["current_time"]

		for key, value := range radiationValues["radiation-values"].(map[string]interface{}) {
			result[key] = value
		}

		result["sensor"] = radiationValues["sensor"]
		result["unit"] = radiationValues["unit"]

		fmt.Println(result)

		insertIntoRadiation(db, result)
	}

	return

}

func main() {

	sensorInteli := map[string]interface{}{
		"latitude": -23.555734690062174,
		"longitude": -46.73388952459531,
		"sensor": "Sensor-Inteli",
		"code": 123,
		"manufacturer": "SmarTopia",
		"author": "Luana Parra",
		"date": time.Now(),
	}
	
	sensorPaulista := map[string]interface{}{
		"latitude": -23.561472985154808,
		"longitude": -46.65594627611366,
		"sensor": "Sensor-Paulista",
		"code": 124,
		"manufacturer": "SmarTopia",
		"author": "Emanuele Martins",
		"date": time.Now(),
	}
	
	sensorFariaLima := map[string]interface{}{
		"latitude": -23.587039143730863,
		"longitude": -46.68163586869637,
		"sensor": "Sensor-Faria-Lima",
		"code": 125,
		"manufacturer": "SmarTopia",
		"author": "Felipe Leão",
		"date": time.Now(),
	}
	
	sensorShare := map[string]interface{}{
		"latitude": -23.572985847044286,
		"longitude": -46.706362987750225,
		"sensor": "Sensor-Share",
		"code": 126,
		"manufacturer": "SmarTopia",
		"author": "Felipe Campos",
		"date" : time.Now(),
	}

	db, err := sql.Open("sqlite3", "./database.db")	

	catchErrors(err)

	
	ConnectToDatabase("sqlite3", "./database.db")
	
	InsertIntoSensors(db, sensorInteli)
	InsertIntoSensors(db, sensorPaulista)
	InsertIntoSensors(db, sensorFariaLima)
	InsertIntoSensors(db, sensorShare)

	client := DefaultClient.CreateClient(DefaultClient.Broker, DefaultClient.IdSubscriber, DatabaseHandler)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := client.Subscribe("sensor/#", 1, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		return
	}

	fmt.Println("Subscriber running...")
	select {}
}
