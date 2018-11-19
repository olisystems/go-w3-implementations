package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3" // SQL Lite Driver
)

type Database struct {
	db *sql.DB
}

func Connect() *Database {

	db, err := sql.Open("sqlite3", "./meter.db")
	if err != nil {
		log.Fatal(err)
	}

	return &Database{db: db}
}

func (database Database) Disconnect() {
	database.db.Close()
}

func (database Database) WriteToMeter(timestamp int, consumption int, production int) sql.Result {
	db := database.db

	stmt, err := db.Prepare("insert into meter(timestamp, consumption, production) values(?,?,?)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(timestamp, consumption, production)
	if err != nil {
		log.Fatal(err)
	}

	return res

}

type Meter struct {
	Timestamp   int
	Consumption int
	Production  int
}

func (meter Meter) ToString() string {
	return fmt.Sprintf("%t %c %p", meter.Timestamp, meter.Consumption, meter.Production)
}

// Queries for every column in the MeterTable
// QueryAppend Filter fore query e.g "ORDER BY timestamp DESC LIMIT 1"
func (database Database) ReadFromMeter(queryAppend string) []Meter {
	db := database.db
	query := "select * from meter " + queryAppend

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	output := make([]Meter, 0)

	var timestamp int
	var consumption int
	var production int

	for rows.Next() {
		err = rows.Scan(&timestamp, &consumption, &production)
		if err != nil {
			log.Fatal(err)
		}
		meter := Meter{
			Timestamp:   timestamp,
			Consumption: consumption,
			Production:  production,
		}

		output = append(output, meter)
	}

	rows.Close()

	return output
}
