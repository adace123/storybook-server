package gql

import (
	"log"

	"../models"
	"github.com/graphql-go/graphql"
)

var authorsDAO = models.AuthorsDAO{Collection: "authors"}
var storiesDAO = models.StoriesDAO{Collection: "stories"}
var commentsDAO = models.CommentsDAO{Collection: "comments"}
var StorybookSchema graphql.Schema

func init() {

	AuthorType.AddFieldConfig("comments", &graphql.Field{
		Type: graphql.NewList(CommentType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return commentsDAO.FindCommentsByAuthor(p.Source.(models.Author).ID)
		},
	})

	AuthorType.AddFieldConfig("stories", &graphql.Field{
		Type: graphql.NewList(StoryType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return storiesDAO.FindByVisibility(p.Source.(models.Author).ID, "PUBLIC")
		},
	})

	schemaConfig := graphql.SchemaConfig{Query: QueryType, Mutation: MutationType}
	Schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("Could not create schema: %v", err)
	}
	StorybookSchema = Schema
}

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
