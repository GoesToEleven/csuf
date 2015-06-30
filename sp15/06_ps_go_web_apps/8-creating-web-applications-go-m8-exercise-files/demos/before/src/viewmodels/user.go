package viewmodels

import (

)

type Profile struct {
	Title string
	Active string
	User User
}

type User struct {
	Id int
	Email string
	FirstName string
	LastName string
	Stand Stand
}

type Stand struct {
	Address string
	City string
	State string
	Zip string
}

func GetProfile() Profile {
	result := Profile{
		Title: "Lemonade Stand Supply - Profile",
	}
	
	return result
}