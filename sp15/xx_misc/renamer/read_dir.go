package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	files, _ := ioutil.ReadDir("/Users/tm002/Desktop/01_javascript")
	fmt.Println(files)
	for _, f := range files {
		fmt.Println(f.Name())
	}
}

//  http://golang.org/pkg/os/#FileInfo
