package static

import (
	"net/http"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w,r, "public/"+r.URL.Path)
}
