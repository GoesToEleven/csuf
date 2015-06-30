package controllers

import (
	"text/template"
	"net/http"
	"viewmodels"
	"github.com/gorilla/mux"
	"strconv"
	"controllers/util"
)

type categoriesController struct {
	template *template.Template
}

func (this *categoriesController) get(w http.ResponseWriter, req *http.Request) {
	vm := viewmodels.GetCategories()
	
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
		vm := viewmodels.GetProducts(id)
		
		w.Header().Add("Content Type", "text/html")
		responseWriter := util.GetResponseWriter(w, req)
		defer responseWriter.Close()
		
		this.template.Execute(responseWriter, vm)
	} else {
		w.WriteHeader(404)
	}
}