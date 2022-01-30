package resolver

import (
	"database/sql"
	"fmt"
	"log"
)

type Resolver struct{}

type datasourceName struct {
	host     string
	port     int
	user     string
	password string
	dbname   string
	sslmode  string
}

func getDataSourceNameString(dsn datasourceName) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", dsn.host, dsn.port, dsn.user, dsn.password, dsn.dbname, dsn.sslmode)
}

var db *sql.DB

func init() {
	dsn := datasourceName{
		host:     "faaaar-db",
		port:     5432,
		user:     "postgres",
		password: "postgres",
		dbname:   "postgres",
		sslmode:  "disable",
	}

	dsnString := getDataSourceNameString(dsn)
	_db, err := sql.Open("postgres", dsnString)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	db = _db
}
