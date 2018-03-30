package gql

import (
	"log"

	"github.com/graphql-go/graphql"
)

// ExecuteQuery - executes a GraphQL query
func ExecuteQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		log.Printf("Something went wrong: %v", result.Errors)
	}
	return result
}
