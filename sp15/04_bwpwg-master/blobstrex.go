package blobstrex

import (
        "fmt"
        "html/template"
        "io"
        "net/http"
        "github.com/rwcarlsen/goexif/exif"
        "strconv"
        "appengine"
        "appengine/blobstore"
)

func init() {
        http.HandleFunc("/", handleRoot)
        http.HandleFunc("/serve/", handleServe)
        http.HandleFunc("/upload", handleUpload)
}

func serveError(c appengine.Context, w http.ResponseWriter, err error) {
        w.WriteHeader(http.StatusInternalServerError)
        w.Header().Set("Content-Type", "text/plain")
        io.WriteString(w, "Internal Server Error")
        c.Errorf("%v", err)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
        c := appengine.NewContext(r)
        uploadURL, err := blobstore.UploadURL(c, "/upload", nil)
        if err != nil {
                serveError(c, w, err)
                return
        }
        w.Header().Set("Content-Type", "text/html")
        err = rootTemplate.Execute(w, uploadURL)
        if err != nil {
                c.Errorf("%v", err)
        }
}

var rootTemplate = template.Must(template.New("root").Parse(rootTemplateHTML))

const rootTemplateHTML = `
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<link rel="stylesheet" href="css/upper.css">
<title>Upload your Photo</title>
</head>
<body>
<form action="{{.}}" method="POST" enctype="multipart/form-data">
Upload File: <input type="file" name="file"><br />
<input type="submit" name="submit" value="Submit">
</form></body></html>
`

func handleServe(w http.ResponseWriter, r *http.Request) {
        // Instantiate blobstore reader
        reader := blobstore.NewReader(appengine.NewContext(r), 
                                      appengine.BlobKey(r.FormValue("blobKey")))
        
        lat, lng, _ := getLatLng(reader)
        
        blobstore.Delete(appengine.NewContext(r), 
                         appengine.BlobKey(r.FormValue("blobKey")))
        
        if lat == "" {
                io.WriteString(w, "Sorry but your photo has no GeoTag information...")
                return
        }        

        s := "http://maps.googleapis.com/maps/api/staticmap?sensor=false&zoom=5
              &size=600x300&maptype=roadmap&amp;center="
        s = s + lat + "," + lng + "&markers=color:blue%7Clabel:I%7C" + lat + "," + lng

        img := "<img src='" + s + "' alt='map' />"
        fmt.Fprint(w, img)
        
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
        c := appengine.NewContext(r)
        blobs, _, err := blobstore.ParseUpload(r)
        if err != nil {
                serveError(c, w, err)
                return
        }
        file := blobs["file"]
        if len(file) == 0 {
                c.Errorf("no file uploaded")
                http.Redirect(w, r, "/", http.StatusFound)
                return
        }
        http.Redirect(w, r, "/serve/?blobKey="+string(file[0].BlobKey), 
                      http.StatusFound)
}

func getLatLng(f io.Reader) (string, string, error) {
	//f, err := os.Open(fname)
	//if err != nil {
	//	return "", "", err
	//}

	x, err := exif.Decode(f)
	//defer f.Close()
	if err != nil {
		return "", "", err
	}
	lat, _ := x.Get("GPSLatitude")
	latdeg_numer, _ := lat.Rat2(0)
	latmin_numer, _ := lat.Rat2(1)
	latsec_numer, latsec_denom := lat.Rat2(2)
	var latitude float64 = float64(latdeg_numer) +
		((float64(latmin_numer) +
			((float64(latsec_numer) /
				float64(latsec_denom)) /
				float64(60.0))) /
			float64(60.0))

	latstr := strconv.FormatFloat(latitude, 'f', 15, 64)
	latRef, _ := x.Get("GPSLatitudeRef")
	if latRef.StringVal() == "S" {
		latstr = "-" + latstr
	}

	lng, _ := x.Get("GPSLongitude")
	lngdeg_numer, _ := lng.Rat2(0)
	lngmin_numer, _ := lng.Rat2(1)
	lngsec_numer, lngsec_denom := lng.Rat2(2)
	var longitude float64 = float64(lngdeg_numer) +
		((float64(lngmin_numer) +
			((float64(lngsec_numer) /
				float64(lngsec_denom)) /
				float64(60.0))) /
			float64(60.0))

	lngstr := strconv.FormatFloat(longitude, 'f', 15, 64)
	lngRef, _ := x.Get("GPSLongitudeRef")
	if lngRef.StringVal() == "W" {
		lngstr = "-" + lngstr
	}

	return latstr, lngstr, nil
}

