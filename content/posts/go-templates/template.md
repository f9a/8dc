---
title: Go Templates
date: 2019-05-01T23:00:00+01:00
categories:
  - go
url: /post/uuid/6e6e83fd-24a1-5ed1-b763-3eceefe50c5d
---

Ein Einführung in die Go [Template Bibliothek](https://golang.org/pkg/text/template).

Ein Beispiel sagt mehr als 1000 Worte. Starten wir daher mit einem klassischen "Hello, World", nicht, stattdessen gehen wir vorran mit ["Mary had a little lamb"](https://en.wikipedia.org/wiki/Mary_Had_a_Little_Lamb).

```golang
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
```

Fantastisch! Wir haben einen String nach Stdout geschrieben. Jetzt, lass uns etwas Sinnvolles machen. Geben wir dem Kinder Lied einen neuen Context.

## Einführung Context

```golang
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
```

Nicht nur Mary ist nun verwirrt. Was macht diese {{.}} Gebilde im Template String?

Wir haben beim erzeugen des Templates den String "enormous" übergeben

```golang
err = rhyme.Execute(os.Stdout, "enormous")
```

auf diesen können wir nun innerhalb des Templates zugreifen. Der Zugriff auf den String erfolgt über

```text
{{.}}
```

Ding die sich innerhalb {{ und }} befinden werden beim erzeugen des Templates verarbeitet.

Der Punkt repräsentiert einen Context in diesem Fall den Globalen. Wir können und das ganze vortsellen wie eine Bastelkiste in dieser Kiste können unterschiedlichste Bastelmaterialien abgelegt werden wie z.Bsp. Klebestift, Buchstaben aus Karton, ein Stempel. Mit diesen in der Kiste abgelegten Dinge können wir dem Template Farbe verleihen.

!!! Todo(ttochtermann): Bild Bastelkiste

In unserem Fall befindet sich nur ein einsamer String in der Kiste, nicht sehr spannend. Bringen wir daher mehr Fabre ins Spiel.

```golang
three.go
package main

import (
	"html/template"
	"os"
)

var littleLamb = `
Mary had a {{.Size}} lamb, {{.Size}} lamb,
{{.Size}} lamb, Mary had a {{.Size}} lamb
whose fleece was {{.Fleece.Color}} as {{.Fleece.Attr}}.
And everywhere that Mary went
Mary went, Mary went, everywhere
that Mary went
The lamb was sure to go.
`

// Struct Felder müssen öffentlich sein.
type sheep struct {
	Size   string
	Fleece struct {
		Color string
		Attr  string
	}
}

func main() {
	rhyme := template.New("mary had a little lamb")

	_, err := rhyme.Parse(littleLamb)
	if err != nil {
		panic(err)
	}

	sheep := sheep{}
	sheep.Size = "giant"
	sheep.Fleece.Color = "bloody"
	sheep.Fleece.Attr = "hell"

	context := sheep

	// Schreibe Template zu Stdout
	err = rhyme.Execute(os.Stdout, context)
	if err != nil {
		panic(err)
	}

}
```

Wir übergeben als context ein Struct auf die Felder diese Structs können wir wie folgt zugreifen.

```text
{{.Feldname}}
```

Auch ist es möglich durch verkettung auf Verschachtelte Daten zuzugreifen

```text
{{.Feldname1.Feldname2}}
```

Wie oben angemerkt befinden wir uns in diesem Fall im globalen Context, auf diesen kann ebenso mit hilfe des \$ zugegriffen werden.

```golang
{{$.}}
```

D.h. wir hätten im Beispiel oben auch mittels dem folgenden auf die übergeben Struktur zugreifen können.

```text
{{$.Feldname}}
```

Warum wir das benötigt? Der Context kann sich ändern (es wird nicht immer gebastelt) zum Beispiel innerhalb einer Schleife. Möchte man innerhalb dieser auf den gloablen Context zugreifen muss man in diesem Fall das \$ vorstellen. Später dazu mehr.

## Zugriff auf Variablen und Methods

Erweiteren wir nun unser Template um weiter Variablen und Methoden Aufrufe.

```golang
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
```

### Maps

Im Beispiel oben wird auf die places map wie folgt zugegriffen.

```
{{.Places.eternias.Name}}
```

wobei eternias der map Schlüssel ist und Name zum location Struct gehört.

Allgemein kann wie folgt mit maps gearbeitet werden

```golang
context = map[string]string{"key": "value"}
templateStr := "{{.key}}"
```

Wobei der key aus aus Buchstaben bestehen muss. Es ist möglich den Aufruf zu verketten, wie im Beispiel oben gezeigt.

### Methods

Im Beispiel wird auf die Method von places wie folgt zugegriffen.

```
{{.Places.Total}}
```

Wobei Total die Method ist. Es ist mölgich den Aufruf zu verketten. Die Methode muss mindestens einen Wert zurückgegeben, jedoch maximal zwei. Der Zweite Rückgabewert kann ein Error sein. Ist dieser nicht nil wird die Verarbeitung unterbrochen und der Fehler wird ausgegeben.

```golang
package main

import (
	"errors"
	"html/template"
	"os"
)

var littleLamb = `Mary had a {{.Size.Grow}} lamb, {{.Size}} lamb`

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
```

Die Welt ist nun schon ziemlich bunt, geht aber noch bunter.

## Funktionen

Es ist möglich in einem Template Funktionen zu verwenden. Der Aufruf dafür lautet wie folgt.

```
{{func arg1 arg2 argX}}
```

Der erste Teil ist der Funktionsname alle weiter Teile sind Argumente. Argumente können wiederum Structs, Maps, Strings, Ints, Funktionen, Methoden und Verknüpfungen dieser sein. Es gibt einige [vordefiniert Funktionen](https://golang.org/pkg/text/template/#hdr-Functions). Ebenfalls können wir weiter Funktionen für ein Template definieren.

```golang
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
```

## Piplines

In der Package Dokumentation wird viel über piplines gesprochen. Piplines ist ein Befehl oder mehere Befehle verbunden durch |. Wer sich schon eimal mit Linux oder ähnlichem beschäftigt hat wird die schon kennen. Man kan die Ausgabe von einem Befehl mittels | an den nächsten weiterleiten. Was ist ein Befehl? Ein Befehl kann sein, ein Wert (wie zum Beispiel ein String, Int), eine Funktion oder eine Methode.

```golang
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
```

Es ist möglich das ein Befehl in der Kette ein Error zurückgibt dann wird die Ausführung unterbrochen und der Error wird zurückgegeben.

## Weiters zu Varaibeln

Möchte man eine Variable innerhalb des Templates definieren ist dies ebenfalls machbar.

```
{{$var := "value"}}
```

Bringen wir nun etwas mehr Kontrolle in das Template.

## If

In der Package Dokumentantation wird der if block wie folgt spezifiziert.

{{if pipeline}} T1 {{else if pipeline}} T0 {{end}}
{{if pipeline}} T1 {{else}} T0 {{end}}
{{if pipeline}} T1 {{end}}

Wir wollen nun unterscheiden welcher Version des Liede eine Person zu sehen bekommt.

- Ist die Person unter 16 bekommt diese Person eine jugendfreundliche Version des Liedes zu sehen.
- Ist die Person zwischen 16 und 60 bekommt Sie eine die nicht Jugendfreundliche Version zu lesen.
- Der dritte Fall ist das eine Person älter als 60 Jahre ist dann werden alle unzüchtige Stellen groß geschrieben.

```golang
package main

import (
	"flag"
	"html/template"
	"os"
)

var littleLamb = `
Mary had a little lamb, little lamb,
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
```

Schauen wir uns den if Block genauer an.

```golang
{{if lt .Age 17}}
```

lt ist eine von go vordefiniert Funktion (shocking news if ist ebenfalls eine vordefiniert Funktion) welche prüft ob der Wert1 < Wert2 ist. Wenn nicht prüfe das else if Statement.

```golang
{{else if and (gt .Age 16) (le .Age 60) }}
```

Hier sind gt, le und and ebefalls vordefiniert Funktionen. gt prüft ob Wert1 > Wert2 ist, le prüft ob Wert1 <= Wert2 ist und and gibt wahr zurück wenn Wert1 und Wert2 wahr sind ansonsten unwahr. Da and nur zwei Argumente erwaret muss mittels ( und ) die Funktionsaufrufe für gt und le gruppiert werden. Das else if Statment könnte auch noch wie folgt weiter gruppiert werden.

```golang
{{else if (and (gt .Age 16) (le .Age 60)) }}
```

#### Tip

Bei der Ausgabe werden nun Leerzeilen ausgegeben dort wo ein if Block ist, diese kann vermieden werden in dem man die folgende Schreibweise verwendet.

```golang
{{- if pipeline -}}
```

## range

range ist das neue for. Mit rang kann über array, slices, maps und channels iteriert werden.

Schauen wir uns an was Mary den schon alles hatte.

```golang
package main

import (
	"html/template"
	"os"
)

var littleLamb = `
What mary had already:
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
```

2 Punkte? Hier kommte es zu dem angesprochenden Context wechsel innerhalb des range Blocks wird der "Context Punkt" auf das aktuelle Element gesetzt über das gerade iteriert wird. Möchten wir nun auf eine Wert zugreifen der im globalen Context abgelget ist können wir das mittels \$.

```golang
package main

import (
	"html/template"
	"os"
)

var littleLamb = `
What Mary had already:
{{range .Lambs}}
{{$.ListStyle}} {{.}}
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
		ListStyle string
		Lambs     []string
	}{
		ListStyle: "*",
		Lambs:     []string{"lamb #1", "lamb #2"},
	}

	err = rhyme.Execute(os.Stdout, context)
	if err != nil {
		panic(err)
	}

}
```

Schauen wir ob Mary auch schon kleine Katzen hatte.

```golang
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
```
