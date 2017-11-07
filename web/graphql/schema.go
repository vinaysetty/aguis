package graphql

import (
	"github.com/graphql-go/graphql"
)

var Schema = graphql.NewSchema(graphql.SchemaConfig{
	Query: QueryType,
})
