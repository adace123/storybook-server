package gql

import (
	"github.com/graphql-go/graphql"
	"gopkg.in/mgo.v2/bson"
)

// QueryType - main GraphQL query
var QueryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"stories": &graphql.Field{
				Type: graphql.NewList(StoryType),
				Args: graphql.FieldConfigArgument{
					"authorID": &graphql.ArgumentConfig{
						Type: graphql.ID,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if id, ok := p.Args["authorID"]; ok {
						return storiesDAO.FindStoriesByAuthor(bson.ObjectIdHex(id.(string)))
					}
					return storiesDAO.FindAll()
				},
			},
			"comments": &graphql.Field{
				Type: graphql.NewList(CommentType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return commentsDAO.FindAll()
				},
			},
			"authors": &graphql.Field{
				Type: graphql.NewList(AuthorType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return authorsDAO.FindAll()
				},
			},
			"story": &graphql.Field{
				Type: StoryType,
				Args: graphql.FieldConfigArgument{
					"storyID": &graphql.ArgumentConfig{
						Type: graphql.ID,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return storiesDAO.FindByID(bson.ObjectIdHex(p.Args["storyID"].(string)))
				},
			},
			"author": &graphql.Field{
				Type: AuthorType,
				Args: graphql.FieldConfigArgument{
					"authorID": &graphql.ArgumentConfig{
						Type: graphql.ID,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return authorsDAO.FindByID(bson.ObjectIdHex(p.Args["authorID"].(string)))
				},
			},
			"comment": &graphql.Field{
				Type: CommentType,
				Args: graphql.FieldConfigArgument{
					"commentID": &graphql.ArgumentConfig{
						Type: graphql.ID,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return commentsDAO.FindByID(bson.ObjectIdHex(p.Args["commentID"].(string)))
				},
			},
		},
	},
)
