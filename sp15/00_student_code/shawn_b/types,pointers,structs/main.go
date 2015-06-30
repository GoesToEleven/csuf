package main

import (
    "fmt"
)

type Robots struct {
    manufacturer string
    rtype string
    model string
    axis int
}

func main() {
    var greeting string
    greeting = "Robot types"

    fmt.Println(greeting)

    var f_lrmate = Robots{
        manufacturer: "fanuc",
        rtype: "articulated",
        model: "LR-Mate 200iD",
        axis: 6}

    fmt.Println(f_lrmate)

    var f_delta = Robots{
        manufacturer: "fanuc",
        rtype: "parallel",
        model: "M2iA",
        axis: 4}

    var r_ptr *Robots = &f_delta
    fmt.Println(r_ptr.manufacturer, r_ptr.rtype, r_ptr.model, r_ptr.axis)
    fmt.Println(r_ptr)
    fmt.Println(*r_ptr)

    var hello = "Hello, world"
    println(hello)
    var ptr_hello *string = &hello
    println(hello, *ptr_hello)
    *ptr_hello = "Hello, New World"
    println(hello, *ptr_hello)
}

