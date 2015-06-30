package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type User struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type HomePageHandler struct{}

func (h *HomePageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	json.NewDecoder(r.Body).Decode(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	user.CreatedAt = time.Now()
	data, _ := json.Marshal(user)
	w.Write(data)
}

func main() {
	http.Handle("/", &HomePageHandler{})

	log.Fatal(http.ListenAndServe(":3000", nil))
}
