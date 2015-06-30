package viewmodels

import (

)

type Categories struct {
	Title string
	Active string
	Categories []Category
}

type Category struct {
	ImageUrl string
	Title string
	Description string
	IsOrientRight bool
	Id int
	Products []Product
}

func GetCategories() Categories {
	result := Categories{
		Title: "Lemonade Stand Society - Shop",
		Active: "shop",
	}
	
	return result
}