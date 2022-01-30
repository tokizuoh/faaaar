package fields

import (
	"github/tokizuoh/faaaar/server/models"
	"github/tokizuoh/faaaar/server/resolvers"

	"github.com/graphql-go/graphql"
)

var UnitsFieldKey = "units"

var UnitsField = &graphql.Field{
	Type: graphql.NewList(models.UnitType),
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		// TODO: [#30] 値が入っていない時は idolIdがゼロ値になっていることを確認
		idolId, _ := p.Args["idolId"].(int)
		result, err := resolvers.GetUnitsByIdolID(idolId)
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
