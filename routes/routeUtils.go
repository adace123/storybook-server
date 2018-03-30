package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"../models"
	mgo "gopkg.in/mgo.v2"
)

// HandleErr - handles error
func handleErr(err error, message string) bool {
	if err != nil {
		log.Fatal(message, err)
	}
	return true
}

func configResponse(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json")
}

func convertAllToJSON(entity interface{}) []byte {
	log.Printf("%T", entity)
	res, err := json.Marshal(entity)
	handleErr(err, "Could not convert model to JSON")
	return res
}

func getConfig() models.Config {
	file, err := os.Open("config.json")
	handleErr(err, "Could not open config file")
	decoder := json.NewDecoder(file)
	var config models.Config
	err = decoder.Decode(&config)
	handleErr(err, "Could not convert config file to JSON")
	return config
}

func connect() {
	config := getConfig()
	session, err := mgo.Dial(config.MongoURL)
	handleErr(err, "Could not connect to DB")
	db = session.DB("storybooks")
}
