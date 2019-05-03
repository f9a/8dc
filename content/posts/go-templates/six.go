package main

import (
	"errors"
	"html/template"
	"os"
)

var littleLamb = `Mary had a {{.Size.Grow}} lamb {{.Size}}`

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
		Size size
	}{
		Size: size{
			Size: "big",
		},
	}

	// Übergebe einen Context in diesem Fall einen String
	err = rhyme.Execute(os.Stdout, context)
	if err != nil {
		panic(err)
	}

}

func (s size) Grow() (string, error) {
	if s.Size == "little" {
		return "big", nil
	}

	return "", errors.New("Is already big")
}
