package utils

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	err := godotenv.Load()
	HandleErr(err, "Couldn't load env file ")
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

// func GetConfigFile() Config {
// 	file, err := os.Open("config.json")
// 	HandleErr(err, "Could not open config file")
// 	decoder := json.NewDecoder(file)
// 	var config Config
// 	err = decoder.Decode(&config)
// 	HandleErr(err, "Could not convert config file to JSON")
// 	return config
// }
