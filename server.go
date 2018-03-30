package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"./gql"
	"./utils"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		utils.ConfigResponse(&w)
		resp, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal("Could not read body")
		}
		result := gql.ExecuteQuery(string(resp), gql.StorybookSchema)
		json.NewEncoder(w).Encode(result)
	}).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
