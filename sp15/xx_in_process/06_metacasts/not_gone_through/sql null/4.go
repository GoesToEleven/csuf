package main

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type MyNullString sql.NullString

func (ns *MyNullString) Scan(value interface{}) error {
	n := sql.NullString{String: ns.String}
	err := n.Scan(value)
	ns.String, ns.Valid = n.String, n.Valid
	return err
}

func (ns MyNullString) Value() (driver.Value, error) {
	n := sql.NullString{String: ns.String}
	return n.Value()
}

func (s MyNullString) MarshalJSON() ([]byte, error) {
	if s.Valid {
		return json.Marshal(s.String)
	}
	return json.Marshal(nil)
}

type User struct {
	Id    int          `db:"id" json:"id"`
	Email string       `db:"email" json:"email"`
	Name  MyNullString `db:"name" json:"name"`
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

	json.NewEncoder(os.Stdout).Encode(user)
}
