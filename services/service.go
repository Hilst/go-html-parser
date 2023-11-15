package service

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	mdl "github.com/Hilst/go-ui-html-template/models"
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

func (s *Service) RequestData(id string, ch chan mdl.DataResponse) {
	defer close(ch)
	path := pathForId(id)
	data, err := os.ReadFile(s.dataRoot + path)
	if err != nil {
		ch <- mdl.NewDataResp(nil, err)
		return
	}
	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	ch <- mdl.NewDataResp(result, err)
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

func (s *Service) RequestLayout(layoutName string, ch chan mdl.LayoutResponse) {
	defer close(ch)
	abs, err := filepath.Abs(filepath.Join(s.layoutRoot, layoutName))
	if err != nil {
		ch <- mdl.NewLayoutRespError(err)
		return
	}
	dir, err := os.ReadDir(abs)
	if err != nil {
		ch <- mdl.NewLayoutRespError(err)
		return
	}

	var name string
	for _, file := range dir {
		if file.IsDir() {
			continue
		}

		name = file.Name()
		data, err := os.ReadFile(filepath.Join(abs, name))
		if err != nil {
			ch <- mdl.NewLayoutRespError(err)
			return
		}

		name = strings.Replace(name, ".html", "", -1)
		ch <- mdl.NewLayoutResp(string(data), name)
	}
}
