package main

import (
	"fmt"
	"io/ioutil"
)

type page struct {
	Title string
	Body  []byte
}

func (p *page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &page{Title: title, Body: body}, nil
}

func main() {
	p1 := &page{Title: "TestPage", Body: []byte("This is a sample page.")}
	p1.save()
	p2, _ := loadPage("TestPage")
	// fmt.Println("p1", string(p1.Body))
	fmt.Println("p2", string(p2.Body))
}
