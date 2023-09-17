package main

import (
	HTMLtmpl "html/template"
	"os"
)

type (
	Data struct {
		Items []Item
	}
	Item struct {
		Title       string
		Description string
		Type        string
	}
)

func main() {
	data := Data{
		[]Item{
			{"A", "a", "primary"},
			{"B", "b", "secondary"},
			{"C", "c", "cancel"},
		},
	}
	all := HTMLtmpl.Must(HTMLtmpl.New("ALL").ParseGlob("./templates/**/*.html"))
	println(all.DefinedTemplates())
	all.ExecuteTemplate(os.Stdout, "MAIN", data)
}
