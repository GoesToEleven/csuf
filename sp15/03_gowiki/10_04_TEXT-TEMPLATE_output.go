package main

import (
	"fmt"
	"os"
	"text/template"
)

type inventory struct {
	Material string
	Count    uint
}

func main() {
	const letter = `{{.Count}} items are made of {{.Material}}`
	// const letter string = `{{.Count}} items are made of {{.Material}}`
	// const letter = '{{.Count}} items are made of {{.Material}}'
	// const letter = "{{.Count}} items are made of {{.Material}}"

	sweaters := inventory{"wool", 17}

	// Create a new template and parse the letter into it.
	tmpl := template.Must(template.New("dog").Parse(letter))

	err := tmpl.Execute(os.Stdout, sweaters)

	fmt.Println()
	if err != nil {
		panic(err)
	}

	/*





















	*/
	// CONSTANTS can't use := but only =
	const x string = "Hello World"
	// const x := "Hello World"
	fmt.Println(x)

}
