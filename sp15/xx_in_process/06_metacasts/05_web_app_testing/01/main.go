package main

import (
	"net/http"

	"github.com/goestoeleven/mc/05_web_app_testing/01/my_app"
)

func main() {
	http.ListenAndServe(":3000", my_app.NewMux())
}
