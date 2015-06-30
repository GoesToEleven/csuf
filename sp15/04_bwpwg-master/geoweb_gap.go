package geoweb

import (
	"encoding/json"
	"fmt"
        "html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"appengine"
        "appengine/urlfetch"
)

func init() {
        http.HandleFunc("/", handler)
        http.HandleFunc("/showimage", showimage)
}

func handler(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, rootForm)
}

const rootForm = `
  <!DOCTYPE html>
    <html>
      <head>
        <meta charset="utf-8">
        <title>Go View</title>
        <link rel="stylesheet" href="/stylesheets/goview.css">        
      </head>
      <body>
        <h1><img style="margin-left: 120px;" src="images/gsv.png" alt="Go View" />GoView</h1>
        <h2>Accept Address</h2>
        <p>Please enter your address:</p>
        <form style="margin-left: 120px;" action="/showimage" method="post" accept-charset="utf-8">
      <input type="text" name="str" value="Type address..." id="str" />
      <input type="submit" value=".. and see the image!" />
        </form>
      </body>
    </html>
`

var upperTemplate = template.Must(template.New("showimage").Parse(upperTemplateHTML))

func showimage(w http.ResponseWriter, r *http.Request) {
        // Sample address "1600 Amphitheatre Parkway, Mountain View, CA"
        addr := r.FormValue("str")

	// QueryEscape escapes the addr string so
	// it can be safely placed inside a URL query
	// safeAddr := url.QueryEscape(addr)
        safeAddr := url.QueryEscape(addr)
        fullUrl := fmt.Sprintf("http://maps.googleapis.com/maps/api/geocode/json?sensor=false&address=%s", safeAddr)

        c := appengine.NewContext(r)
        client := urlfetch.Client(c)
    
        resp, err := client.Get(fullUrl)
        if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
        }
    
	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	// Read the content into a byte array
	body, dataReadErr := ioutil.ReadAll(resp.Body)
	if dataReadErr != nil {
		log.Fatal("ReadAll: ", dataReadErr)
	}

        res := make(map[string][]map[string]map[string]map[string]interface{}, 0)

	// We will be using the Unmarshal function
	// to transform our JSON bytes into the
	// appropriate structure.
	// The Unmarshal function accepts a byte array
	// and a reference to the object which shall be
	// filled with the JSON data (this is simplifying,
	// it actually accepts an interface)
	json.Unmarshal(body, &res)
        
	lat, _ := res["results"][0]["geometry"]["location"]["lat"]
	lng, _ := res["results"][0]["geometry"]["location"]["lng"]
	
	// %.13f is used to convert float64 to a string
	queryUrl := fmt.Sprintf("http://maps.googleapis.com/maps/api/streetview?sensor=false&size=600x300&location=%.13f,%.13f", lat, lng)

        tempErr := upperTemplate.Execute(w, queryUrl)
        if tempErr != nil {
	        http.Error(w, tempErr.Error(), http.StatusInternalServerError)
        }
}

const upperTemplateHTML = ` 
<!DOCTYPE html>
  <html>
    <head>
      <meta charset="utf-8">
      <title>Display Image</title>
      <link rel="stylesheet" href="/stylesheets/goview.css">              
    </head>
    <body>
      <h1><img style="margin-left: 120px;" src="images/gsv.png" alt="Street View" />GoView</h1>
      <h2>Image at your Address</h2>
      <img style="margin-left: 120px;" src="{{html .}}" alt="Image" />
    </body>
  </html>
`

