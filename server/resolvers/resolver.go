package resolvers

import (
	"database/sql"
	"github/tokizuoh/faaaar/server/types"
	"log"

	_ "github.com/lib/pq"
)

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

	dsnString := dsn.string()
	_db, err := sql.Open("postgres", dsnString)
	if err != nil {
		log.Fatal(err)
	}

	db = _db
}

func GetIdolsByAge(age int) ([]types.Idol, error) {
	idols, err := types.IdolsByAge(db, age)
	if err != nil {
		return nil, err
	}

	return idols, nil
}

func GetUnitsByIdolID(idolId int) ([]types.Unit, error) {
	units, err := types.UnitsByIdolID(db, idolId)
	if err != nil {
		return nil, err
	}

	return units, nil
}
