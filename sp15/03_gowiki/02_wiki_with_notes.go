package main

import (
	"fmt"
	"io/ioutil"
)

// create a struct to hold page data
type Page struct {
	Title string
	Body  []byte
}

/*
Sometimes we need to work with strings as binary data.
To convert a string to a slice of bytes (and vice-versa) do this:

	arr := []byte("test")
	fmt.Println(arr)
	fmt.Println(string(arr))

see
.. / golang-book / 52_chp13_CORE_PACKAGES_strings_byte_slices

*/

// the Page struct stores the page in memory
// the function 'save' below saves the page to persistent (secondary) storage
// the function will write the file to disk
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

/*
the receiver is a pointer to a struct
this allows us to change the data in the struct (pass by reference)
it also allows us to only pass the memory address of the struct as opposed
to passing all of the data in the struct



ioutil.WriteFile
https://golang.org/pkg/io/ioutil/#WriteFile
click the "WriteFile" title in the above url to be taken here:
https://golang.org/src/io/ioutil/ioutil.go?s=2520:2588#L66

re: 0600
http://en.wikipedia.org/wiki/Umask
*/

// this loads the page from secondary storage to memory
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

/*
the function returns a pointer to the page ("returns a pointer to a Page literal")
Is this different from returning the memory address where the page is stored?

the &Page returns the memory location where Page is stored
the *Page means that a pointer to a memory location is being returned

you can print that return and see the memory location
or you can de-reference that return (with *) and see the value

see
.. / golang-book / 34_chp08_pointers_pass_by_reference

ioutil.WriteFile
https://golang.org/pkg/io/ioutil/#ReadFile
click the "ReadFile" title in the above url to be taken here:
https://golang.org/src/io/ioutil/ioutil.go?s=1464:1510#L39
*/

func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.save()
	p2, _ := loadPage("TestPage")
	// // see the struct as slice of bytes
	// fmt.Println(p2)
	// // see the struct as slice of bytes
	// fmt.Println(*p2)
	// // see the struct as slice of bytes
	// fmt.Println(p2.Body)
	// // access struct values
	fmt.Println(string(p2.Body))
	fmt.Println(string(p2.Title))
}

/*
Why is the &Page use the &?
It runs with and without it.
&Page means we are storing the memory location in p1

What happens if you take 'string' out from the last line?
*/
