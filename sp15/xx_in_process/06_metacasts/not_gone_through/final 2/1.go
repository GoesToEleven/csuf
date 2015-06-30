package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost/godos_development?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	// Getting a single value
	var subject string

	rows, err := db.Query("SELECT subject FROM todos")
	for rows.Next() {
		if err := rows.Scan(&subject); err != nil {
			log.Fatal(err)
		}
		log.Printf("Subject is %s", subject)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
