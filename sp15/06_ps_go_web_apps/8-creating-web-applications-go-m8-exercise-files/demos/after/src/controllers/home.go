package controllers

import (
	"net/http"
	"viewmodels"
	"text/template"
	"controllers/util"
	"models"
)

type homeController struct {
	template *template.Template
	loginTemplate *template.Template
}

func (this *homeController) get(w http.ResponseWriter, req *http.Request) {
	vm := viewmodels.GetHome()
	
	sessionCookie, err := req.Cookie("sessionId")
	if err == nil {
		member, err := models.GetMemberBySessionId(sessionCookie.Value)
		if err == nil {
			vm.Member.IsLoggedIn = true
			vm.Member.FirstName = member.FirstName()
		}
	}
	
	w.Header().Add("Content Type", "text/html")
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()
	
	this.template.Execute(responseWriter, vm)
}

func (this *homeController) login(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()
	
	responseWriter.Header().Add("Content Type", "text/html")
	
	if req.Method == "POST" {
		email := req.FormValue("email")
		password := req.FormValue("password")
		
		member, err := models.GetMember(email, password) 
		
		if err == nil {
			session, err := models.CreateSession(member)
			if err == nil {
				var cookie http.Cookie
				cookie.Name = "sessionId"
				cookie.Value = session.SessionId()
				responseWriter.Header().Add("Set-Cookie", cookie.String())
			}
		}
	}
	
	vm := viewmodels.GetLogin()
	
	this.loginTemplate.Execute(responseWriter, vm)	
}