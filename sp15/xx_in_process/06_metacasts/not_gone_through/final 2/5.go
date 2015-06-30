package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"time"
)

type Todo struct {
	Id          int       `db:"id"`
	Subject     string    `db:"subject"`
	Description string    `db:"description"`
	Completed   bool      `db:"completed"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func main() {
	db, err := sqlx.Open("postgres", "postgres://postgres:postgres@localhost/godos_development?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	todos := []Todo{}
	db.Select(&todos, "select * from todos")

	for _, todo := range todos {
		log.Printf("Subject is %s", todo.Subject)
	}
}
