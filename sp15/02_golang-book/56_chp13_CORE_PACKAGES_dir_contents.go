package main

import (
	"fmt"
	"os"
)

/*
To get the contents of a directory we use the same os.Open function but
give it a directory path instead of a file name. Then we call the Readdir method
*/

func main() {
	dir, err := os.Open(".")
	if err != nil {
		return
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return
	}
	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}
}
