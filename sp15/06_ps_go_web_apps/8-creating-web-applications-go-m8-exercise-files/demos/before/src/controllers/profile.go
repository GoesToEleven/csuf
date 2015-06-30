package controllers

import (
	"controllers/util"
	"net/http"
	"text/template"
	"viewmodels"
)

type profileController struct {
	template *template.Template
}

func (this *profileController) handle(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()
	
	vm := viewmodels.GetProfile()
	if req.Method == "POST" {
		vm.User.Email = req.FormValue("email")
		vm.User.FirstName = req.FormValue("firstName")
		vm.User.LastName = req.FormValue("lastName")
		vm.User.Stand.Address = req.FormValue("standAddress")
		vm.User.Stand.City = req.FormValue("standCity")
		vm.User.Stand.State = req.FormValue("standState")
		vm.User.Stand.Zip = req.FormValue("standZip")
	}
	
	responseWriter.Header().Add("Content Type", "text/html")
	this.template.Execute(responseWriter, vm)
}