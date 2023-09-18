package main

import (
	"encoding/json"
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
