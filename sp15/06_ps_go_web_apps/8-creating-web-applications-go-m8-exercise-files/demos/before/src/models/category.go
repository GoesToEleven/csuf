package models

import (
	"errors"
)

type Category struct {
	imageUrl string
	title string
	description string
	isOrientRight bool
	id int
	products []Product
}

func (this *Category) ImageUrl() string {
	return this.imageUrl
}
func (this *Category) Title() string {
	return this.title
}
func (this *Category) Description() string {
	return this.description
}
func (this *Category) IsOrientRight() bool {
	return this.isOrientRight
}
func (this *Category) Id() int {
	return this.id
}
func (this *Category) Products() []Product {
	return this.products
}

func (this *Category) SetImageUrl(value string) {
	this.imageUrl = value
}
func (this *Category) SetTitle(value string) {
	this.title = value
}
func (this *Category) SetDescription(value string) {
	this.description = value
}
func (this *Category) SetIsOrientRight(value bool) {
	this.isOrientRight = value
}
func (this *Category) SetId(value int) {
	this.id = value
}
func (this *Category) SetProducts(value []Product) {
	this.products = value
}

func GetCategories() []Category {
	result := []Category{
		Category{
	 		imageUrl: "lemon.png",
	 		title: "Juices and Mixes",
	 		description: `Explore our wide assortment of juices and mixes expected by
								today's lemonade stand clientelle. Now featuring a full line of
								organic juices that are guaranteed to be obtained from trees that
								have never been treated with pesticides or artificial
								fertilizers.`,
			isOrientRight: false,
			id: 1,
			products: GetJuiceProducts(),
	 	}, Category{
	 		imageUrl: "kiwi.png",
	 		title: "Cups, Straws, and Other Supplies",
	 		description: `From paper cups to bio-degradable plastic to straws and
							napkins, LSS is your source for the sundries that keep your stand
							running smoothly.`,
			isOrientRight: true,
			id: 2,
	 	}, Category{
	 		imageUrl: "pineapple.png",
	 		title: "Signs and Advertising",
	 		description: `Sure, you could just wait for people to find your stand
							along the side of the road, but if you want to take it to the next
							level, our premium line of advertising supplies.`,
			isOrientRight: false,
			id: 3,
	 	},		
	}
	
	return result
}

func GetCategoryById(id int) (Category, error) {
	for _, category := range GetCategories() {
		if category.Id() == id {
			return category, nil
		}
	}
	
	return Category{}, errors.New("Category Not Found")
}