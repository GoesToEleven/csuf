package controllers

import (
	"text/template"
	"net/http"
	"github.com/gorilla/mux"
	"viewmodels"
	"strconv"
	"controllers/util"
	"models"
	"converters"
)

type productController struct {
	template *template.Template
}

func (this *productController) get(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	
	idRaw := vars["id"]
	
	id, err := strconv.Atoi(idRaw)
	if err == nil {
		product, err := models.GetProductById(id)
		
		if err == nil {
			
			w.Header().Add("Content Type", "text/html")
			responseWriter := util.GetResponseWriter(w, req)
			defer responseWriter.Close()
		
			vm := viewmodels.GetProduct(product.Name())
			vm.Product = converters.ConvertProductToViewModel(product)
			
			this.template.Execute(responseWriter, vm)
		}
	} else {
		w.WriteHeader(404)
	}
}