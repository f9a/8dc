package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"
)

var littleLamb = `Mary had a little {{colorize "lamb" | lovely | toUpper}}`

type size struct {
	Size string
}

func main() {
	funcs := template.FuncMap{
		"colorize": func(stuff string) string {
			return fmt.Sprintf("red %v", stuff)
		},
		"lovely": func(stuff string) string {
			return fmt.Sprintf("lovely %v", stuff)
		},
		"toUpper": strings.ToUpper,
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
