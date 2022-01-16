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

type Idol struct {
	Id          string
	Name        string
	Age         int
	Height      int
	Birth_place string
	Birth_day   string
	Blood_type  string
	Unit        string
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
	},
})

func getSameAgeIdols(db *sql.DB, age int) []Idol {
	st := fmt.Sprintf("select * from idol where age=%d", age)
	rows, err := db.Query(st)
	if err != nil {
		log.Fatal(err)
	}

	var result []Idol
	for rows.Next() {
		var i Idol
		rows.Scan(&i.Id, &i.Name, &i.Age, &i.Height, &i.Birth_place, &i.Birth_day, &i.Blood_type, &i.Unit)

		if i.Age == age {
			result = append(result, i)
		}

	}

	return result
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

	scheme, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"idols": &graphql.Field{
					Type: graphql.NewList(IdolType),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						ageQuery, ok := p.Args["age"].(int)
						if ok {
							result := getSameAgeIdols(db, ageQuery)
							return result, nil
						}
						return nil, nil
					},
					Args: graphql.FieldConfigArgument{
						"age": &graphql.ArgumentConfig{
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

	query := `
		{
			idols(age: 20) {
				id
				name
			}
		}
	`

	params := graphql.Params{
		Schema:        scheme,
		RequestString: query,
	}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatal(r.Errors)
	}

	rJSON, _ := json.Marshal(r)
	log.Printf("%s \n", rJSON) // {"data":{"idols":[{"id":16,"name":"有栖川 夏葉"},{"id":26,"name":"斑鳩 ルカ"}]}}
}
