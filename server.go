package main

import (
	"log"
	"net/http"

	"./routes"
	"github.com/gorilla/handlers"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS()(routes.Router)))
}
