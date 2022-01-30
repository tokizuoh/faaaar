package fields

import (
	"github/tokizuoh/faaaar/server/models"

	"github.com/graphql-go/graphql"
)

var UnitsField = &graphql.Field{
	Type: graphql.NewList(models.UnitType),
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		idolIdQuery, ok := p.Args["idolId"].(int)
		if ok {
			result, err := models.GetUnitsByIdolID(models.UnitsByIdolIdOption{IdolId: idolIdQuery})
			if err != nil {
				return nil, err
			}
			return result, nil
		} else {
			result, err := models.GetUnitsByIdolID(models.UnitsByIdolIdOption{})
			if err != nil {
				return nil, err
			}
			return result, nil
		}
	},
	Args: graphql.FieldConfigArgument{
		"idolId": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
}
