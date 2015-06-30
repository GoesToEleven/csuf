package main

import (
	"net/http"
	"fmt"
	"html/template"
//	"log"
//	"os"
//	"strings"
	"io/ioutil"
)

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/formResult", formHandler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	rootForm, err := ioutil.ReadFile("templates/prompt.html");
	if err != nil {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, string(rootForm))
}

var resultFile, _ = ioutil.ReadFile("templates/result.html");
var resultHtmlTemplate = template.Must(template.New("result").Parse(string(resultFile)))

func formHandler(w http.ResponseWriter, r *http.Request) {
	strEntered := r.FormValue("str")

	var err error

	if strEntered == "Shawn" {
		err = resultHtmlTemplate.Execute(w, "You spelled it correctly")
	} else {
		err = resultHtmlTemplate.Execute(w, "...you spelled it wrong")
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
