package router

import (
	"../gql"
	"../utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

var Router = mux.NewRouter()

func init() {
	Router.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		log.Println("graphql route")
		defer r.Body.Close()
		utils.ConfigResponse(&w)
		resp, err := ioutil.ReadAll(r.Body)
		utils.HandleErr(err, "Could not ready body")
		result := gql.ExecuteQuery(string(resp), gql.StorybookSchema)
		json.NewEncoder(w).Encode(result)
	}).Methods("POST")

}

func StartServer() {
	log.Fatal(http.ListenAndServe(":8080", Router))
}
