package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// ModelDAO interface
type ModelDAO interface {
	FindAll(db *mgo.Database) (interface{}, error)
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
	ID            bson.ObjectId `bson:"_id" json:"_id"`
	Title         string        `bson:"title" json:"title"`
	Content       string        `bson:"content" json:"content"`
	AllowComments bool          `bson:"allowComments" json:"allowComments"`
	Author        bson.ObjectId `bson:"author" json:"author"`
	Visibility    string        `bson:"visibility" json:"visibility"`
}

// Comment model
type Comment struct {
	ID      bson.ObjectId `bson:"_id" json:"_id"`
	Content string        `bson:"content" json:"content"`
	Story   bson.ObjectId `bson:"story" json:"story"`
	Author  bson.ObjectId `bson:"author" json:"author"`
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
