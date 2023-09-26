package service

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
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

func (s *Service) RequestLayout(layoutName string) []string {
	abs, err := filepath.Abs(filepath.Join(s.layoutRoot, layoutName))
	if err != nil {
		log.Fatalln(err.Error())
	}
	dir, err := os.ReadDir(abs)
	if err != nil {
		log.Fatalln(err.Error())
	}

	out := make([]string, 0)
	for _, file := range dir {
		if file.IsDir() {
			continue
		}

		data, err := os.ReadFile(filepath.Join(abs, file.Name()))
		if err != nil {
			log.Fatalln(err.Error())
		}
		out = append(out, string(data))
	}
	return out
}
