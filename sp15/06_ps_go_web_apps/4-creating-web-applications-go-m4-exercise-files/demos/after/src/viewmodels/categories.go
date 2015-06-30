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
}

func GetCategories() Categories {
	result := Categories{
		Title: "Lemonade Stand Society - Shop",
		Active: "shop",
	}
	
	juiceCategory := Category{
 		ImageUrl: "lemon.png",
 		Title: "Juices and Mixes",
 		Description: `Explore our wide assortment of juices and mixes expected by
							today's lemonade stand clientelle. Now featuring a full line of
							organic juices that are guaranteed to be obtained from trees that
							have never been treated with pesticides or artificial
							fertilizers.`,
		IsOrientRight: false,
 	}
	
	supplyCategory := Category{
 		ImageUrl: "kiwi.png",
 		Title: "Cups, Straws, and Other Supplies",
 		Description: `From paper cups to bio-degradable plastic to straws and
						napkins, LSS is your source for the sundries that keep your stand
						running smoothly.`,
		IsOrientRight: true,
 	}
	
	advertiseCategory := Category{
 		ImageUrl: "pineapple.png",
 		Title: "Signs and Advertising",
 		Description: `Sure, you could just wait for people to find your stand
						along the side of the road, but if you want to take it to the next
						level, our premium line of advertising supplies.`,
		IsOrientRight: false,
 	}
	
	result.Categories = []Category{
		juiceCategory,
		supplyCategory,
		advertiseCategory,
	}
	
	return result
}