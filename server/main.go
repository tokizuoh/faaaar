package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
	_ "github.com/lib/pq"
)

type datasourceName struct {
	host     string
	port     int
	user     string
	password string
	dbName   string
	sslMode  string
}

func getDataSourceNameString(dsn datasourceName) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", dsn.host, dsn.port, dsn.user, dsn.password, dsn.dbName, dsn.sslMode)
}

type idol struct {
	id          string
	name        string
	age         int
	height      int
	birth_place string
	birth_day   string
	blood_type  string
	unit        string
}

func main() {
	dsn := datasourceName{
		host:     "172.26.0.1", // TODO: composeで固定化する
		port:     5423,
		user:     "postgres",
		password: "postgres",
		dbName:   "postgres",
		sslMode:  "disable",
	}
	dsnString := getDataSourceNameString(dsn)
	db, err := sql.Open("postgres", dsnString)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT * FROM idol")
	if err != nil {
		log.Fatal(err)
	}

	var idols []idol
	for rows.Next() {
		var i idol
		rows.Scan(&i.id, &i.name, &i.age, &i.height, &i.birth_place, &i.birth_day, &i.blood_type, &i.unit)
		idols = append(idols, i)
	}

	fields := graphql.Fields{
		"age": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "hoge", nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{
		Name:   "RootQuery",
		Fields: fields,
	}

	schemeConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(rootQuery),
	}

	scheme, err := graphql.NewSchema(schemeConfig)
	if err != nil {
		log.Fatal(err)
	}

	query := `
		{
			age
		}
	`

	params := graphql.Params{Schema: scheme, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatal(r.Errors)
	}

	rJSON, _ := json.Marshal(r)
	log.Printf("%s \n", rJSON) // {"data":{"age":"hoge"}}
}
