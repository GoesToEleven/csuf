package models

import (
	"testing"
)

func Test_ReturnsNonEmptySlice(t *testing.T) {
	categories := GetCategories()
	
	if len(categories) == 0 {
		t.Log("Non empty slice returned")
		t.Fail()
	}
}