package main

import (
    "fmt"
    "net/url"
)

func main() {

    v := url.Values{}
    v.Set("name", "Ava")
    v.Add("friend", "Jess")
    v.Add("friend", "Sarah")
    v.Add("friend", "Zoe")
    // v.Encode() == "name=Ava&friend=Jess&friend=Sarah&friend=Zoe"
    fmt.Println(v.Get("name"))
    fmt.Println(v.Get("friend"))
    fmt.Println(v["friend"])
}
