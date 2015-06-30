package main

import (
	"fmt"
	"os"
	"text/template"
)

type Person struct {
	fname     string
	lname     string
	honorific string
	attended  bool
	donated   bool
}

func CreatePeople() []Person {
	return []Person{{"Brandy", "Dihel", "Mrs.", true, true},
		{"Alex", "Taylor", "Mr.", true, true},
		{"James", "Bengtson", "Mr.", false, false},
		{"Megan", "Newkirk", "Mrs.", false, true},
		{"Taylor", "Dihel", "Mr.", true, false},
	}
}

type LetterInput struct {
	Fname     string
	Lname     string
	Honorific string
	Attended  bool
	Donated   bool
	Events    []string
}

func (l *LetterInput) SetPerson(p Person) {
	l.Fname = p.fname
	l.Lname = p.lname
	l.Honorific = p.honorific
	l.Attended = p.attended
	l.Donated = p.donated
}

const letter string = `
Dear {{.Honorific}} {{.Lname}},
{{if .Attended}}
Thank you for attending the Fundraiser; we hope that you had a good time.{{else}}
We are sorry that you were unable to attend the Fundraiser.{{end}}
{{if .Donated}}
Also, thank you very much for your generous donation.{{end}}

We have a few upcoming events that we hope you will keep in mind:

{{range .Events}}-{{.}}
{{end}}

Best Wishes,
Corey
`

func main() {
	var myinput = LetterInput{Events: []string{"Lemoore High School Varsity Golf Car Wash", "Relay For Life Event at the Lemoore High School Stadium", "Young Farmers of America Fishing Trip to Bass Lake"}}
	p := CreatePeople()
	t := template.Must(template.New("letter").Parse(letter))
	for _, input := range p {
		f, err := os.Create(input.fname + "_" + input.lname + ".txt")
		if err != nil {
			fmt.Println(err)
		}
		myinput.SetPerson(input)
		err = t.Execute(f, myinput)
		if err != nil {
			fmt.Println(err)
		}
		f.Close()
	}
}
