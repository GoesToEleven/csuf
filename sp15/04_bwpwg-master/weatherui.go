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
	"sync"
)

type DataPoint struct {
	Time                   float64
	Summary                string
	Icon                   string
	SunriseTime            float64
	SunsetTime             float64
	PrecipIntensity        float64
	PrecipIntensityMax     float64
	PrecipIntensityMaxTime float64
	PrecipProbability      float64
	PrecipType             string
	PrecipAccumulation     float64
	Temperature            float64
	TemperatureMin         float64
	TemperatureMinTime     float64
	TemperatureMax         float64
	TemperatureMaxTime     float64
	DewPoint               float64
	WindSpeed              float64
	WindBearing            float64
	CloudCover             float64
	Humidity               float64
	Pressure               float64
	Visibility             float64
	Ozone                  float64
}

type Forecast struct {
	Latitude  float64
	Longitude float64
	Timezone  string
	Offset    float64
	Currently DataPoint
	Junk      string
}

func main() {
        http.Handle("/stylesheets/", http.StripPrefix("/stylesheets/", http.FileServer(http.Dir("stylesheets"))))
        http.HandleFunc("/", handler)
        http.HandleFunc("/display", display)
	fmt.Println("listening...")
	err := http.ListenAndServe(GetPort(), nil)
	if err != nil {
	        log.Fatal("ListenAndServe: ", err)
        }
}

func handler(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, rootForm)
}

const rootForm = `
<!DOCTYPE html>
<html lang="en">
<head>

  <!-- Basic Page Needs
  ================================================== -->
  <meta charset="utf-8">
  <title>Weather Forecast</title>
  <meta name="description" content="Weather Forecast">
  <meta name="author" content="Satish Talim">

  <!-- Mobile Specific Metas
  ================================================== -->
  <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">

  <!-- CSS
  ================================================== -->
  <link rel="stylesheet" href="stylesheets/base.css">
  <link rel="stylesheet" href="stylesheets/skeleton.css">
  <link rel="stylesheet" href="stylesheets/layout.css">
  <link rel="stylesheet" href="stylesheets/table.css">

  <!-- Favicons
  ================================================== -->
  <link rel="shortcut icon" href="images/favicon.ico">
  <link rel="apple-touch-icon" href="images/apple-touch-icon.png">
  <link rel="apple-touch-icon" sizes="72x72" href="images/apple-touch-icon-72x72.png">
  <link rel="apple-touch-icon" sizes="114x114" href="images/apple-touch-icon-114x114.png">

  <!-- Validation
  ================================================== -->
  <script>
  function validateForm()
  {
    var c1=document.forms["myForm"]["city1"].value;
    var c2=document.forms["myForm"]["city2"].value;
    var c3=document.forms["myForm"]["city3"].value;
    var c4=document.forms["myForm"]["city4"].value;
    if ((c1==null || c1=="") || (c2==null || c2=="") ||
        (c3==null || c3=="") || (c4==null || c4==""))
    {
      alert("City Name must be filled out");
      return false;
    }
  }
  </script>

</head>
<body>

  <!-- Primary Page Layout
  ================================================== -->
  <div class="container">
    <div class="sixteen columns">
      <h1 class="remove-bottom" style="margin-top: 40px">Weather Forecast</h1>
      <h5>Version 1.0</h5>
      <hr />
    </div>

    <div class="sixteen columns">
      <p>Please enter the cities for which you want a weather forecast.</p>
      <form name ="myForm" action="/display" onsubmit="return validateForm()" method="post" accept-charset="utf-8">
        <!-- Label and text input -->
        <label for="regularInput1">City Name 1</label>
        <input type="text" name="city1" id="regularInput1" />
        <label for="regularInput2">City Name 2</label>
        <input type="text" name="city2" id="regularInput2" />
        <label for="regularInput3">City Name 3</label>
        <input type="text" name="city3" id="regularInput3" />
        <label for="regularInput4">City Name 4</label>
        <input type="text" name="city4" id="regularInput4" />
        <button type="submit">Submit Form</button>
      </form>
    </div>
  </div><!-- container -->

<!-- End Document
================================================== -->
</body>
</html>
`

var displayTemplate = template.Must(template.New("display").Parse(displayTemplateHTML))

