package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"time"
)

type Todo struct {
	Id          int
	Subject     string
	Description string
	Completed   bool
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func main() {
	db, err := sqlx.Open("postgres", "postgres://postgres:postgres@localhost/godos_development?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	tx := db.MustBegin()
	now := time.Now()
	t := Todo{
		Subject:     "Mow Lawn!",
		Description: "Yuck!",
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	tx.Exec("INSERT INTO todos (subject, description, created_at, updated_at) VALUES ($1, $2, $3, $4)", t.Subject, t.Description, t.CreatedAt, t.UpdatedAt)
	// // demonstrate the transaction:
	// tx.Exec("INSERT INTO todos (subject, description, created_at, updated_at) VALUES ($1, $2, $3, $4)", t.Subject, t.Description, t.CreatedAt, t.UpdatedAt)
	// tx.Exec("INSERT INTO todos (subject, description, created_at, updated_at) VALUES ($1, $2, $3, $4)", nil, t.Description, t.CreatedAt, t.UpdatedAt)
	tx.Commit()

	todos := []Todo{}
	db.Select(&todos, "select * from todos")

	for _, todo := range todos {
		log.Printf("Subject is %s", todo.Subject)
	}
}
