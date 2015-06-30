package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

type GeoCodeJson struct {
	Status string `json: "status"`
	Results []GeoJsonResults `json: "results"`
}

type GeoJsonResults struct {
	Formatted_address string `json: "formatted_address"`
	Types []string `json: "types"`
	Address_components []GeoJsonAddrComp `json: "address_components"`
	Geometry GeoJsonGeometry `json: "geometry"`
}

type GeoJsonAddrComp struct {
	Long_name string `json: "long_name"`
	Short_name string `json: "short_name"`
	Types []string `json: "types"`
}

type GeoJsonGeometry struct {
	Location_type string `json: "location_type"`
	Location GeoJsonCoordinate `json: "location"`
	Viewport GeoJsonViewport `json: "viewport"`
}

type GeoJsonViewport struct {
	Northeast GeoJsonCoordinate `json: "northeast"`
	Southwest GeoJsonCoordinate `json: "southwest"`
}

type GeoJsonCoordinate struct {
	Lat float64 `json: "lat"`
	Lng float64 `json: "lng"`
}

//Alternatively pack everything into one struct
//type GeoCodeJson struct {
//	Status string `json: "status"`
//	Results []struct {
//		Address_components []struct {
//			Long_name string `json: "long_name"`
//			Short_name string `json: "short_name"`
//			Types []string `json: "types"`
//		}`json: "address_components"`
//		Formatted_address string `json: "formatted_address"`
//		Geometry struct {
//			Location struct {
//				Lat float64 `json: "lat"`
//				Lng float64 `json: "lng"`
//			}`json: "location"`
//			Location_type string `json: "location_type"`
//			Viewport struct {
//				Northeast struct {
//					Lat float64 `json: "lat"`
//					Lng float64 `json: "lng"`
//				}`json: "northeast"`
//				Southwest struct {
//					Lat float64 `json: "lat"`
//					Lng float64 `json: "lng"`
//				}`json: "southwest"`
//			}`json: "viewport"`
//		}`json: "geometry"`
//		Types []string `json: "types"`
//	}`json: "results"`
//}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/showimage", showimage)
	fmt.Println("listening...")
	err := http.ListenAndServe(GetPort(), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

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

// STEP 1: create a new template
//  // http://golang.org/pkg/html/template/#New *** OR THIS ONE? *** http://golang.org/pkg/html/template/#Template.New
// STEP 2: parse the string into the template
//  // in lay terms: "give the template your form letter"
//  // in lay terms: "put your form letter into the template"
//  // http://golang.org/pkg/html/template/#Template.Parse
//  // http://golang.org/pkg/html/template/#Must
var upperTemplate = template.Must(template.New("showimage").Parse(upperTemplateHTML))

func showimage(w http.ResponseWriter, r *http.Request) {
	// Sample address "1600 Amphitheatre Parkway, Mountain View, CA"
	addr := r.FormValue("str")

	// QueryEscape escapes the addr string so
	// it can be safely placed inside a URL query
	// safeAddr := url.QueryEscape(addr)
	safeAddr := url.QueryEscape(addr)
	fullUrl := fmt.Sprintf(
		"http://maps.googleapis.com/maps/api/geocode/json?sensor=false&address=%s",
		safeAddr)

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	client := &http.Client{}

	// Build the request
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
	}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, requestErr := client.Do(req)
	if requestErr != nil {
		log.Fatal("Do: ", requestErr)
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
	var res2 GeoCodeJson
	// https://developers.google.com/maps/documentation/geocoding/#JSON
	// making a map
	// with a key of [string]
	// and a value of a slice of maps: []map
	// and the slice of maps will have THREE more maps in it
	//   // map 1: [string]map
	//   // map 2: [string]map
	//   // map 3: [string]interface{}
	//   // and the initial size of the map is 0
	//   // http://golang.org/pkg/builtin/#make

	/* REMINDERS

	   reminder from samples/54...
	   	// Maps - Shorthand Notation
	   	myGreeting := map[string]string{
	   		"Tim":     "Good morning!",
	   		"Jenny":   "Bonjour!",
	   }

	   reminder from samples/60...
	   mySlice := []int{1, 5, 15, 20, 25, 30}

	   reminder from samples/85...
	   // (1) - declare a type
	   // (2) - say that it's an interface
	   // (3) - declare the method(s)
	   type renamable interface {
	   	rename(newName string)
	   }

	*/

	// We will be using the Unmarshal function
	// to transform our JSON bytes into the
	// appropriate structure.
	// The Unmarshal function accepts a byte array
	// and a reference to the object which shall be
	// filled with the JSON data (this is simplifying,
	// it actually accepts an interface)
	json.Unmarshal(body, &res)
	json.Unmarshal(body, &res2)


//	lat := res["results"][0]["geometry"]["location"]["lat"]
//	lng := res["results"][0]["geometry"]["location"]["lng"]
	lat := res2.Results[0].Geometry.Location.Lat
	lng := res2.Results[0].Geometry.Location.Lng

	// %.13f is used to convert float64 to a string
	// https://gobyexample.com/string-formatting
	queryUrl :=
		fmt.Sprintf(
//			"http://maps.googleapis.com/maps/api/streetview?sensor=false&size=600x300&location=%.13f,%.13f", lat, lng)
//			"http://maps.googleapis.com/maps/api/staticmap?zoom=20&size=640x640&center=%.13f,%.13f&maptype=satellite", lat, lng)
			"http://maps.googleapis.com/maps/api/staticmap?size=640x640&markers=color:red%%7Alabel:A%%7C%.13f,%.13f&maptype=satellite", lat, lng)

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
