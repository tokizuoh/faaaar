package fields

import (
	"github/tokizuoh/faaaar/server/models"

	"github.com/graphql-go/graphql"
)

var IdolsFieldKey = "idols"

var IdolsField = &graphql.Field{
	Type: graphql.NewList(models.IdolType),
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		age, ok := p.Args["age"].(int)
		var o models.IdolsByAgeOption
		if ok {
			o = models.IdolsByAgeOption{Age: age}
		}

		result := models.GetSameAgeIdols(o)
		return result, nil
	},
	Args: graphql.FieldConfigArgument{
		"age": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
}
