package fields

import (
	"github/tokizuoh/faaaar/server/models"

	"github.com/graphql-go/graphql"
)

var UnitsFieldKey = "units"

var UnitsField = &graphql.Field{
	Type: graphql.NewList(models.UnitType),
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		idolId, ok := p.Args["idolId"].(int)
		var o models.UnitsByIdolIdOption
		if ok {
			o = models.UnitsByIdolIdOption{IdolId: idolId}
		}

		result, err := models.GetUnitsByIdolID(o)
		if err != nil {
			return nil, err
		}
		return result, nil
	},
	Args: graphql.FieldConfigArgument{
		"idolId": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
}
