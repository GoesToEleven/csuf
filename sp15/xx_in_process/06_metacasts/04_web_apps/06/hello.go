package main

import (
	"encoding/json"
	"net/http"
	"time"
)

// IMPORTANT
// to run this file
// install POSTMAN for google chrome
// https://www.google.ca/webhp?sourceid=chrome-instant&ion=1&espv=2&ie=UTF-8#q=chrome%20postman
// then use POSTMAN to make a POST request
// here's the prebuilt postmant JSON file you want to send
// you can load this into postman
// https://www.getpostman.com/collections/21646f5b294eb318519a

type User struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func AnyName(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	json.NewDecoder(r.Body).Decode(user)
	user.CreatedAt = time.Now()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	data, _ := json.Marshal(user)
	w.Write(data)

}

func main() {
	http.HandleFunc("/", AnyName)

	http.ListenAndServe(":8080", nil)
}
