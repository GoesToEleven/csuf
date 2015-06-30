package msgmartini

import (
        "net/http"
        "github.com/codegangsta/martini"
)

func init() {
        m := martini.Classic()
        m.Get("/", func() string {
                return "Hello from RubyLearning.org - 
                        Learn Go with 1000s of other participants!"
        })

        http.Handle("/", m)
}
