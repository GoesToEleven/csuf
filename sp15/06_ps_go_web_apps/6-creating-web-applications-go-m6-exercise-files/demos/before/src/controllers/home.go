package controllers

import (
	"net/http"
	"viewmodels"
	"text/template"
	"controllers/util"
)

type homeController struct {
	template *template.Template
}

func (this *homeController) get(w http.ResponseWriter, req *http.Request) {
	vm := viewmodels.GetHome()
	
	w.Header().Add("Content Type", "text/html")
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()
	
	this.template.Execute(responseWriter, vm)
}