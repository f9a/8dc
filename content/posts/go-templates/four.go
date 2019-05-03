package main

import (
	"flag"
	"html/template"
	"os"
)

var littleLamb = `Mary had a little lamb, little lamb,
little lamb, Mary had a little lamb
{{if lt .Age 17}}
whose fleece was white as snow.
{{else if and (gt .Age 16) (le .Age 60) }}
whose fleece was bloody as hell.
{{else}}
whose fleece was BLOODY as HELL.
{{end}}
And everywhere that Mary went
Mary went, Mary went, everywhere
that Mary went
The lamb was sure to go.
`

func main() {
	// Lese Alter aus den übergeben Argumenten
	age := flag.Int("age", 0, "age")
	flag.Parse()

	rhyme := template.New("mary had a little lamb")

	_, err := rhyme.Parse(littleLamb)
	if err != nil {
		panic(err)
	}

	context := struct {
		Age int
	}{
		Age: *age,
	}

	// Schreibe Template zu Stdout
	err = rhyme.Execute(os.Stdout, context)
	if err != nil {
		panic(err)
	}

}
