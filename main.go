package main

import (
	HTMLtmpl "html/template"
	"os"
)

func main() {
	// DECLARE SERVICE
	service := Service{
		dataPath:   "./mocks/",
		layoutPath: "./screens/",
	}
	// GET DATA JSON
	data := service.RequestData("zero.json")
	// GET LAYOUT HTML AS STRING
	layout := service.RequestLayout("layout-items.html")
	// READ BASE ALL TEMPLATES
	all := HTMLtmpl.Must(HTMLtmpl.New("ALL").ParseGlob("./templates/**/*.html"))
	// ADD LAYOUT TEMPLATE FROM STRING
	all.New("LAYOUT").Parse(layout)
	// BUILD LAYOUT INTO OUTPUT
	all.ExecuteTemplate(os.Stdout, "MAIN", data)
}
