package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"time"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost/godos_development?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	now := time.Now()
	res, err := db.Exec("INSERT INTO todos (subject, description, created_at, updated_at) VALUES ($1, $2, $3, $4)", "Mow Lawn", "", now, now)
	if err != nil {
		log.Fatal(err)
	}
	affected, _ := res.RowsAffected()
	log.Printf("Rows affected %d", affected)

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
