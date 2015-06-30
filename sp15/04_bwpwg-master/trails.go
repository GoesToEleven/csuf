package main
 
import (
	"fmt"
	"html/template"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson" // short for Binary JSON
	"log"
	"net/http"
	"os"
)

// TrailName contains information for an individual trail
type TrailName struct {
        ID            bson.ObjectId `bson:"_id,omitempty"`
        Name          string        `bson:"name"`
        LocDesc       string        `bson:"location_desc"`
        TrailFeatures string        `bson:"trail_features"`
        RoundTrip     float64       `bson:"round_trip"`
        ElevationGain int           `bson:"elevation_gain"`
        Difficulty    float64       `bson:"difficulty"`
}
 
func main() {
        http.Handle("/stylesheets/", http.StripPrefix("/stylesheets/", http.FileServer(http.Dir("stylesheets"))))
        http.HandleFunc("/", display)
	fmt.Println("listening...")
	err := http.ListenAndServe(GetPort(), nil)
	if err != nil {
	        log.Fatal("ListenAndServe: ", err)
        }
}

var displayTemplate = template.Must(template.New("display").Parse(displayTemplateHTML))

func display(w http.ResponseWriter, r *http.Request) {
        // In the open command window set the following for Heroku:
        // heroku config:set MONGOHQ_URL=mongodb://IndianGuru:password@troup.mongohq.com:10059/trails
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

	collection := sess.DB("trails").C("trail_names")
	
        result := []TrailName{}
        err = collection.Find(nil).All(&result)
        if err != nil {
                log.Fatal("Find: ", err)
        }

        displayTemplate.Execute(w, result)
}

const displayTemplateHTML = ` 
<!DOCTYPE html>
<html lang="en">
<head>

  <!-- Basic Page Needs
  ================================================== -->
  <meta charset="utf-8">
  <title>Smoky Mountain Hiking Trails Directory</title>
  <meta name="description" content="Smoky Mountain Hiking Trails Directory">
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
      <h1 class="remove-bottom" style="margin-top: 40px">Smoky Mountain Hiking Trails Directory</h1>
      <h5>Version 1.0</h5>
      <hr />
    </div>

    <div class="sixteen columns">
      <p>The table below provides basic information for some of Smoky Mountain hiking trails. This list of trails are in alphabetical order.</p>
      <p><strong>Difficulty Rating Defined</strong>: a difficulty rating of less than 5 is generally considered to be an easy hike. Between 5 and 10 is moderate, and anything over 10 is considered to be strenuous.</p>
      <p><strong>R/T</strong> on the table heading below stands for total round trip miles.</p>
      <div class="example">
        <!-- https://github.com/dstotijn/Skeleton-tables -->
        <table>
          <thead>
            <tr>
              <th>Trail</th>
              <th>Location</th>
              <th>Trail Features</th>
              <th>R/T miles</th>
              <th>Elevation Gain</th>
              <th>Difficulty Rating</th>
            </tr>
          </thead>
          <tbody>
            {{range .}}
            <tr>
              <td>{{.Name |html}}</td>
              <td>{{.LocDesc}}</td>
              <td>{{.TrailFeatures}}</td>
              <td>{{.RoundTrip}}</td>
              <td>{{.ElevationGain}}</td>
              <td>{{.Difficulty}}</td>
            </tr>
            {{ end }}
          </tbody>
        </table>
      </div>
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

