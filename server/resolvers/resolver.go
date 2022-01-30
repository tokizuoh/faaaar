package resolvers

import (
	"database/sql"
	"github/tokizuoh/faaaar/server/models"
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

func GetIdolsByAge(age int) ([]models.Idol, error) {
	idols, err := models.IdolsByAge(db, age)
	if err != nil {
		return nil, err
	}

	return idols, nil
}

func GetUnitsByIdolID(idolId int) ([]models.Unit, error) {
	units, err := models.UnitsByIdolID(db, idolId)
	if err != nil {
		return nil, err
	}

	return units, nil
}
