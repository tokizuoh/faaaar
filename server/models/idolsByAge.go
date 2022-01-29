package models

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
)

type Idol struct {
	Id         string
	Name       string
	Age        int
	Height     int
	Birthplace string
	Birthday   string
	Bloodtype  string
	Unit       string
}

type IdolsByAgeOption struct {
	Age int
}

var IdolType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Idol",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				idol := p.Source.(Idol)
				return idol.Id, nil
			},
		},
		"name": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				idol := p.Source.(Idol)
				return idol.Name, nil
			},
		},
		"age": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				idol := p.Source.(Idol)
				return idol.Age, nil
			},
		},
		"height": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				idol := p.Source.(Idol)
				return idol.Height, nil
			},
		},
		"birth_place": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				idol := p.Source.(Idol)
				return idol.Birthplace, nil
			},
		},
		"birth_day": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				idol := p.Source.(Idol)
				return idol.Birthday, nil
			},
		},
		"blood_type": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				idol := p.Source.(Idol)
				return idol.Bloodtype, nil
			},
		},
	},
})

func GetSameAgeIdols(db *sql.DB, o IdolsByAgeOption) []Idol {
	var stx string
	if o.Age == 0 {
		// TODO: [#25] SQL文のリテラルをSQLファイル読み込みに変更する
		stx = "select * from idol order by id"
	} else {
		// TODO: [#25] SQL文のリテラルをSQLファイル読み込みに変更する
		stx = fmt.Sprintf("select * from idol where age=%d order by id", o.Age)
	}
	rows, err := db.Query(stx)
	if err != nil {
		log.Fatal(err)
	}

	var result []Idol
	for rows.Next() {
		var i Idol
		rows.Scan(&i.Id, &i.Name, &i.Age, &i.Height, &i.Birthplace, &i.Birthday, &i.Bloodtype)
		if o.Age == 0 || i.Age == o.Age {
			result = append(result, i)
		}

	}

	return result
}
