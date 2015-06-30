package main

import (
	"log"
	"reflect"
	"regexp"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string
	Email     string `json:"email,omitempty"`
	Password  string `json:"-"`
}

func ColumnsForStruct(s interface{}) []string {
	columns := []string{}
	st := reflect.TypeOf(s)
	field_count := st.NumField()

	for i := 0; i < field_count; i++ {
		field := st.Field(i)

		tag := field.Tag.Get("json")
		if tag == "" {
			tag = field.Name
		}
		if tag != "-" {
			s := regexp.MustCompile(",").Split(tag, -1)
			columns = append(columns, s[0])
		}
	}

	return columns
}
func main() {
	u := User{"Mark", "Bates", "", "password"}
	columns := ColumnsForStruct(u)

	for _, column := range columns {
		log.Println(column)
	}
}
