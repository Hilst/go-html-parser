package main

import (
	"log"

	"github.com/Hilst/go-ui-html-template/builder"
)

func main() {
	// DECLARE SERVICE
	service := Service{
		dataPath:   "./mocks/",
		layoutPath: "./screens/",
	}
	// GET DATA
	data := service.RequestData("zero.json")
	// GET LAYOUT HTML AS STRING
	layout := service.RequestLayout("layout-items.html")

	tb := builder.NewBuilder(data, layout)
	str, err := tb.Build()
	if err != nil {
		log.Fatalln(err.Error())
	}
	println(str)
}
