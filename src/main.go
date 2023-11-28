package main

import (
	ctr "github.com/Hilst/go-ui-html-template/controller"
	srv "github.com/Hilst/go-ui-html-template/services"
	tmp "github.com/Hilst/go-ui-html-template/services/templates"
)

func main() {
	service := srv.NewIService()

	ts := tmp.NewTemplateService()
	ts.Ready()

	c := ctr.NewController(service, ts)
	c.Main()
}
