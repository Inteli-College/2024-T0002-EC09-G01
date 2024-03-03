package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
	"log"
)

func ConnectToDatabase(source string, database string) *sql.DB {
	db, err := sql.Open(source, database)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func InsertIntoSensors(data Sensor) string {
	db := ConnectToDatabase("sqlite3", "database.db")

	query := `
			INSERT INTO sensors
				(latitude, longitude, sensor_name, sensor_code, manufacturer, registered_by, registration_date) 
				VALUES (?, ?, ?, ?, ?, ?, ?);
		`
	statement, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	_, err = statement.Exec(
		data.Latitude,
		data.Longitude,
		data.Sensor,
		data.Code,
		data.Manufacturer,
		data.Author,
		data.Date,
	)

	if err != nil {
		log.Fatal(err)
	}

	response := fmt.Sprintf("Sensor: %s added successfully", data.Sensor)

	return response
}

type SQLSensor struct {
	Id int `json:"id"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Sensor string `json:"sensor"`
	Code int `json:"code"`
	Manufacturer string `json:"manufacturer"`
	Author string `json:"author"`
	Date string `json:"date"`
}

func GetInSensors() []SQLSensor {
	db := ConnectToDatabase("sqlite3", "database.db")

	query := `
		SELECT * FROM sensors;
	`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	var sensors []SQLSensor
	for rows.Next() {
		var sensor SQLSensor
		err := rows.Scan(
			&sensor.Id,
			&sensor.Latitude,
			&sensor.Longitude,
			&sensor.Sensor,
			&sensor.Code,
			&sensor.Manufacturer,
			&sensor.Author,
			&sensor.Date,
		)
		if err != nil {
			log.Fatal(err)
		}
		sensors = append(sensors, sensor)
	}

	return sensors
}

func insertIntoGases(data Gas) string {
	db := ConnectToDatabase("sqlite3", "database.db")

	query := `
		INSERT INTO gases
		(sensorId, sensorName, unit, time, NH3, CO, C2H5OH, H2, iC4H10, CH4, NO2, C3H8)
		VALUES (?,?,?,?,?,?,?,?,?,?,?,?)
	`
	statement, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	_, err = statement.Exec(
		data.SensorId,
		data.SensorName,
		data.Unit,
		data.Time,
		data.NH3,
		data.CO,
		data.C2H5OH,
		data.H2,
		data.IC4H10,
		data.CH4,
		data.NO2,
		data.C3H8,
	)

	if err != nil {
		log.Fatal(err)
	}

	response := fmt.Sprintf("Gas: %s added successfully", data.SensorName)

	return response
}

type SQLGas struct {
	Id int `json:"id"`
	SensorId int `json:"sensorId"`
	SensorName string `json:"sensorName"`
	Unit string `json:"unit"`
	Time string `json:"time"`
	NH3 float64 `json:"NH3"`
	CO float64 `json:"CO"`
	C2H5OH float64 `json:"C2H5OH"`
	H2 float64 `json:"H2"`
	IC4H10 float64 `json:"IC4H10"`
	CH4 float64 `json:"CH4"`
	NO2 float64 `json:"NO2"`
	C3H8 float64 `json:"C3H8"`
}

func GetInGases() []SQLGas {
	db := ConnectToDatabase("sqlite3", "database.db")

	query := `
		SELECT * FROM gases;
	`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	var gases []SQLGas
	for rows.Next() {
		var gas SQLGas
		err := rows.Scan(
			&gas.Id,
			&gas.SensorId,
			&gas.SensorName,
			&gas.Unit,
			&gas.Time,
			&gas.NH3,
			&gas.CO,
			&gas.C2H5OH,
			&gas.H2,
			&gas.IC4H10,
			&gas.CH4,
			&gas.NO2,
			&gas.C3H8,
		)
		if err != nil {
			log.Fatal(err)
		}
		gases = append(gases, gas)
	}

	return gases
}

func insertIntoRadiation(data Radiation) string {
	db := ConnectToDatabase("sqlite3", "database.db")

	query := `
		INSERT INTO radiation
		(sensorId, sensorName, unit, time, radiation)
		VALUES (?,?,?,?,?)
	`
	statement, err := db.Prepare(query)

	if err != nil {
		log.Fatal(err)
	}

	_, err = statement.Exec(
		data.SensorId,
		data.SensorName,
		data.Unit,
		data.Time,
		data.Radiation,
	)

	if err != nil {
		log.Fatal(err)
	}

	response := fmt.Sprintf("Radiation: %s added successfully", data.SensorName)

	return response
}

type SQLRadiation struct {
	Id int `json:"id"`
	SensorId int `json:"sensorId"`
	SensorName string `json:"sensorName"`
	Unit string `json:"unit"`
	Time string `json:"time"`
	Radiation float64 `json:"radiation"`
}

func GetInRadiation() []SQLRadiation {
	db := ConnectToDatabase("sqlite3", "database.db")

	query := `
		SELECT * FROM radiation;
	`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	var radiation []SQLRadiation
	for rows.Next() {
		var rad SQLRadiation
		err := rows.Scan(
			&rad.Id,
			&rad.SensorId,
			&rad.SensorName,
			&rad.Unit,
			&rad.Time,
			&rad.Radiation,
		)
		if err != nil {
			log.Fatal(err)
		}
		radiation = append(radiation, rad)
	}

	return radiation
}

func CreateTable() {
	db := ConnectToDatabase("sqlite3", "database.db")

	sqlQuery := `
		CREATE TABLE IF NOT EXISTS gases
		(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
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
			id INTEGER PRIMARY KEY AUTOINCREMENT,
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
		);
	`

	command, err := db.Prepare(sqlQuery)
	if err != nil {
		log.Fatal(err)
	}

	command2, err := db.Prepare(anotherQuery)
	if err != nil {
		log.Fatal(err)
	}

	command3, err := db.Prepare(oneMoreQuery)
	if err != nil {
		log.Fatal(err)
	}

	command.Exec()
	command2.Exec()
	command3.Exec()
}
