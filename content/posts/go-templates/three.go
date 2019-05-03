package main

import (
	"html/template"
	"os"
)

var littleLamb = `
{{.Human.Name}} had a {{.Animal.Size}} {{.Animal.Type}}, {{.Animal.Size}} {{.Animal.Type}},
{{.Animal.Size}} {{.Animal.Type}}, {{.Human.Name}} had a {{.Animal.Size}} {{.Animal.Type}}
whose fleece was {{.Animal.Fleece.Color}} as {{.Animal.Fleece.Attr}}.
And everywhere that {{.Human.Name}} went
{{.Human.Name}} went, {{.Human.Name}} went, everywhere
that {{.Human.Name}} went
The {{.Animal.Type}} was sure to go.

---

Places where {{.Human.Name}} and the {{.Animal.Type}} have been:
{{.Places.eternias.Name}} - X: {{.Places.eternias.X}}, Y: {{.Places.eternias.Y}}
Total places: {{.Places.Total}}
`

type (
	human struct {
		Name string
	}

	fleece struct {
		Color string
		Attr  string
	}

	animal struct {
		Type   string
		Size   string
		Fleece fleece
	}

	location struct {
		Name string
		X    float32
		Y    float32
	}

	places map[string]location
)

func main() {
	rhyme := template.New("mary had a little lamb")

	_, err := rhyme.Parse(littleLamb)
	if err != nil {
		panic(err)
	}

	heMan := human{Name: "He-Man"}
	tiger := animal{
		Type: "Tiger",
		Size: "giant",
		Fleece: fleece{
			Color: "green",
			Attr:  "jelly",
		},
	}
	pl := places{
		"eternias": location{
			Name: "Eternias",
			X:    144.33,
			Y:    1323.32,
		},
	}

	context := struct {
		Human  human
		Animal animal
		Places places
	}{
		Human:  heMan,
		Animal: tiger,
		Places: pl,
	}

	err = rhyme.Execute(os.Stdout, context)
	if err != nil {
		panic(err)
	}

}

func (p places) Total() int {
	return len(p)
}
