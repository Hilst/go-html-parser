package service

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	mdl "github.com/Hilst/go-ui-html-template/models"
	"github.com/redis/go-redis/v9"
)

type Service struct {
	redis      redis.Client
	layoutRoot string
}

func NewService(layoutRoot string) *Service {
	rClient := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	return &Service{
		*rClient,
		layoutRoot,
	}
}

func (s *Service) RequestData(id string) mdl.DataResponse {
	ctx := context.Background()
	val, err := s.redis.Get(ctx, id).Result()
	if err != nil {
		return mdl.NewDataRespError(err)
	}
	var result mdl.MiddleDataResp
	err = json.Unmarshal([]byte(val), &result)
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
