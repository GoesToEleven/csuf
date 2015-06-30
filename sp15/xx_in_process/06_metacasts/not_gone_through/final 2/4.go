package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"time"
)

type Todo struct {
	id          int
	subject     string
	description string
	completed   bool
	created_at  time.Time
	updated_at  time.Time
}

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost/godos_development?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("select * from todos")
	for rows.Next() {
		todo := Todo{}
		if err := rows.Scan(&todo.id, &todo.subject, &todo.description, &todo.completed, &todo.created_at, &todo.updated_at); err != nil {
			log.Fatal(err)
		}
		log.Printf("ID is %d", todo.id)
		log.Printf("Subject is %s", todo.subject)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