func display(w http.ResponseWriter, r *http.Request) {
        addr := []string{r.FormValue("city1"), r.FormValue("city2"), r.FormValue("city3"), r.FormValue("city4")}
       
        f := make([]Forecast, 4, 4)

        // Create a wait group to manage the goroutines
        var waitGroup sync.WaitGroup
        
        // Perform 4 concurrent queries
	waitGroup.Add(4)
	
        // Perform 4 concurrent queries
        for query := 0; query < 4; query++ {
                go Get(query, &waitGroup, addr, f)
        }
        
        // Wait for all the queries to complete.
        waitGroup.Wait()

        displayTemplate.Execute(w, f)  
}

// Get is a function that is launched as a goroutine 
func Get(query int, waitGroup *sync.WaitGroup, addr []string, f []Forecast) {
        // Decrement the wait group count so the program knows this
        // has been completed once the goroutine exits.
        defer waitGroup.Done()
        
        // Geocoding API
	// QueryEscape escapes the addr string so
	// it can be safely placed inside a URL query
	// safeAddr := url.QueryEscape(addr)
        safeAddr := url.QueryEscape(addr[query])
        fullUrl := fmt.Sprintf("http://maps.googleapis.com/maps/api/geocode/json?sensor=false&address=%s", safeAddr)

	// Build the request
	req, err1 := http.NewRequest("GET", fullUrl, nil)
	if err1 != nil {
		log.Fatal("NewRequest: ", err1)
	}

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err2 := client.Do(req)
	if err2 != nil {
		log.Fatal("Do: ", err2)
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
        
        // lat, lng as float64
	lat, _ := res["results"][0]["geometry"]["location"]["lat"]
	lng, _ := res["results"][0]["geometry"]["location"]["lng"]

        
        // Forecast API
	// %.13f is used to convert float64 to a string
	url := fmt.Sprintf("https://api.forecast.io/forecast/yourapikey/%.13f,%.13f?units=ca", lat, lng)

        resp, err := http.Get(url)
        if err != nil {
                log.Fatal("Get: ", err)
        }
        defer resp.Body.Close()
        fbody, err := ioutil.ReadAll(resp.Body)
        if err != nil {
                log.Fatal("ReadAll: ", err)
        }
        
        json.Unmarshal(fbody, &f[query])
}

const displayTemplateHTML = ` 
<!DOCTYPE html>
<html lang="en">
<head>

  <!-- Basic Page Needs
  ================================================== -->
  <meta charset="utf-8">
  <title>Weather Forecast</title>
  <meta name="description" content="Weather Forecast">
  <meta name="author" content="Satish Talim">

  <!-- Mobile Specific Metas
  ================================================== -->
  <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">

  <!-- CSS
  ================================================== -->
  <link rel="stylesheet" href="stylesheets/base.css">
  <link rel="stylesheet" href="stylesheets/skeleton.css">
  <link rel="stylesheet" href="stylesheets/layout.css">
  <link rel="stylesheet" href="stylesheets/table.css">

  <!-- Favicons
  ================================================== -->
  <link rel="shortcut icon" href="images/favicon.ico">
  <link rel="apple-touch-icon" href="images/apple-touch-icon.png">
  <link rel="apple-touch-icon" sizes="72x72" href="images/apple-touch-icon-72x72.png">
  <link rel="apple-touch-icon" sizes="114x114" href="images/apple-touch-icon-114x114.png">

</head>
<body>

  <!-- Primary Page Layout
  ================================================== -->
  <div class="container">
    <div class="sixteen columns">
      <h1 class="remove-bottom" style="margin-top: 40px">Weather Forecast</h1>
      <h5>Version 1.0</h5>
      <hr />
    </div>

    <div class="sixteen columns">
      <p>The table below provides current weather forecast for the cities you entered.</p>
      <div class="example">
        <!-- https://github.com/dstotijn/Skeleton-tables -->
        <table>
          <thead>
            <tr>
              <th>Timezone</th>
              <th>Temperature</th>
              <th>Summary</th>
            </tr>
          </thead>
          <tbody>
            {{range .}}
            <tr>
              <td>{{.Timezone}}</td>
              <td>{{.Currently.Temperature}}</td>
              <td>{{.Currently.Summary}}</td>
            </tr>
            {{ end }}
          </tbody>
        </table>
      </div>
      <p><a href="/">Start again!</a></p>
    </div>
  </div><!-- container -->

<!-- End Document
================================================== -->
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

