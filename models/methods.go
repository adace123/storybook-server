package models

import mgo "gopkg.in/mgo.v2"

// FindAll - returns all authors from DB
func (authorsDAO AuthorsDAO) FindAll(db *mgo.Database) (interface{}, error) {
	var authors []Author

	authorDB := db.C(authorsDAO.Collection)
	err := authorDB.Find(nil).All(&authors)
	return authors, err
}

// FindAll - returns all stories from DB
func (storiesDAO StoriesDAO) FindAll(db *mgo.Database) (interface{}, error) {
	var stories []Story
	storyDB := db.C(storiesDAO.Collection)
	err := storyDB.Find(nil).All(&stories)
	return stories, err
}

// FindAll - returns all comments from DB
func (commentsDAO CommentsDAO) FindAll(db *mgo.Database) (interface{}, error) {
	var comments []Comment
	commentDB := db.C(commentsDAO.Collection)
	err := commentDB.Find(nil).All(&comments)
	return comments, err
}
