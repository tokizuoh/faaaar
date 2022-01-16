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

	// TODO: PostgreSQLからデータを取得する処理を追加する
}
