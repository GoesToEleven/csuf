package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	files, _ := ioutil.ReadDir("/Users/tm002/Desktop/01_javascript")
	fmt.Println(files)
	for i, f := range files {
		fmt.Println("i -------------- ", i)
		fmt.Println(f.Name())
	}
}
