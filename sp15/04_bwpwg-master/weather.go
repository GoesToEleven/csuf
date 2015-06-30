package main

import (
        "encoding/json"
        "fmt"
        "io/ioutil"
        "log"
        "net/http"
        "net/url"
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
        // Create a wait group to manage the goroutines.
        var waitGroup sync.WaitGroup
        
        // Perform 4 concurrent queries
        waitGroup.Add(4)
        for query := 0; query < 4; query++ {
                go Get(query, &waitGroup)
        }
        
        // Wait for all the queries to complete.
        waitGroup.Wait()
        fmt.Printf("All Queries Completed")
}

// Get is a function that is launched as a goroutine 
func Get(query int, waitGroup *sync.WaitGroup) {
        // Decrement the wait group count so the program knows this
        // has been completed once the goroutine exits.
        defer waitGroup.Done()
        
        addr := [4]string{"Pune,India", "Franklin,TN", "Sydney,Australia", "Vientiane,Lao PDR"}
        
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
         
        var f Forecast
        json.Unmarshal(fbody, &f)

        fmt.Println("The Weather at ", addr[query])
        fmt.Println("Timezone = ", f.Timezone)
        fmt.Println("Temp in Celsius = ", f.Currently.Temperature)
        fmt.Println("Summary = ", f.Currently.Summary)
}
