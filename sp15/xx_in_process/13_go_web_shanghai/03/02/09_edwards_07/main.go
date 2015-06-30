package main

import (
    "log"
    "net/http"
    "time"
)

// closure
func timeHandler(format string) http.Handler {
    fn := func(w http.ResponseWriter, r *http.Request) {
        tm := time.Now().Format(format)
        w.Write([]byte("The time is: " + tm))
    }
    return http.HandlerFunc(fn)
}

func main() {
    mux := http.NewServeMux()

    th := timeHandler(time.RFC1123)
    mux.Handle("/time", th)

    log.Println("Listening walrus...")
    http.ListenAndServe(":3000", mux)
}


/*
Processing HTTP requests with Go is primarily about two things:
(1) ServeMux aka Request Router aka MultiPlexor
(2) Handlers

In the examples before, we hardcoded the time format in the timeHandler function.
Instead, we can pass information as variables from main() to a handler.
One way to do this: put our handler logic into a closure,
and close over the variables we want to use.

First it creates fn, an anonymous function which accesses ‐ or closes over –
the format variable forming a closure. Regardless of what we do with the closure
it will always be able to access the variables that are local to the scope it was created in
– which in this case means it'll always have access to the format variable.

Secondly our closure has the signature func(http.ResponseWriter, *http.Request).
As you may remember from earlier, this means that we can convert it into a HandlerFunc type
(so that it satisfies the Handler interface).
Our timeHandler function then returns this converted closure.

In this example we've just been passing a simple string to a handler.
But in a real-world application you could use this method to pass database connection,
template map, or any other application-level context. It's a good alternative to using
global variables, and has the added benefit of making neat self-contained handlers for testing.

You might also see this same pattern written as:

func timeHandler(format string) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    tm := time.Now().Format(format)
    w.Write([]byte("The time is: " + tm))
  })
}

Or using an implicit conversion to the HandlerFunc type on return:

func timeHandler(format string) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    tm := time.Now().Format(format)
    w.Write([]byte("The time is: " + tm))
  }
}

*/