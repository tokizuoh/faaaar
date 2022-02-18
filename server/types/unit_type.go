package types

import (
	"database/sql"
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
	idolId   int
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

func Units(db *sql.DB, idolId int) ([]Unit, error) {
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
		rows.Scan(&gur.unitName, &gur.idolName, &gur.idolId)

		response = append(response, gur)
	}

	// "unit_name": "idol_id"
	mi := map[string][]int{}
	for _, r := range response {
		mi[r.unitName] = append(mi[r.unitName], r.idolId)
	}

	f := func(targetId int, idolIds []int) bool {
		for _, id := range idolIds {
			if id == targetId {
				return true
			}
		}

		return false
	}

	// "unit_name": "idol_name"
	mn := map[string][]string{}
	for _, r := range response {
		if idolId == 0 || f(idolId, mi[r.unitName]) {
			mn[r.unitName] = append(mn[r.unitName], r.idolName)
		}
	}

	var result []Unit
	for key, value := range mn {
		u := Unit{
			Name:  key,
			Idols: value,
		}
		result = append(result, u)
	}

	return result, nil
}
