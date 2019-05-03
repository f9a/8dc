package main

import (
	"html/template"
	"os"
)

var littleLamb = `
What Mary had already:
{{range .}}> {{.}}
{{end}}
`

type size struct {
	Size string
}

func main() {
	rhyme := template.New("mary had a little lamb")

	_, err := rhyme.Parse(littleLamb)
	if err != nil {
		panic(err)
	}

	context := []string{"lamb #1", "lamb #2"}

	err = rhyme.Execute(os.Stdout, context)
	if err != nil {
		panic(err)
	}

}
