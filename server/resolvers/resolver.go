package resolvers

import (
	"database/sql"
	"fmt"
	"github/tokizuoh/faaaar/server/models"
	"io/ioutil"
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

func (dsn datasourceName) String() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", dsn.host, dsn.port, dsn.user, dsn.password, dsn.dbname, dsn.sslmode)
}

var db *sql.DB

type sqlcfg struct {
	base  string
	where string
}

func (s sqlcfg) Query() string {
	if s.where == "" {
		return s.base
	}
	return s.base + " " + "WHERE" + " " + s.where
}

func init() {
	dsn := datasourceName{
		host:     "faaaar-db",
		port:     5432,
		user:     "postgres",
		password: "postgres",
		dbname:   "postgres",
		sslmode:  "disable",
	}

	dsnString := dsn.String()
	_db, err := sql.Open("postgres", dsnString)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	db = _db
}

func readSQLFile(filepath string) (string, error) {
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", err
	}

	return string(b), nil
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
