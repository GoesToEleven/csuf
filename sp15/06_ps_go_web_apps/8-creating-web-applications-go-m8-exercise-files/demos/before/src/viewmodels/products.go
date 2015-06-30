package viewmodels

import ()

type Products struct {
	Title    string
	Active   string
	Products []Product
}

func GetProducts(name string) Products {
	var result Products
	result.Active = "shop"
	result.Title = "Lemonade Stand Society - " + name + " Shop"

	return result
}

type ProductVM struct {
	Title   string
	Active  string
	Product Product
}

func GetProduct(name string) ProductVM {
	var result ProductVM
	
	result.Active = "shop"
	result.Title = "Lemonade Stand Society - " + name

	return result
}

type Product struct {
	Name             string
	DescriptionShort string
	DescriptionLong  string
	PricePerLiter    float32
	PricePer10Liter  float32
	Origin           string
	IsOrganic        bool
	ImageUrl         string
	Id               int
}