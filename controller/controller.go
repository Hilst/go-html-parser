package controller

import (
	"net/http"

	b "github.com/Hilst/go-ui-html-template/builder"
	s "github.com/Hilst/go-ui-html-template/service"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	builder *b.TemplateBuilder
	service *s.Service
}

func NewController(builder *b.TemplateBuilder, service *s.Service) *Controller {
	return &Controller{
		builder,
		service,
	}
}

func (c *Controller) Main() {
	router := gin.Default()

	router.GET("/layout/:layoutname", func(ctx *gin.Context) {
		layoutName := ctx.Param("layoutname")
		layout := c.service.RequestLayout(layoutName + ".html")
		texthtml, err := c.builder.Build(layout)
		if err != nil {
			ctx.Data(http.StatusInternalServerError, "text/plain", []byte(err.Error()))
			return
		}

		ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(texthtml))
	})

	router.Run(":8080")
}
