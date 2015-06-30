package main

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"log"
	"strings"

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

func (z *MyNullString) UnmarshalJSON(text []byte) error {
	z.Valid = false
	if string(text) == "null" {
		return nil
	}
	s := ""
	err := json.Unmarshal(text, &s)
	if err == nil {
		z.String = s
		z.Valid = true
	}
	return err
}

type User struct {
	Id    int          `db:"id" json:"id"`
	Email string       `db:"email" json:"email"`
	Name  MyNullString `db:"name" json:"name"`
}

func main() {
	u := &User{}
	x := `{"id":1,"email":"john@example.com","name":null}`
	json.NewDecoder(strings.NewReader(x)).Decode(u)
	log.Println(u)
	log.Println(u.Email)
	log.Println(u.Name)
	log.Println(u.Name.String)
}
