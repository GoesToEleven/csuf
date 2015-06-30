package controllers

import (
	"text/template"
	"net/http"
	"viewmodels"
	"github.com/gorilla/mux"
	"strconv"
	"controllers/util"
	"models"
	"converters"
)

type categoriesController struct {
	template *template.Template
}

func (this *categoriesController) get(w http.ResponseWriter, req *http.Request) {
	categories := models.GetCategories()
	
	categoriesVM := []viewmodels.Category{}
	for _, category := range categories {
		categoriesVM = append(categoriesVM, converters.ConvertCategoyToViewModel(category))
	}
	
	vm := viewmodels.GetCategories()
	vm.Categories = categoriesVM
	
	w.Header().Add("Content Type", "text/html")
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()
	
	this.template.Execute(responseWriter, vm)
}

type categoryController struct {
	template *template.Template
}

func (this *categoryController) get(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	
	idRaw := vars["id"]
	
	id, err := strconv.Atoi(idRaw)
	if err == nil {
		category, err := models.GetCategoryById(id)
		
		if err == nil {
		
			w.Header().Add("Content Type", "text/html")
			responseWriter := util.GetResponseWriter(w, req)
			defer responseWriter.Close()
		
			vm := viewmodels.GetProducts(category.Title())
			productVMs := []viewmodels.Product{}
			
			for _, product := range category.Products() {
				productVMs = append(productVMs, converters.ConvertProductToViewModel(product))
			}
			
			vm.Products = productVMs
			
			this.template.Execute(responseWriter, vm)
		}
	} else {
		w.WriteHeader(404)
	}
}