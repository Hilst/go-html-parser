package main

import (
	"encoding/json"
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

// SERVICE
type IService interface {
	RequestData(path string) map[string]interface{}
	RequestLayout(layoutName string) string
}

type Service struct {
	dataPath   string
	layoutPath string
}

func (s Service) RequestData(path string) map[string]interface{} {
	data, err := os.ReadFile(s.dataPath + path)
	if err != nil {
		panic("CANT READ DATA FILE")
	}
	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		panic("CANT UNMARSHAL DATA")
	}
	return result
}

func (s Service) RequestLayout(layoutName string) string {
	data, err := os.ReadFile(s.layoutPath + layoutName)
	if err != nil {
		panic("CANT READ LAYOUT FILE")
	}
	layout := string(data)
	return layout
}
