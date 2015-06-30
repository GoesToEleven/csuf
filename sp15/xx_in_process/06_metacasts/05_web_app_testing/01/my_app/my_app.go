package my_app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func HomePageHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello, World!")
}

func WelcomeByNameHandler(res http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(res, "Hello, %s!", name)
}

func JsonHandler(res http.ResponseWriter, req *http.Request) {
	user := new(User)
	json.NewDecoder(req.Body).Decode(user)
	user.Id = 1

	res.WriteHeader(http.StatusCreated)
	res.Header().Set("Content-Type", "application/json")
	data, _ := json.Marshal(user)
	res.Write(data)
}

func NewMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/json", JsonHandler)
	mux.HandleFunc("/hello", WelcomeByNameHandler)
	mux.HandleFunc("/", HomePageHandler)
	return mux
}
