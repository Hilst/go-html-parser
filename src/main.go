package main

import (
	"log"

	bld "github.com/Hilst/go-ui-html-template/builder"
	ctr "github.com/Hilst/go-ui-html-template/controller"
	srv "github.com/Hilst/go-ui-html-template/service"
)

func main() {
	service := srv.NewService("./res/mocks/", "./res/screens/")
	data := service.RequestData("zero.json")
	tb, err := bld.NewBuilder(data)
	if err != nil {
		log.Fatalln(err.Error())
	}
	c := ctr.NewController(tb, service)
	c.Main()
}
