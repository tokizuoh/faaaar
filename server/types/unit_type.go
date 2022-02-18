package types

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/graphql-go/graphql"
)

type Unit struct {
	Name  string
	Idols []string
}

type GetUnitsReponse struct {
	unitName string
	idolName string
}

var UnitType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Unit",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				unit := p.Source.(Unit)
				return unit.Name, nil
			},
		},
		"idols": &graphql.Field{
			Type: graphql.NewList(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				unit := p.Source.(Unit)
				return unit.Idols, nil
			},
		},
	},
})

func readSQLFile(filepath string) (string, error) {
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func UnitsByIdolID(db *sql.DB, idolId int) ([]Unit, error) {
	stx, err := readSQLFile("./sqls/get_units_by_idol_id.sql")
	if err != nil {
		return nil, err
	}

	var where string
	if idolId != 0 {
		where = fmt.Sprintf("idl.id=%d", idolId)
	}

	cfg := Sqlcfg{
		base:  stx,
		where: where,
	}

	rows, err := db.Query(cfg.Query())
	if err != nil {
		log.Fatal(err)
	}

	var result []Unit
	for rows.Next() {
		var u Unit
		// TODO: [#43]
		// rows.Scan(&u.Id, &u.Name)

		result = append(result, u)
	}

	return result, nil
}

// TODO: [#43]
func Units(db *sql.DB) ([]Unit, error) {
	stx, err := readSQLFile("./sqls/get_units.sql")
	if err != nil {
		return nil, err
	}

	cfg := Sqlcfg{
		base: stx,
	}

	rows, err := db.Query(cfg.Query())
	if err != nil {
		log.Fatal(err)
	}

	var response []GetUnitsReponse
	for rows.Next() {
		var gur GetUnitsReponse
		rows.Scan(&gur.unitName, &gur.idolName)

		response = append(response, gur)
	}

	m := map[string][]string{}
	for _, r := range response {
		m[r.unitName] = append(m[r.unitName], r.idolName)
	}

	var result []Unit
	for key, value := range m {
		u := Unit{
			Name:  key,
			Idols: value,
		}
		result = append(result, u)
	}

	return result, nil
}
