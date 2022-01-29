package models

import (
	"database/sql"
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

func GetUnitsByIdolID(db *sql.DB, o UnitsByIdolIdOption) []Unit {
	stx := "select * from unit"

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

	return result
}