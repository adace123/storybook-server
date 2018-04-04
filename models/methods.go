package models

import (
	"../utils"
	"errors"
	"golang.org/x/crypto/bcrypt"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"os"
)

// DB - mLab connector
var DB *mgo.Database

func connect() {
	// config := utils.GetConfigFile()
	session, err := mgo.Dial(os.Getenv("MONGO_URL"))
	utils.HandleErr(err, "Could not connect to DB")
	DB = session.DB("storybooks")
}

func init() {
	connect()
}

// FindAll - returns all authors from DB
func (authorsDAO AuthorsDAO) FindAll() (interface{}, error) {
	var authors []Author

	authorDB := DB.C(authorsDAO.Collection)
	err := authorDB.Find(nil).All(&authors)
	return authors, err
}

// FindAll - returns all stories from DB
func (storiesDAO StoriesDAO) FindAll() (interface{}, error) {
	var stories []Story
	storyDB := DB.C(storiesDAO.Collection)
	err := storyDB.Find(bson.M{"visibility": "PUBLIC"}).All(&stories)
	return stories, err
}

// FindAll - returns all comments from DB
func (commentsDAO CommentsDAO) FindAll() (interface{}, error) {
	var comments []Comment
	commentDB := DB.C(commentsDAO.Collection)
	err := commentDB.Find(nil).All(&comments)
	return comments, err
}

func (authorsDAO AuthorsDAO) FindByEmail(email string) (interface{}, error) {
	var author Author
	authorDB := DB.C(authorsDAO.Collection)
	err := authorDB.Find(bson.M{"email": email}).One(&author)
	return author, err
}

// FindByID - returns single author
func (authorsDAO AuthorsDAO) FindByID(id bson.ObjectId) (interface{}, error) {
	var author Author
	authorDB := DB.C(authorsDAO.Collection)
	err := authorDB.FindId(id).One(&author)
	return author, err
}

// FindByID - returns single story
func (storiesDAO StoriesDAO) FindByID(id bson.ObjectId) (interface{}, error) {
	var story Story
	storyDB := DB.C(storiesDAO.Collection)
	err := storyDB.FindId(id).One(&story)
	return story, err
}

// FindByID - returns single comment
func (commentsDAO CommentsDAO) FindByID(id bson.ObjectId) (interface{}, error) {
	var comment Comment
	commentDB := DB.C(commentsDAO.Collection)
	err := commentDB.FindId(id).One(&comment)
	return comment, err
}

// FindByVisibility - get stories by author and filter by visibility
func (storiesDAO StoriesDAO) FindByVisibility(authorID bson.ObjectId, visibility string) (interface{}, error) {
	var stories []Story
	storyDB := DB.C(storiesDAO.Collection)
	err := storyDB.Find(bson.M{"visibility": visibility, "author": authorID}).Sort("-timestamp").All(&stories)
	return stories, err
}

// FindCommentsByStory - get all comments for a story
func (commentsDAO CommentsDAO) FindCommentsByStory(storyID bson.ObjectId) (interface{}, error) {
	var comments []Comment
	commentDB := DB.C(commentsDAO.Collection)
	err := commentDB.Find(bson.M{"story": storyID}).Sort("-timestamp").All(&comments)
	return comments, err
}

// FindCommentsByAuthor - get all comments for an author
func (commentsDAO CommentsDAO) FindCommentsByAuthor(authorID bson.ObjectId) (interface{}, error) {
	var comments []Comment
	commentDB := DB.C(commentsDAO.Collection)
	err := commentDB.Find(bson.M{"author": authorID}).Sort("-timestamp").All(&comments)
	return comments, err
}

// FindStoriesByAuthor - get all stories for an author
func (storiesDAO StoriesDAO) FindStoriesByAuthor(authorID bson.ObjectId) (interface{}, error) {
	var stories []Story
	storyDB := DB.C(storiesDAO.Collection)
	err := storyDB.Find(bson.M{"author": authorID}).Sort("-timestamp").All(&stories)
	return stories, err
}

// CreateAuthor - create new author
func (authorsDAO AuthorsDAO) SignUp(author *Author) (interface{}, error) {
	_, err := authorsDAO.FindByEmail(author.Email)
	if err != nil {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(author.Password), bcrypt.DefaultCost)
		utils.HandleErr(err, "Could not generate password")
		(*author).Password = string(hashedPassword)
		(*author).ID = bson.NewObjectId()
		e := DB.C(authorsDAO.Collection).Insert(&author)
		utils.HandleErr(e, "Could not insert author")
		return authorsDAO.FindByEmail(author.Email)
	}
	return Author{}, errors.New("Author already exists")
}
