package models

import (
	"database/sql"
	"io/ioutil"
	"log"

	"github.com/graphql-go/graphql"
)

type Unit struct {
	Id   string
	Name string
}

type UnitsByIdolIdOption struct {
	IdolId int
}

var UnitType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Unit",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				unit := p.Source.(Unit)
				return unit.Id, nil
			},
		},
		"name": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				unit := p.Source.(Unit)
				return unit.Name, nil
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

func GetUnitsByIdolID(db *sql.DB, o UnitsByIdolIdOption) ([]Unit, error) {
	stx, err := readSQLFile("./sqls/get_units_by_idol_id.sql")
	if err != nil {
		return nil, err
	}

	if o.IdolId == 0 {
		// TODO: [#25] SQL文のリテラルをSQLファイル読み込みに変更する
		// stx = "SELECT idl.id          AS id, idl.name        AS name, idl.age         AS age, idl.height      AS height, idl.birth_place AS birth_place, idl.blood_type  AS blood_type, unt.name        AS unit FROM idol idl INNER JOIN idol_unit idlunt ON idl.id = idlunt.idol INNER JOIN unit unt ON idlunt.unit = unt.id"
	} else {
		// TODO: [#25] SQL文のリテラルをSQLファイル読み込みに変更する
		// whereでidの指定する
		// stx = fmt.Sprintf("SELECT idl.id          AS id, unt.name        AS unit FROM idol idl INNER JOIN idol_unit idlunt ON idl.id = idlunt.idol INNER JOIN unit unt ON idlunt.unit = unt.id WHERE idl.id=%d", o.IdolId)
	}

	rows, err := db.Query(stx)
	if err != nil {
		log.Fatal(err)
	}

	var result []Unit
	for rows.Next() {
		var u Unit
		rows.Scan(&u.Id, &u.Name)

		result = append(result, u)
	}

	return result, nil
}
