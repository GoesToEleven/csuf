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

	row := db.QueryRow("SELECT subject FROM todos where id = $1", 1)
	var subject string
	row.Scan(&subject)
	log.Printf("Subject is %s", subject)
}
