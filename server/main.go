package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github/tokizuoh/faaaar/server/models"
	"io/ioutil"
	"log"

	"github.com/graphql-go/graphql"
	_ "github.com/lib/pq"
)

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

func readQuery(filepath string) (string, error) {
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

const QUERY_FILE_PATH = "./query.txt"

func main() {
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
		log.Fatal(err)
	}

	scheme, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"idols": &graphql.Field{
					Type: graphql.NewList(models.IdolType),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						ageQuery, ok := p.Args["age"].(int)
						if ok {
							result := models.GetSameAgeIdols(db, models.IdolsByAgeOption{Age: ageQuery})
							return result, nil
						} else {
							result := models.GetSameAgeIdols(db, models.IdolsByAgeOption{})
							return result, nil
						}
					},
					Args: graphql.FieldConfigArgument{
						"age": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
					},
				},
				"units": &graphql.Field{
					Type: graphql.NewList(models.UnitType),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						idolIdQuery, ok := p.Args["idolId"].(int)
						if ok {
							result := models.GetUnitsByIdolID(db, models.UnitsByIdolIdOption{IdolId: idolIdQuery})
							return result, nil
						} else {
							result := models.GetUnitsByIdolID(db, models.UnitsByIdolIdOption{})
							return result, nil
						}
					},
					Args: graphql.FieldConfigArgument{
						"idolId": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
					},
				},
			},
		}),
	})

	if err != nil {
		log.Fatal(err)
	}

	query, err := readQuery(QUERY_FILE_PATH)
	if err != nil {
		log.Fatal(err)
	}

	params := graphql.Params{
		Schema:        scheme,
		RequestString: query,
	}

	r := graphql.Do(params)
	if r.HasErrors() {
		log.Fatal(r.Errors)
	}

	output, err := json.MarshalIndent(r, "", "\t")
	log.Printf("%s \n", output)
}
