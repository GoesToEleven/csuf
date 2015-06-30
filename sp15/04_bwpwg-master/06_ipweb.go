package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// A struct is a collection of fields
// This is our type which matches the JSON object.
type IpRecord struct {
	// These two fields use the json: tag to specify which field they map to
	CountryName string `json:"country_name"`
	CountryCode string `json:"country_code"`
	// These fields are mapped directly by name (note the different case)
	City string
	Ip   string
	// As these fields can be nullable, we use a pointer to a string rather than a string
	Lat *string
	Lng *string
}

func main() {
	http.HandleFunc("/", myHandlerFunc)
	fmt.Println("listening...")
	err := http.ListenAndServe(GetPort(), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
		return
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

func myHandlerFunc(w http.ResponseWriter, r *http.Request) {
	baseURL := "http://api.hostip.info/get_json.php?position=true&ip="
	ip := "198.252.210.32"

	// QueryEscape escapes the ip string so
	// it can be safely placed inside a URL query
	// safeIp := url.QueryEscape(ip)
	// url := baseURL + safeIp

	url := baseURL + ip

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	// Read the content into a byte array
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("ReadAll: ", err)
		return
	}
	fmt.Printf("%s", body)

	// We will be using the Unmarshal function
	// to transform our JSON bytes into the
	// appropriate structure.
	// The Unmarshal function accepts a byte array
	// and a reference to the object which shall be
	// filled with the JSON data (this is simplifying,
	// it actually accepts an interface)

	// Fill the record with the data from the JSON
	var record IpRecord
	err = json.Unmarshal(body, &record)
	if err != nil {
		// An error occurred while converting
		// our JSON to an object
		log.Fatal("Unmarshal: ", err)
	}

	fmt.Fprintf(w, "Latitude = %s and Longitude = %s", *record.Lat, *record.Lng)
	fmt.Println()
	fmt.Fprintf(w, "%+v", record)
	fmt.Fprintln(w, record.CountryName)
	fmt.Fprintln(w, record.CountryCode)
	fmt.Fprintln(w, record.City)
	fmt.Fprintln(w, record.Ip)
	fmt.Fprintln(w, *record.Lat)
	fmt.Fprintln(w, *record.Lng)
}
