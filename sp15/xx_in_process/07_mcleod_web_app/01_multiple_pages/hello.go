package hello

import (
	"html/template"
	"io/ioutil"
	"net/http"
)

var (
	guestbookForm []byte
	signTemplate  = template.Must(template.ParseFiles("guestbook.html"))
)

func init() {
	content, err := ioutil.ReadFile("guestbookform.html")
	if err != nil {
		panic(err)
	}
	guestbookForm = content

	http.HandleFunc("/", root)
	http.HandleFunc("/sign", sign)
}

func root(w http.ResponseWriter, r *http.Request) { w.Write(guestbookForm) }

func sign(w http.ResponseWriter, r *http.Request) {
	err := signTemplate.Execute(w, r.FormValue("content"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

/*
NOTES:
-this is the original appengine-guestbook-go site
--we've already looked at this
--you can find this site on github
https://github.com/GoesToEleven/appengine-guestbook-go
-notice:
--"root" uses w.Write
--"sign" uses Execute
http://golang.org/pkg/net/http/#ResponseWriter
http://golang.org/pkg/html/template/#Template.Execute
---Question: why this choice?
---Consideration:
-----"root" doesn't "merge your data into the form letter (template)"
-----"sign" is "merging data into the form letter (template)"

see this running here:
http://decoded-theme-89623.appspot.com/
--note:
---- if it's not running, I took it down
---- you can follow these instructions to run it locally
https://cloud.google.com/appengine/docs/go/gettingstarted/helloworld
---- you can follow these instructions to run on app engine
https://cloud.google.com/appengine/docs/go/gettingstarted/uploading

-notice:
-- func init()
https://golang.org/doc/effective_go.html#init

*/

// STEP 1: create a new template with some name
// STEP 2: parse the string into the template
// // in lay terms: "give the template your form letter"
// // in lay terms: "put your form letter into the template"
// STEP 3: execute the template
// // merge template 'tmpl' with content of 's'
// // lay terms: "merge your data into the form letter"

/*

handle
handlefunc
handler
handlerfunc

http://golang.org/pkg/net/http/#Handle				// takes a string & a handler
http://golang.org/pkg/net/http/#HandleFunc		// takes a string & a handlerfunc
http://golang.org/pkg/net/http/#Handler				// defines the handler interface
http://golang.org/pkg/net/http/#HandlerFunc		// a func that implements the handler interface

*/
