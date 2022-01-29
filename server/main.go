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

// type Idol struct {
// 	Id         string
// 	Name       string
// 	Age        int
// 	Height     int
// 	Birthplace string
// 	Birthday   string
// 	Bloodtype  string
// 	Unit       string
// }

// type Option struct {
// 	age int
// }

// var IdolType = graphql.NewObject(graphql.ObjectConfig{
// 	Name: "Idol",
// 	Fields: graphql.Fields{
// 		"id": &graphql.Field{
// 			Type: graphql.Int,
// 			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 				idol := p.Source.(Idol)
// 				return idol.Id, nil
// 			},
// 		},
// 		"name": &graphql.Field{
// 			Type: graphql.String,
// 			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 				idol := p.Source.(Idol)
// 				return idol.Name, nil
// 			},
// 		},
// 		"age": &graphql.Field{
// 			Type: graphql.Int,
// 			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 				idol := p.Source.(Idol)
// 				return idol.Age, nil
// 			},
// 		},
// 		"height": &graphql.Field{
// 			Type: graphql.Int,
// 			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 				idol := p.Source.(Idol)
// 				return idol.Height, nil
// 			},
// 		},
// 		"birth_place": &graphql.Field{
// 			Type: graphql.String,
// 			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 				idol := p.Source.(Idol)
// 				return idol.Birthplace, nil
// 			},
// 		},
// 		"birth_day": &graphql.Field{
// 			Type: graphql.String,
// 			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 				idol := p.Source.(Idol)
// 				return idol.Birthday, nil
// 			},
// 		},
// 		"blood_type": &graphql.Field{
// 			Type: graphql.String,
// 			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 				idol := p.Source.(Idol)
// 				return idol.Bloodtype, nil
// 			},
// 		},
// 	},
// })

// func getSameAgeIdols(db *sql.DB, o Option) []Idol {
// 	var stx string
// 	if o.age == 0 {
// 		stx = "select * from idol order by id"
// 	} else {
// 		stx = fmt.Sprintf("select * from idol where age=%d order by id", o.age)
// 	}
// 	rows, err := db.Query(stx)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	var result []Idol
// 	for rows.Next() {
// 		var i Idol
// 		rows.Scan(&i.Id, &i.Name, &i.Age, &i.Height, &i.Birthplace, &i.Birthday, &i.Bloodtype, &i.Unit)

// 		if o.age == 0 || i.Age == o.age {
// 			result = append(result, i)
// 		}

// 	}

// 	return result
// }

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
							result := models.GetSameAgeIdols(db, models.IdolByAgeOption{Age: ageQuery})
							return result, nil
						} else {
							result := models.GetSameAgeIdols(db, models.IdolByAgeOption{})
							return result, nil
						}
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
