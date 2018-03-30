package gql

import (
	"../models"
	"github.com/graphql-go/graphql"
)

var AuthorType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "author",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"password": &graphql.Field{
				Type: graphql.String,
			},
			"imageURL": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var StoryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "story",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"content": &graphql.Field{
				Type: graphql.String,
			},
			"allowComments": &graphql.Field{
				Type: graphql.Boolean,
			},
			"author": &graphql.Field{
				Type: AuthorType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return authorsDAO.FindByID(p.Source.(models.Story).Author)
				},
			},
		},
	},
)

var CommentType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "comment",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
			},
			"content": &graphql.Field{
				Type: graphql.String,
			},
			"author": &graphql.Field{
				Type: AuthorType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return authorsDAO.FindByID(p.Source.(models.Comment).Author)
				},
			},
			"story": &graphql.Field{
				Type: StoryType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return authorsDAO.FindByID(p.Source.(models.Comment).Story)
				},
			},
		},
	},
)

var StoryInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "StoryInput",
	Fields: graphql.Fields{
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"content": &graphql.Field{
			Type: graphql.String,
		},
		"allowComments": &graphql.Field{
			Type: graphql.Boolean,
		},
		"visibility": &graphql.Field{
			Type: graphql.String,
		},
	},
})

// AuthorInput - author input object
var AuthorInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "AuthorInput",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"password": &graphql.Field{
			Type: graphql.String,
		},
		"imageURL": &graphql.Field{
			Type: graphql.String,
		},
	},
})
