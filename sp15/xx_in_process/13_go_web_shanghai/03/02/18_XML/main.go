package main

import (
    "encoding/xml"
    "net/http"
)

type Profile struct {
    Name    string
    Hobbies []string `xml:"Hobbies>Hobby"`
}

func main() {
    http.HandleFunc("/", foo)
    http.ListenAndServe(":3000", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
    profile := Profile{"Alex", []string{"snowboarding", "programming"}}

    x, err := xml.MarshalIndent(profile, "", "  ")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/xml")
    w.Write(x)
}

/*
$ curl -i localhost:3000

HTTP/1.1 200 OK
Content-Type: application/xml
Content-Length: 128

<Profile>
    <Name>Alex</Name>
    <Hobbies>
        <Hobby>snowboarding</Hobby>
        <Hobby>programming</Hobby>
    </Hobbies>
</Profile>

*/

// http://www.alexedwards.net/blog/golang-response-snippets
