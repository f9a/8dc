package main

import (
	"html/template"
	"os"
)

var littleLamb = `
What Mary had already:
{{range .Kitten}}
{{.}}
{{else}}
I hate kitten! {{.Curse}}
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

	context := struct {
		Curse  string
		Kitten []string
	}{
		Curse:  "f#$k",
		Kitten: []string{},
	}

	err = rhyme.Execute(os.Stdout, context)
	if err != nil {
		panic(err)
	}

}
