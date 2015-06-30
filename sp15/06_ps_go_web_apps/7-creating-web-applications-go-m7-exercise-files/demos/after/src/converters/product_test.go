package converters

import (
	"testing"
	"models"
)

func Test_ConvertsProductToViewModel(t *testing.T) {
	product := models.Product{}
	product.SetName("the name")
	product.SetDescriptionShort("the short description")
	product.SetDescriptionLong("the long description")
	product.SetPricePerLiter(42.127)
	product.SetPricePer10Liter(27.314)
	product.SetOrigin("the origin")
	product.SetIsOrganic(true)
	product.SetImageUrl("the image URL")
	product.SetId(42)
	
	result := ConvertProductToViewModel(product)
	
		if product.Name() != result.Name {
		t.Fail()
	}
	if product.DescriptionShort() != result.DescriptionShort {
		t.Fail()
	}
	if product.DescriptionLong() != result.DescriptionLong {
		t.Fail()
	}
	if product.PricePerLiter() != result.PricePerLiter {
		t.Fail()
	}
	if product.PricePer10Liter() != result.PricePer10Liter {
		t.Fail()
	}
	if product.Origin() != result.Origin {
		t.Fail()
	}
	if product.IsOrganic() != result.IsOrganic {
		t.Fail()
	}
	if product.ImageUrl() != result.ImageUrl {
		t.Fail()
	}
	if product.Id() != result.Id {
		t.Fail()
	}
}