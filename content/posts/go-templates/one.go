package main

import (
	"html/template"
	"os"
)

var littleLamb = `
Mary had a little lamb, little lamb,
little lamb, Mary had a little lamb
whose fleece was white as snow.
And everywhere that Mary went
Mary went, Mary went, everywhere
that Mary went
The lamb was sure to go.
`

func main() {
	// Definiere neues Template
	rhyme := template.New("mary had a little lamb")

	// Lese Template String
	_, err := rhyme.Parse(littleLamb)
	if err != nil {
		panic(err)
	}

	// Schreibe Template zu Stdout
	rhyme.Execute(os.Stdout, nil)
}
