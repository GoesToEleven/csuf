package FormPage

import (
	"html/template"
	"net/http"
)

func init() {
	http.Handle("/hiddenDir/", http.StripPrefix("/hiddenDir/", http.FileServer(http.Dir("assets/"))))

	http.HandleFunc("/", root)
	http.HandleFunc("/response", response)
}

func root(rw http.ResponseWriter, req *http.Request) {
	mytemp := template.Must(template.ParseFiles("assets/layout.html", "assets/root.html"))
	mytemp.ExecuteTemplate(rw, "layout", nil)
}

func response(rw http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")
	switch name {
	case "Corey Dihel":
		mytemp := template.Must(template.ParseFiles("assets/layout.html", "assets/master.html"))
		mytemp.ExecuteTemplate(rw, "layout", nil)
	default:
		mytemp := template.Must(template.ParseFiles("assets/layout.html", "assets/guest.html"))
		mytemp.ExecuteTemplate(rw, "layout", name)
	}
}
