package fields

import (
	"github/tokizuoh/faaaar/server/models"
	"github/tokizuoh/faaaar/server/resolvers"

	"github.com/graphql-go/graphql"
)

var IdolsFieldKey = "idols"

var IdolsField = &graphql.Field{
	Type: graphql.NewList(models.IdolType),
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		// TODO: [#30] 値が入っていない時は ageがゼロ値になっていることを確認
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
