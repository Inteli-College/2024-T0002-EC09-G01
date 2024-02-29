package main

import (
	"database/sql"
	"log"
	"math/rand"

	_ "github.com/mattn/go-sqlite3"
)

func CreateDatabase(database string, source string) *sql.DB{
	db, _ := sql.Open(database, source)

	sqlQuery := `
		CREATE TABLE IF NOT EXISTS measurements
		(
			id INTEGER PRIMARY KEY,
			name VARCHAR(15),
			latitude FLOAT,
			longitude FLOAT,
			measurement FLOAT,
			rate FLOAT,
			unit VARCHAR(15)
		);
	`

	command, err := db.Prepare(sqlQuery)

	catchErrors(err)

	command.Exec()

	return db
}

func catchErrors(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func main() {
	database := CreateDatabase("sqlite3", "./database.db")

	insertData := func(db *sql.DB, name string, measurement float32) {
		query := `INSERT INTO measurements(name, measurement) VALUES (?, ?);`

		statement, err := db.Prepare(query)

		catchErrors(err)

		_, err = statement.Exec(name, measurement)

		catchErrors(err)
	}

	insertData(database, "SPS30", rand.Float32())

	displayUsers(database)

}

func displayUsers(db *sql.DB) {
	row, err := db.Query("SELECT * FROM measurements ORDER BY name")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() {
		var id int
		var name string
		var measurement float32
		row.Scan(&id, &name, &measurement)
		log.Println("Sensor: ", id, " ", name, " ", measurement)
	}
}
