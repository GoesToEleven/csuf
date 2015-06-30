package main

import (
	"encoding/json"
	"os"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string
	Email     string `json:"email,omitempty"`
	Password  string `json:"-"`
}

func main() {
	u := User{"Mark", "Bates", "mark@example.com", "password"}
	b, _ := json.Marshal(u)
	os.Stdout.Write(b)
}
