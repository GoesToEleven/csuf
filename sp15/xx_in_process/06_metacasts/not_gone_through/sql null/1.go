package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type User struct {
	Id    int    `db:"id" json:"id"`
	Email string `db:"email" json:"email"`
	Name  string `db:"name" json:"name"`
}

func main() {
	db, err := sqlx.Open("postgres", "postgres://postgres:postgres@localhost/sql_development?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	user := &User{}
	err = db.Get(user, "select * from users")

	if err != nil {
		log.Fatal(err)
	}

	log.Println(user)
	log.Println(user.Email)
	log.Println(user.Name)
}
