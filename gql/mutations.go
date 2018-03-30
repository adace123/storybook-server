package gql

import (
	"../models"
	"github.com/graphql-go/graphql"
)

// MutationType - all GraphQL mutations
var MutationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createAuthor": &graphql.Field{
				Type: AuthorType,
				Args: graphql.FieldConfigArgument{
					"author": &graphql.ArgumentConfig{
						Type: AuthorInput,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var authorFields = p.Args["author"].(map[string]interface{})
					var author models.Author = models.Author{
						Name:     authorFields["name"].(string),
						Email:    authorFields["email"].(string),
						Password: authorFields["password"].(string),
						ImageURL: authorFields["imageURL"].(string),
					}
					newAuthor, err := authorsDAO.CreateAuthor(&author)
					return newAuthor, err
				},
			},
		},
	},
)
