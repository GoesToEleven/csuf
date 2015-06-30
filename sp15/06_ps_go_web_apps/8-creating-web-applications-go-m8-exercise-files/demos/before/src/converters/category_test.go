package converters

import (
	"testing"
	"models"
)

func Test_ConvertsCategoryToViewModel(t *testing.T) {
	category := models.Category{}
	category.SetImageUrl("the image URL")
	category.SetTitle("the title")
	category.SetDescription("the description")
	category.SetIsOrientRight(true)
	category.SetId(42)
	category.SetProducts([]models.Product{
			models.Product{},
			models.Product{},
		})
	
	result := ConvertCategoyToViewModel(category)
	
	if result.ImageUrl != category.ImageUrl() {
		t.Log("Image URL not converted properly")
		t.Fail()
	}
	
	if result.Title != category.Title() {
		t.Log("Title not converted properly")
		t.Fail()
	}
	
	if result.Description != category.Description() {
		t.Log("Description not converted properly")
		t.Fail()
	}
	
	if result.IsOrientRight != category.IsOrientRight() {
		t.Log("IsOrientRight not converted properly")
		t.Fail()
	}
	
	if result.Id != category.Id() {
		t.Log("Id not converted properly")
		t.Fail()
	}
	
	if len(result.Products) != len(category.Products()) {
		t.Log("Products not converted properly")
		t.Fail()
	}
}