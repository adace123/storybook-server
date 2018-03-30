package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// ModelDAO interface
type ModelDAO interface {
	FindAll(db *mgo.Database) (interface{}, error)
}

// Config represents config.json
type Config struct {
	MongoURL string
}

// Author model
type Author struct {
	ID       bson.ObjectId `bson:"_id" json:"_id"`
	Email    string        `bson:"email" json:"email"`
	Password string        `bson:"password" json:"password"`
	ImageURL string        `bson:"imageURL" json:"imageURL"`
	Name     string        `bson:"name" json:"name"`
}

// Story model
type Story struct {
	ID            bson.ObjectId     `bson:"_id" json:"_id"`
	Title         string            `bson:"title" json:"title"`
	Content       string            `bson:"content" json:"content"`
	AllowComments bool              `bson:"allowComments" json:"allowComments"`
	Author        Author            `bson:"author" json:"author"`
	Visibility    map[string]string `bson:"visibility" json:"visibility"`
}

// Comment model
type Comment struct {
	ID      bson.ObjectId `bson:"_id" json:"_id"`
	Content string        `bson:"content" json:"content"`
	Story   Story         `bson:"story" json:"story"`
	Author  Author        `bson:"author" json:"author"`
}

// Visibility model
type Visibility struct {
	PUBLIC      string `bson:"PUBLIC" json:"PUBLIC"`
	PRIVATE     string `bson:"PRIVATE" json:"PRIVATE"`
	UNPUBLISHED string `bson:"UNPUBLISHED" json:"UNPUBLISHED"`
}

// AuthorsDAO - interacts with Authors collection
type AuthorsDAO struct {
	Collection string
}

// StoriesDAO - interacts with Stories collection
type StoriesDAO struct {
	Collection string
}

// CommentsDAO - interacts with Comments collection
type CommentsDAO struct {
	Collection string
}
