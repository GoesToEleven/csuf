package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
    myPath := "/Users/tm002/Desktop/course_screenshots/lynda_unix/"
	files, _ := ioutil.ReadDir(myPath)
	for i, f := range files {
		fmt.Println(f.Name())
		oldfile := (myPath + f.Name())
		newfile := (myPath + strconv.Itoa(i) + ".png")
		//        fmt.Println(oldfile)
		//        fmt.Println(newfile)
		err := os.Rename(oldfile, newfile)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

/*
    age := 12

    // will NOT display properly
    fmt.Println(":", string(age))

    // convert int to string
    // will display properly
    fmt.Println(":", strconv.Itoa(age))
}
*/
