package fields

import (
	"github/tokizuoh/faaaar/server/models"

	"github.com/graphql-go/graphql"
)

var IdolsField = &graphql.Field{
	Type: graphql.NewList(models.IdolType),
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		ageQuery, ok := p.Args["age"].(int)
		if ok {
			result := models.GetSameAgeIdols(models.IdolsByAgeOption{Age: ageQuery})
			return result, nil
		} else {
			result := models.GetSameAgeIdols(models.IdolsByAgeOption{})
			return result, nil
		}
	},
	Args: graphql.FieldConfigArgument{
		"age": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
}
