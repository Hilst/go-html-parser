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
	case "ccd110f9-fb3f-4c9b-b4ef-52995e6477b3":
		return "total.json"
	case "f68fb80b-94b7-40f2-933b-1c0e32aade8e":
		return "zero.json"
	default:
		return "data.json"
	}
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
