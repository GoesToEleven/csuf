// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// On App Engine, the framework sets up main; we should be a different package.
package moustachio

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"html/template"
	"image"
	"image/jpeg"
	_ "image/png" // import so we can read PNG files.
	"io"
	"net/http"
	"strconv"

	"appengine"
	"appengine/datastore"

	"resize"
)

var (
	templates = template.Must(template.ParseFiles(
		"edit.html",
		"error.html",
		"upload.html",
	))
)

// Because App Engine owns main and starts the HTTP service,
// we do our setup during initialization.
func init() {
	http.HandleFunc("/", uploadHandler)
	http.HandleFunc("/edit", editHandler)
	http.HandleFunc("/img", imgHandler)
}

// Image is the type used to hold the image in the datastore.
type Image struct {
	Data []byte
}

// uploadHandler is the HTTP handler for uploading images; it handles "/".
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	if r.Method != "POST" {
		// No upload; show the upload form.
		b := &bytes.Buffer{}
		if err := templates.ExecuteTemplate(b, "upload.html", nil); err != nil {
			writeError(w, r, err)
			return
		}
		b.WriteTo(w)
		return
	}

	f, _, err := r.FormFile("image")
	if err != nil {
		writeError(w, r, err)
		return
	}
	defer f.Close()

	// Grab the image data.
	var buf bytes.Buffer
	io.Copy(&buf, f)
	i, _, err := image.Decode(&buf)
	if err != nil {
		writeError(w, r, err)
		return
	}

	// Resize if too large, for more efficient moustachioing.
	// We aim for less than 1200 pixels in any dimension; if the
	// picture is larger than that, we squeeze it down to 600.
	const max = 1200
	if b := i.Bounds(); b.Dx() > max || b.Dy() > max {
		// If it's gigantic, it's more efficient to downsample first
		// and then resize; resizing will smooth out the roughness.
		if b.Dx() > 2*max || b.Dy() > 2*max {
			w, h := max, max
			if b.Dx() > b.Dy() {
				h = b.Dy() * h / b.Dx()
			} else {
				w = b.Dx() * w / b.Dy()
			}
			i = resize.Resample(i, i.Bounds(), w, h)
			b = i.Bounds()
		}
		w, h := max/2, max/2
		if b.Dx() > b.Dy() {
			h = b.Dy() * h / b.Dx()
		} else {
			w = b.Dx() * w / b.Dy()
		}
		i = resize.Resize(i, i.Bounds(), w, h)
	}

	// Encode as a new JPEG image.
	buf.Reset()
	if err := jpeg.Encode(&buf, i, nil); err != nil {
		writeError(w, r, err)
		return
	}

	// Create an App Engine context for the client's request.
	c := appengine.NewContext(r)

	// Save the image under a unique key, a hash of the image.
	key := datastore.NewKey(c, "Image", keyOf(buf.Bytes()), 0, nil)
	if _, err = datastore.Put(c, key, &Image{buf.Bytes()}); err != nil {
		writeError(w, r, err)
		return
	}

	// Redirect to /edit using the key.
	http.Redirect(w, r, "/edit?id="+key.StringID(), http.StatusFound)
}

// keyOf returns (part of) the SHA-1 hash of the data, as a hex string.
func keyOf(data []byte) string {
	sha := sha1.New()
	sha.Write(data)
	return fmt.Sprintf("%x", string(sha.Sum(nil))[0:8])
}

// editHandler is the HTTP handler for editing images; it handles "/edit".
func editHandler(w http.ResponseWriter, r *http.Request) {
	b := &bytes.Buffer{}
	if err := templates.ExecuteTemplate(b, "edit.html", r.FormValue("id")); err != nil {
		writeError(w, r, err)
	}
	b.WriteTo(w)
}

// imgHandler is the HTTP handler for displaying images and painting
// moustaches; it handles "/img".
func imgHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	key := datastore.NewKey(c, "Image", r.FormValue("id"), 0, nil)
	im := new(Image)
	if err := datastore.Get(c, key, im); err != nil {
		writeError(w, r, err)
		return
	}

	m, _, err := image.Decode(bytes.NewBuffer(im.Data))
	if err != nil {
		writeError(w, r, err)
		return
	}

	get := func(n string) int { // helper closure
		i, _ := strconv.Atoi(r.FormValue(n))
		return i
	}
	x, y, s, d := get("x"), get("y"), get("s"), get("d")

	if x > 0 { // only draw if coordinates provided
		m = moustache(m, x, y, s, d)
	}

	b := &bytes.Buffer{}
	if err := jpeg.Encode(w, m, nil); err != nil {
		writeError(w, r, err)
		return
	}
	w.Header().Set("Content-type", "image/jpeg")
	b.WriteTo(w)
}

// writeError renders the error in the HTTP response.
func writeError(w http.ResponseWriter, r *http.Request, err error) {
	c := appengine.NewContext(r)
	c.Errorf("Error: %v", err)
	w.WriteHeader(http.StatusInternalServerError)
	if err := templates.ExecuteTemplate(w, "error.html", err); err != nil {
		c.Errorf("templates.ExecuteTemplate: %v", err)
	}
}
