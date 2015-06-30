package main

import (
  "fmt"
)

func Greet(name string) {
  fmt.Println("Hello, " + name)
}

func GreetNames(names []string, suffix string) {
  for _, n := range names {
    Greet(n + suffix)
  }
}

func main() {
  names := []string {
    "Mark",
    "Rachel",
    "Dylan",
    "Leo",
  }

  comm := make(chan string)

  go func() {
    GreetNames(names, " <C> ")
    comm <- "Finished greeting names concurrently..."
  }()

  GreetNames(names, " <M> ")
  fmt.Println(<- comm)
}
