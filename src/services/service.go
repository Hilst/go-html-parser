package service

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Service struct {
	dataRoot   string
	layoutRoot string
}

func NewService(dataRoot string, layoutRoot string) *Service {
	return &Service{
		dataRoot,
		layoutRoot,
	}
}

func (s *Service) RequestData(id string) map[string]any {
	path := pathForId(id)
	data, err := os.ReadFile(s.dataRoot + path)
	if err != nil {
		log.Fatalln(err.Error())
	}
	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return result
}

func pathForId(id string) string {
	switch id {
	case "big":
		return "total.json"
	case "comp":
		return "zero.json"
	default:
		return "data.json"
	}
}

type LayoutResponse struct {
	Tmpl string
	Name string
}

func (s *Service) RequestLayout(layoutName string) []LayoutResponse {
	abs, err := filepath.Abs(filepath.Join(s.layoutRoot, layoutName))
	if err != nil {
		log.Fatalln(err.Error())
	}
	dir, err := os.ReadDir(abs)
	if err != nil {
		log.Fatalln(err.Error())
	}

	out := make([]LayoutResponse, 0)
	var name string
	for _, file := range dir {
		if file.IsDir() {
			continue
		}

		name = file.Name()
		data, err := os.ReadFile(filepath.Join(abs, name))
		if err != nil {
			log.Fatalln(err.Error())
		}

		name = strings.Replace(name, ".html", "", -1)
		layoutResponse := LayoutResponse{
			Tmpl: string(data),
			Name: name,
		}
		out = append(out, layoutResponse)
	}
	return out
}
