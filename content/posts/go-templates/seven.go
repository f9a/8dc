package main

import (
	"fmt"
	"html/template"
	"os"
)

var littleLamb = `Mary had a little {{colorize "lamb"}}`

type size struct {
	Size string
}

func main() {
	funcs := template.FuncMap{
		"colorize": func(stuff string) string {
			return fmt.Sprintf("red %v", stuff)
		},
	}

	rhyme := template.New("mary had a little lamb").Funcs(funcs)

	_, err := rhyme.Parse(littleLamb)
	if err != nil {
		panic(err)
	}

	err = rhyme.Execute(os.Stdout, nil)
	if err != nil {
		panic(err)
	}

}
