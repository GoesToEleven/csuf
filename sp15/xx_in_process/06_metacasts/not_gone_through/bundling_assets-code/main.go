package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/elazarl/go-bindata-assetfs"
)

func main() {
	log.Println(AssetNames())
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		tmplName := "templates/index.html"
		data, _ := Asset(tmplName)

		tmpl, err := template.New("Index").Parse(string(data))
		if err != nil {
			log.Panic(err)
		}
		tmpl.Execute(res, nil)
	})
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(&assetfs.AssetFS{Asset, AssetDir, "assets"})))
	http.ListenAndServe(":3000", nil)
}
