package main

import (
	"html/template"
	"net/http"
    "log"
)

var mytemp *template.Template

type userstuff struct {
	Name, Email, Message string
}

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/formhandler", formhandler)
	http.HandleFunc("/link1", linkone)
	http.HandleFunc("/link2", linktwo)
	http.HandleFunc("/link3", linkthree)
	http.HandleFunc("/link4", linkfour)
	mytemp = template.Must(template.ParseFiles("form.html", "formhandler.html", "link1.html",
    "link2.html", "link3.html", "link4.html"))
}

func handler(w http.ResponseWriter, r *http.Request) {
	err := mytemp.ExecuteTemplate(w, "form.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func formhandler(w http.ResponseWriter, r *http.Request) {
	user := userstuff{r.FormValue("name"), r.FormValue("email"), r.FormValue("message")}
	err := mytemp.ExecuteTemplate(w, "formhandler.html", user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func linkone(w http.ResponseWriter, r *http.Request) {

    //r.URL.Query.Get
    log.Println(r.URL.Query())
    log.Println(r.URL.Query().Get("name"))
    log.Println(r.URL.Query().Get("email"))
    log.Println(r.URL.Query().Get("message"))

    //r.URL.Query
    mappedUser := r.URL.Query()
    log.Println(mappedUser)
    log.Println(mappedUser["name"])
    log.Println(mappedUser["email"])
    log.Println(mappedUser["message"])
    log.Println(mappedUser["name"][0])
    log.Println(mappedUser["email"][0])
    log.Println(mappedUser["message"][0])
//    log.Println(mappedUser["message"][1])
//    log.Println(mappedUser["message"][2])

    // into struct
    user := userstuff{
        r.URL.Query().Get("name"),
        r.URL.Query().Get("email"),
        r.URL.Query().Get("message")}
    log.Println(user)
    log.Println(user.Name)
    log.Println(user.Email)
    log.Println(user.Message)

    // r.URL
    log.Println(r.URL)



    err := mytemp.ExecuteTemplate(w, "link1.html", user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func linktwo(w http.ResponseWriter, r *http.Request) {
    mappedUser := r.URL.Query()
    err := mytemp.ExecuteTemplate(w, "link2.html", mappedUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func linkthree(w http.ResponseWriter, r *http.Request) {
//    data := []int{1, 5, 15, 20, 25, 30}
//    moredata := []strings{"one", "two", "three"}
    mappedUser := r.URL.Query()
//    mappedUser["data"] = data
//    mappedUser["moredata"] = moredata
//    log.Println("mapped ",mappedUser)
//    log.Println(mappedUser["name"])
//    log.Println(mappedUser["email"])
//    log.Println(mappedUser["message"])
//    log.Println(mappedUser["data"])
//    log.Println(mappedUser["moredata"])
    err := mytemp.ExecuteTemplate(w, "link3.html", mappedUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func linkfour(w http.ResponseWriter, r *http.Request) {
//    mappedUser := r.URL.Query()
    user := userstuff{ r.URL.Query().Get("name"), r.URL.Query().Get("email"), r.URL.Query().Get("message")}
	err := mytemp.ExecuteTemplate(w, "link4.html", user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
