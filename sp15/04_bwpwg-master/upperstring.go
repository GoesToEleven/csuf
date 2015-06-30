package stringupper

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

func init() {
	http.HandleFunc("/", root)
	http.HandleFunc("/upper", upper)
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, rootForm)
}

const rootForm = `
  <!DOCTYPE html>
    <html>
      <head>
        <meta charset="utf-8">
        <link rel="stylesheet" href="css/upper.css">
        <title>String Upper</title>
      </head>
      <body>
        <h1>String Upper</h1>
        <p>The String Upper Service will accept a string from you and 
           return you the Uppercase version of the original string. Have fun!</p>
        <form action="/upper" method="post" accept-charset="utf-8">
	  <input type="text" name="str" value="Type a string..." id="str">
	  <input type="submit" value=".. and change to uppercase!">
        </form>
      </body>
    </html>
`
var upperTemplate = template.Must(template.New("upper").Parse(upperTemplateHTML))

func upper(w http.ResponseWriter, r *http.Request) {
        strEntered := r.FormValue("str")
        strUpper := strings.ToUpper(strEntered)
        err := upperTemplate.Execute(w, strUpper)
        if err != nil {
	        http.Error(w, err.Error(), http.StatusInternalServerError)
        }
}

const upperTemplateHTML = ` 
<!DOCTYPE html>
  <html>
    <head>
      <meta charset="utf-8">
      <link rel="stylesheet" href="css/upper.css">
      <title>String Upper Results</title>
    </head>
    <body>
      <h1>String Upper Results</h1>
      <p>The Uppercase of the string that you had entered is:</p>
      <pre>{{html .}}</pre>
    </body>
  </html>
`



