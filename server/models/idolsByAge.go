package models

import (
	"database/sql"
	"fmt"

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

func GetSameAgeIdols(o IdolsByAgeOption) ([]Idol, error) {
	// TODO: [#30] DB処理共通化 start
	dsn := datasourceName{
		host:     "faaaar-db",
		port:     5432,
		user:     "postgres",
		password: "postgres",
		dbname:   "postgres",
		sslmode:  "disable",
	}

	dsnString := getDataSourceNameString(dsn)
	db, err := sql.Open("postgres", dsnString)
	defer db.Close()

	if err != nil {
		return nil, err
	}
	// TODO: [#30] DB処理共通化 end

	stx, err := readSQLFile("./sqls/get_idols_by_age.sql")

	var where string
	if o.Age != 0 {
		where = fmt.Sprintf("age=%d", o.Age)
	}

	cfg := Sqlcfg{
		base:  stx,
		where: where,
	}

	rows, err := db.Query(cfg.Query())
	if err != nil {
		return nil, err
	}

	var result []Idol
	for rows.Next() {
		var i Idol
		rows.Scan(&i.Id, &i.Name, &i.Age, &i.Height, &i.Birthplace, &i.Birthday, &i.Bloodtype)
		result = append(result, i)
	}

	return result, nil
}
