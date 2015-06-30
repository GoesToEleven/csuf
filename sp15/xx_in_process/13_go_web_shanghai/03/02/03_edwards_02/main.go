package main

import (
    "log"
    "net/http"
)


func main() {

    rh := http.RedirectHandler("http://www.google.com", 307)
    log.Println("Listening hippy ...")
    http.ListenAndServe(":3000", rh)

}

/*
from:
http://www.alexedwards.net/blog/a-recap-of-request-handling

We were able to do this because ServeMux also has a ServeHTTP method,
meaning that it too satisfies the Handler interface.

For me it simplifies things to think of a ServeMux as just being a special kind of handler,
which instead of providing a response itself passes the request on to a second handler.

This isn't as much of a leap as it first sounds –
chaining handlers together is fairly commonplace in Go.

Processing HTTP requests with Go is primarily about two things:
(1) ServeMux aka Request Router aka MultiPlexor
(2) Handlers
*/