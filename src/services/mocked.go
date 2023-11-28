package service

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	mdl "github.com/Hilst/go-ui-html-template/models"
	"github.com/Hilst/go-ui-html-template/services/env"
)

type mocked struct {
	staticsPath string
	jsonsPath   string
}

func newMock() (*mocked, bool) {
	staticsPath, jsonsPath, ok := env.MockEnv()
	return &mocked{staticsPath, jsonsPath}, ok
}

func (m *mocked) RequestData(id string, ch chan mdl.DataResponse) {
	defer close(ch)
	abs, errAbs := filepath.Abs(filepath.Join(m.jsonsPath, id+".json"))
	if errAbs != nil {
		ch <- mdl.NewDataRespError(errAbs)
		return
	}
	data, errRead := os.ReadFile(abs)
	if errRead != nil {
		ch <- mdl.NewDataRespError(errRead)
		return
	}
	var result mdl.MiddleDataResp
	errMarshal := json.Unmarshal(data, &result)
	ch <- mdl.NewDataResp(result, errMarshal)
}

func (m *mocked) RequestLayout(layoutName string, ch chan mdl.LayoutResponse) {
	defer close(ch)
	abs, errAbs := filepath.Abs(filepath.Join(m.staticsPath, layoutName))
	if errAbs != nil {
		ch <- mdl.NewLayoutRespError(errAbs)
		return
	}
	dir, errDir := os.ReadDir(abs)
	if errDir != nil {
		ch <- mdl.NewLayoutRespError(errDir)
		return
	}

	var name string
	for _, file := range dir {
		if file.IsDir() {
			continue
		}

		name = file.Name()
		data, errRead := os.ReadFile(filepath.Join(abs, name))
		if errRead != nil {
			ch <- mdl.NewLayoutRespError(errRead)
			return
		}

		name = strings.Replace(name, ".html", "", -1)
		ch <- mdl.NewLayoutResp(string(data), name)
	}
}
