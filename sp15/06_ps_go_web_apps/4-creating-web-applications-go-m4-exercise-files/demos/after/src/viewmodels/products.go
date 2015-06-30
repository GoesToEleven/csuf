package viewmodels

import (

)

type Products struct {
	Title string
	Active string
	Products []Product
}

func GetProducts() Products {
	var result Products
	result.Active = "shop"
	result.Title = "Lemonade Stand Society - Juice Shop"
	
	lemonJuice := MakeLemonJuiceProduct()
	appleJuice := MakeAppleJuiceProduct()
	watermelonJuice := MakeWatermelonJuiceProduct()
	kiwiJuice := MakeKiwiJuiceProduct()
	mangosteenJuice := MakeMangosteenJuiceProduct()
	orangeJuice := MakeOrangeJuiceProduct()
	pineappleJuice := MakePineappleJuiceProduct()
	strawberryJuice := MakeStrawberryJuiceProduct()
	
	result.Products = []Product{
		lemonJuice,
		appleJuice,
		watermelonJuice,
		kiwiJuice,
		mangosteenJuice,
		orangeJuice,
		pineappleJuice,
		strawberryJuice,
	}
	
	return result
}

type ProductVM struct {
	Title string
	Active string
	Product Product
}

func GetProduct() ProductVM {
	var result ProductVM
	
	result.Active = "shop"
	result.Title = "Lemonade Stand Society - Lemon Juice"
	
	result.Product = MakeLemonJuiceProduct()
	
	return result
}

type Product struct {
	Name string
	DescriptionShort string
	DescriptionLong string
	PricePerLiter float32
	PricePer10Liter float32
	Origin string
	IsOrganic bool
	ImageUrl string
	Id int
}

func MakeLemonJuiceProduct() Product {
	result := Product{
		Name: "Lemon Juice",
		DescriptionShort: "Made from fresh, organic California lemons.",
		DescriptionLong: `Made from premium, organic Meyer lemons. These fruit are left on the tree until they reach the peak of ripeness and then juiced within 8 hours of being picked.
			<br/>
			Lemonade made from our premium juice is sure to make your stand the most popular in the neighborhood.`,
		PricePerLiter: 1.09,
		PricePer10Liter: 1.04,
		Origin: "California",
		IsOrganic: true,
		ImageUrl: "lemon.png",
		Id: 1,
	}
	
	return result
}

func MakeAppleJuiceProduct() Product {
	result := Product{
		Name: "Apple Juice",
		DescriptionShort: "The perfect blend of Washington apples.",
		DescriptionLong: "The perfect blend of Washington apples.",
		PricePerLiter: 0.89,
		PricePer10Liter: 0.85,
		Origin: "Ohio",
		IsOrganic: true,
		ImageUrl: "apple.png",
		Id: 2,
	}
	
	return result
}

func MakeWatermelonJuiceProduct() Product {
	result := Product{
		Name: "Watermelon Juice",
		DescriptionShort: "From sun-drenched fields in Florida.",
		DescriptionLong: "From sun-drenched fields in Florida.",
		PricePerLiter: 3.99,
		PricePer10Liter: 3.84,
		Origin: "Florida",
		IsOrganic: true,
		ImageUrl: "watermelon.png",
		Id: 3,
	}
	
	return result
}

func MakeKiwiJuiceProduct() Product {
	result := Product{
		Name: "Kiwi Juice",
		DescriptionShort: "California sunshine and rain distilled into sweet goodness",
		DescriptionLong: "California sunshine and rain distilled into sweet goodness",
		PricePerLiter: 5.99,
		PricePer10Liter: 5.54,
		Origin: "California",
		IsOrganic: false,
		ImageUrl: "kiwi.png",
		Id: 4,
	}
	
	return result
}

func MakeMangosteenJuiceProduct() Product {
	result := Product{
		Name: "Mangosteen Juice",
		DescriptionShort: "Our latest taste sensation, imported directly from Hawaii",
		DescriptionLong: "Our latest taste sensation, imported directly from Hawaii",
		PricePerLiter: 6.87,
		PricePer10Liter: 6.79,
		Origin: "Hawaii",
		IsOrganic: false,
		ImageUrl: "mangosteen.png",
		Id: 5,
	}
	
	return result
}

func MakeOrangeJuiceProduct() Product {
	result := Product{
		Name: "Orange Juice",
		DescriptionShort: "Fresh squeezed from Florida's best oranges.",
		DescriptionLong: "Fresh squeezed from Florida's best oranges.",
		PricePerLiter: 1.67,
		PricePer10Liter: 1.63,
		Origin: "Florida",
		IsOrganic: false,
		ImageUrl: "orange.png",
		Id: 6,
	}
	
	return result
}

func MakePineappleJuiceProduct() Product {
	result := Product{
		Name: "Pineapple Juice",
		DescriptionShort: "An exotic and refreshing offering. Straight from Hawaii.",
		DescriptionLong: "An exotic and refreshing offering. Straight from Hawaii.",
		PricePerLiter: 2.48,
		PricePer10Liter: 2.27,
		Origin: "Hawaii",
		IsOrganic: false,
		ImageUrl: "pineapple.png",
		Id: 7,
	}
	
	return result
}

func MakeStrawberryJuiceProduct() Product {
	result := Product{
		Name: "Strawberry Juice",
		DescriptionShort: "MThe perfect balance of sweet and tart.",
		DescriptionLong: "The perfect balance of sweet and tart.",
		PricePerLiter: 4.36,
		PricePer10Liter: 4.27,
		Origin: "California",
		IsOrganic: false,
		ImageUrl: "strawberry.png",
		Id: 8,
	}
	
	return result
}