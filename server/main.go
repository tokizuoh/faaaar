package main

import (
	"encoding/json"
	"github/tokizuoh/faaaar/server/fields"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	_ "github.com/lib/pq"
)

func readQuery(filepath string) (string, error) {
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

const QUERY_FILE_PATH = "./query.txt"

func executeQuery(query string) {
	scheme, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				fields.IdolsFieldKey: fields.IdolsField,
				fields.UnitsFieldKey: fields.UnitsField,
			},
		}),
	})

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

func main() {
	http.HandleFunc("/graphql", func(rw http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}

		query := string(body)
		executeQuery(query)
	})
	http.ListenAndServe(":8080", nil)
}
