package gql

import (
	"../models"
	"github.com/graphql-go/graphql"
	"gopkg.in/mgo.v2/bson"
)

// MutationType - all GraphQL mutations
var MutationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createStory": &graphql.Field{
				Type: StoryType,
				Args: graphql.FieldConfigArgument{
					"authorID": &graphql.ArgumentConfig{
						Type: graphql.ID,
					},
					"story": &graphql.ArgumentConfig{
						Type: StoryInput,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var author models.Author = p.Args["story"].(models.Author)
					author.ID = bson.ObjectIdHex(p.Args["authorID"].(string))
					newAuthor, err := authorsDAO.CreateAuthor(&author)
					return newAuthor, err
				},
			},
		},
	},
)
