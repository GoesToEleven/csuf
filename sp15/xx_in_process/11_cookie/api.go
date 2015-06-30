package main

import (
	//"fmt"
	"html/template"
	"net/http"
    "time"
    "log"
    "encoding/json"
)

var mytemp *template.Template

type userstuff struct{
    Name string `json:"Name"`
    Email string `json:"Email"`
    Message string `json:"Message"`
}

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/formhandler", formhandler)
	http.HandleFunc("/link1", linkone)
	http.HandleFunc("/link2", linktwo)
	//http.HandleFunc("/link3", linkthree)
    mytemp = template.Must(template.ParseFiles("form.html","formhandler.html","link1.html","link2.html"/*,"link3.html"*/))
}

func handler(w http.ResponseWriter, r *http.Request) {
	err := mytemp.ExecuteTemplate(w, "form.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func formhandler(w http.ResponseWriter, r *http.Request) {
    expire := time.Now().AddDate(0, 0, 1)

    thisUser := userstuff{
        Name: r.FormValue("name"),
        Email: r.FormValue("email"),
        Message: r.FormValue("msg"),
    }

    userData, _ := json.Marshal(thisUser)
    log.Printf("jSoN encode; %v", userData)
    for i,val := range userData{
        if val==34 {
            userData[i] = 96
        }
    }
    cookie := http.Cookie{
        Name:       "test",
        Value:      string(userData),
        //Path:       "/",
        //Domain:     "localhost:8080",
        Expires:    expire,
        RawExpires: expire.Format(time.UnixDate),
        // MaxAge=0 means no 'Max-Age' attribute specified.
        // MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
        // MaxAge>0 means Max-Age attribute present and given in seconds
        MaxAge:   86400,
        Secure:   false,
        HttpOnly: false,
        //Raw:      "raw",
        //Unparsed: []string{}, // Raw text of unparsed attribute-value pairs}
    }
    http.SetCookie(w, &cookie)

    err := mytemp.ExecuteTemplate(w, "formhandler.html", struct{
        Name string
        Email string
        Message string
        }{r.FormValue("name"),r.FormValue("email"),r.FormValue("msg")})

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func linkone(w http.ResponseWriter, r *http.Request) {
    thisuser := r.FormValue("name")+" "+r.FormValue("email")+" "+r.FormValue("message")
    err := mytemp.ExecuteTemplate(w, "link1.html", thisuser)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func linktwo(w http.ResponseWriter, r *http.Request) {
    //read from cookie
    var thisUser userstuff
    cookie, _ := r.Cookie("test")
    log.Printf("cookie2: %v", cookie.Value)
    //array of string
    mybytes := []byte(cookie.Value)
    for i,val := range mybytes{
        if val==96 {
            mybytes[i] = 34
        }
    }
    err := json.Unmarshal(mybytes, &thisUser)
    log.Print(thisUser)
    err = mytemp.ExecuteTemplate(w, "link2.html", thisUser)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
