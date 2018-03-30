package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"regexp"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

}

// Config represents config.json
type Config struct {
	MongoURL string
}

// HandleErr - handles error
func HandleErr(err error, message string) bool {
	if err != nil {
		log.Fatal(message, err)
	}
	return true
}

func ConfigResponse(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json")
}

func ConvertAllToJSON(entity interface{}) []byte {
	log.Printf("%T", entity)
	res, err := json.Marshal(entity)
	HandleErr(err, "Could not convert model to JSON")
	return res
}

func GetConfigFile() Config {
	file, err := os.Open("config.json")
	HandleErr(err, "Could not open config file")
	decoder := json.NewDecoder(file)
	var config Config
	err = decoder.Decode(&config)
	HandleErr(err, "Could not convert config file to JSON")
	return config
}

func validEmail(email string) bool {
	return regexp.MustCompile("^((\\w+[\\.]?)+)@(\\w+\\.){1,}\\w{2,9}$").MatchString(email)
}
