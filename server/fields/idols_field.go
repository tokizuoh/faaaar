package fields

import (
	"github/tokizuoh/faaaar/server/resolvers"
	"github/tokizuoh/faaaar/server/types"

	"github.com/graphql-go/graphql"
)

var IdolsFieldKey = "idols"

var IdolsField = &graphql.Field{
	Type: graphql.NewList(types.IdolType),
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		age, _ := p.Args["age"].(int)
		result, err := resolvers.GetIdolsByAge(age)
		if err != nil {
			return nil, err
		}
		return result, nil
	},
	Args: graphql.FieldConfigArgument{
		"age": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
}
