package main

import (
	"fmt"
	"html/template"
        "labix.org/v2/mgo"
        "labix.org/v2/mgo/bson"
        "log"
	"net/http"
	"os"
)

type CountryCapital struct {
        Country string
        Capital string
}

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/upper", upper)
        fmt.Println("listening...")
        err := http.ListenAndServe(GetPort(), nil)
        if err != nil {
                log.Fatal("ListenAndServe: ", err)
        }
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
        <title>Country Capital</title>
      </head>
      <body>
        <h1>Country Capital</h1>
        <p>The Country Capital Service will accept a name of a country from you and return you that country's capital. Have fun!</p>
        <form action="/upper" method="post" accept-charset="utf-8">
	  <input type="text" name="str" value="Type a string..." id="str">
	  <input type="submit" value=".. and fetch capital!">
        </form>
      </body>
    </html>
`
var upperTemplate = template.Must(template.New("upper").Parse(upperTemplateHTML))

func upper(w http.ResponseWriter, r *http.Request) {
        // In the open command window set the following for Heroku:
        // heroku config:set MONGOHQ_URL=mongodb://IndianGuru:password@troup.mongohq.com:10080/godata
        uri := os.Getenv("MONGOHQ_URL")
        if uri == "" {
                fmt.Println("no connection string provided")
                os.Exit(1)
        }
 
        sess, err := mgo.Dial(uri)
        if err != nil {
                fmt.Printf("Can't connect to mongo, go error %v\n", err)
                os.Exit(1)
        }
        defer sess.Close()
        
        sess.SetSafe(&mgo.Safe{})
        
        collection := sess.DB("godata").C("capital")

        result := CountryCapital{}

        strEntered := r.FormValue("str")

        err = collection.Find(bson.M{"country": strEntered}).One(&result)
        if err != nil {
                log.Fatal("Find: ", err)
        }
        
        err2 := upperTemplate.Execute(w, result.Capital)
        if err2 != nil {
	        http.Error(w, err2.Error(), http.StatusInternalServerError)
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
      <h1>Country Capital</h1>
      <p>The Capital of the Country that you had entered is:</p>
      <pre>{{html .}}</pre>
    </body>
  </html>
`

// Get the Port from the environment so we can run on Heroku
func GetPort() string {
        var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "4747"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}



