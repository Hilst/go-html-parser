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

func (s *Service) RequestData(id string) mdl.DataResponse {
	data, err := os.ReadFile(s.dataRoot + id + ".json")
	if err != nil {
		return mdl.NewDataRespError(err)
	}
	var result mdl.MiddleDataResp
	err = json.Unmarshal(data, &result)
	return mdl.NewDataResp(result, err)
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
