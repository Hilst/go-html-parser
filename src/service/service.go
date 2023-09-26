package service

import (
	"encoding/json"
	"log"
	"os"
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

func (s *Service) RequestData(path string) map[string]any {
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

func (s *Service) RequestLayout(layoutName string) string {
	data, err := os.ReadFile(s.layoutRoot + layoutName)
	if err != nil {
		panic("CANT READ LAYOUT FILE")
	}
	layout := string(data)
	return layout
}
