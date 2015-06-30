package converters

import (
	"models"
	"viewmodels"
)

func ConvertCategoyToViewModel(category models.Category) viewmodels.Category {
	result := viewmodels.Category{
		ImageUrl: category.ImageUrl(),
		Title: category.Title(),
		Description: category.Description(),
		IsOrientRight: category.IsOrientRight(),
		Id: category.Id(),
	}
	
	for _, p := range category.Products() {
		result.Products = append(result.Products, ConvertProductToViewModel(p))
	}
	
	return result
}