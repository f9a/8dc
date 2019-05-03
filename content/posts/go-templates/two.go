package main

import (
	"html/template"
	"os"
)

var littleLamb = `Mary had a {{.}} lamb, {{.}} lamb`

func main() {
	rhyme := template.New("mary had a little lamb")

	_, err := rhyme.Parse(littleLamb)
	if err != nil {
		panic(err)
	}

	// Übergebe einen Context in diesem Fall einen String
	err = rhyme.Execute(os.Stdout, "enormous")
	if err != nil {
		panic(err)
	}

}
