package models

import (
	"errors"
)

type Product struct {
	name string
	descriptionShort string
	descriptionLong string
	pricePerLiter float32
	pricePer10Liter float32
	origin string
	isOrganic bool
	imageUrl string
	id int
}

func (this *Product) Name() string {
	return this.name
}
func (this *Product) DescriptionShort() string {
	return this.descriptionShort
}
func (this *Product) DescriptionLong() string {
	return this.descriptionLong
}
func (this *Product) PricePerLiter() float32 {
	return this.pricePerLiter
}
func (this *Product) PricePer10Liter() float32 {
	return this.pricePer10Liter
}
func (this *Product) Origin() string {
	return this.origin
}
func (this *Product) IsOrganic() bool {
	return this.isOrganic
}
func (this *Product) ImageUrl() string {
	return this.imageUrl
} 
func (this *Product) Id() int {
	return this.id
} 

func (this *Product) SetName(value string) {
	this.name = value
}
func (this *Product) SetDescriptionShort(value string) {
	this.descriptionShort = value
}
func (this *Product) SetDescriptionLong(value string) {
	this.descriptionLong = value
}
func (this *Product) SetPricePerLiter(value float32) {
	this.pricePerLiter = value
}
func (this *Product) SetPricePer10Liter(value float32) {
	this.pricePer10Liter = value
}
func (this *Product) SetOrigin(value string) {
	this.origin = value
}
func (this *Product) SetIsOrganic(value bool) {
	this.isOrganic = value
}
func (this *Product) SetImageUrl(value string) {
	this.imageUrl = value
}
func (this *Product) SetId(value int) {
	this.id = value
}

func GetJuiceProducts() []Product {
	result := []Product{
		makeLemonJuiceProduct(),
		makeAppleJuiceProduct(),
		makeWatermelonJuiceProduct(),
		makeKiwiJuiceProduct(),
		makeMangosteenJuiceProduct(),
		makeOrangeJuiceProduct(),
		makePineappleJuiceProduct(),
		makeStrawberryJuiceProduct(),
	}
	
	return result
}

func GetProductById(id int) (Product, error) {
	products := GetJuiceProducts()
	
	for _, p := range products {
		if p.Id() == id {
			return p, nil
		}
	}
	
	return Product{}, errors.New("Product not found")
}

func makeLemonJuiceProduct() Product {
	result := Product{
		name: "Lemon Juice",
		descriptionShort: "Made from fresh, organic California lemons.",
		descriptionLong: `Made from premium, organic Meyer lemons. These fruit are left on the tree until they reach the peak of ripeness and then juiced within 8 hours of being picked.
			<br/>
			Lemonade made from our premium juice is sure to make your stand the most popular in the neighborhood.`,
		pricePerLiter: 1.09,
		pricePer10Liter: 1.04,
		origin: "California",
		isOrganic: true,
		imageUrl: "lemon.png",
		id: 1,
	}
	
	return result
}

func makeAppleJuiceProduct() Product {
	result := Product{
		name: "Apple Juice",
		descriptionShort: "The perfect blend of Washington apples.",
		descriptionLong: "The perfect blend of Washington apples.",
		pricePerLiter: 0.89,
		pricePer10Liter: 0.85,
		origin: "Ohio",
		isOrganic: true,
		imageUrl: "apple.png",
		id: 2,
	}
	
	return result
}

func makeWatermelonJuiceProduct() Product {
	result := Product{
		name: "Watermelon Juice",
		descriptionShort: "From sun-drenched fields in Florida.",
		descriptionLong: "From sun-drenched fields in Florida.",
		pricePerLiter: 3.99,
		pricePer10Liter: 3.84,
		origin: "Florida",
		isOrganic: true,
		imageUrl: "watermelon.png",
		id: 3,
	}
	
	return result
}

func makeKiwiJuiceProduct() Product {
	result := Product{
		name: "Kiwi Juice",
		descriptionShort: "California sunshine and rain distilled into sweet goodness",
		descriptionLong: "California sunshine and rain distilled into sweet goodness",
		pricePerLiter: 5.99,
		pricePer10Liter: 5.54,
		origin: "California",
		isOrganic: false,
		imageUrl: "kiwi.png",
		id: 4,
	}
	
	return result
}

func makeMangosteenJuiceProduct() Product {
	result := Product{
		name: "Mangosteen Juice",
		descriptionShort: "Our latest taste sensation, imported directly from Hawaii",
		descriptionLong: "Our latest taste sensation, imported directly from Hawaii",
		pricePerLiter: 6.87,
		pricePer10Liter: 6.79,
		origin: "Hawaii",
		isOrganic: false,
		imageUrl: "mangosteen.png",
		id: 5,
	}
	
	return result
}

func makeOrangeJuiceProduct() Product {
	result := Product{
		name: "Orange Juice",
		descriptionShort: "Fresh squeezed from Florida's best oranges.",
		descriptionLong: "Fresh squeezed from Florida's best oranges.",
		pricePerLiter: 1.67,
		pricePer10Liter: 1.63,
		origin: "Florida",
		isOrganic: false,
		imageUrl: "orange.png",
		id: 6,
	}
	
	return result
}

func makePineappleJuiceProduct() Product {
	result := Product{
		name: "Pineapple Juice",
		descriptionShort: "An exotic and refreshing offering. Straight from Hawaii.",
		descriptionLong: "An exotic and refreshing offering. Straight from Hawaii.",
		pricePerLiter: 2.48,
		pricePer10Liter: 2.27,
		origin: "Hawaii",
		isOrganic: false,
		imageUrl: "pineapple.png",
		id: 7,
	}
	
	return result
}

func makeStrawberryJuiceProduct() Product {
	result := Product{
		name: "Strawberry Juice",
		descriptionShort: "MThe perfect balance of sweet and tart.",
		descriptionLong: "The perfect balance of sweet and tart.",
		pricePerLiter: 4.36,
		pricePer10Liter: 4.27,
		origin: "California",
		isOrganic: false,
		imageUrl: "strawberry.png",
		id: 8,
	}
	
	return result
}