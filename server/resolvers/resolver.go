package resolvers

import (
	"database/sql"
	"github/tokizuoh/faaaar/server/types"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	dsn := datasourceName{
		user:     "root",
		password: "root",
		protocol: "tcp(faaaar-db:3306)",
		dbname:   "shiny_colors_db",
	}

	dsnString := dsn.string()
	_db, err := sql.Open("mysql", dsnString)
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

func GetUnits() ([]types.Unit, error) {
	units, err := types.Units(db)
	if err != nil {
		return nil, err
	}

	return units, nil
}
