package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("testCreated.txt")
	if err != nil {
		fmt.Println("There was an error", err)
		return
	}
	defer file.Close()

	file.WriteString("This file was created in a program!!!")
}
