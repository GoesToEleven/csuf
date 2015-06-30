package models

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func getDBConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", "dbname=lss_db user=lss_admin password=admin sslmode=disable")
	return db, err
}