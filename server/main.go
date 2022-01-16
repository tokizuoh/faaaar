package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type datasourceName struct {
	host     string
	port     int
	user     string
	password string
	dbName   string
	sslMode  string
}

func getDataSourceNameString(dsn datasourceName) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", dsn.host, dsn.port, dsn.user, dsn.password, dsn.dbName, dsn.sslMode)
}

type idol struct {
	id          string
	name        string
	age         int
	height      int
	birth_place string
	birth_day   string
	blood_type  string
	unit        string
}

func main() {
	dsn := datasourceName{
		host:     "127.0.0.1",
		port:     5423,
		user:     "postgres",
		password: "postgres",
		dbName:   "postgres",
		sslMode:  "disable",
	}
	dsnString := getDataSourceNameString(dsn)
	db, err := sql.Open("postgres", dsnString)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT * FROM idol")
	if err != nil {
		log.Fatal(err)
	}

	var idols []idol
	for rows.Next() {
		var i idol
		rows.Scan(&i.id, &i.name, &i.age, &i.height, &i.birth_place, &i.birth_day, &i.blood_type, &i.unit)
		idols = append(idols, i)
	}
}
