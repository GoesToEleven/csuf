package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("testReadFrom.txt")
	if err != nil {
		fmt.Println("There was an error", err)
		return
	}
	defer file.Close()

	// get the file size
	stat, err := file.Stat()
	if err != nil {
		return
	}
	// read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}

	str := string(bs)
	fmt.Println(str)
}
