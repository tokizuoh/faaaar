package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5423 user=postgres password=postgres dbname=postgres")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	// TODO: PostgreSQLからデータを取得する処理を追加する
}
