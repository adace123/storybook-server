package api_tests

import (
	"../router"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestGetStories(t *testing.T) {
	go router.StartServer()

	gql := `
        query {
            stories {
                title
            }
        }
    `
	res, _ := http.NewRequest("POST", "localhost:8080/graphql", strings.NewReader(gql))
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
