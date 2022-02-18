package fields

import (
	"github/tokizuoh/faaaar/server/resolvers"
	"github/tokizuoh/faaaar/server/types"

	"github.com/graphql-go/graphql"
)

var UnitsFieldKey = "units"

var UnitsField = &graphql.Field{
	Type: graphql.NewList(types.UnitType),
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		idolId, _ := p.Args["idolId"].(int)
		if idolId != 0 {
			result, err := resolvers.GetUnitsByIdolID(idolId)
			if err != nil {
				return nil, err
			}
			return result, nil
		} else {
			result, err := resolvers.GetUnits()
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
