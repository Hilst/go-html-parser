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

type LayoutResponse struct {
	LayoutsHtmls []string `json:"layouts_htmls"`
	Errors       []error  `json:"errors_list"`
}

func NewController(builder *b.TemplateBuilder, service *s.Service) *Controller {
	return &Controller{
		builder,
		service,
	}
}

func (c *Controller) layout_Playoutname(ctx *gin.Context) {
	layoutName := ctx.Param("layoutname")
	layouts := c.service.RequestLayout(layoutName)
	htmls, errs := c.builder.Build(layouts)
	status := http.StatusOK
	ne := numberErrors(errs)
	if ne > 0 && ne < len(layouts) {
		status = http.StatusPartialContent
	}
	if ne == len(layouts) {
		status = http.StatusNotFound
	}
	response := LayoutResponse{
		LayoutsHtmls: htmls,
		Errors:       errs,
	}
	ctx.JSON(status, response)
}

func numberErrors(errs []error) int {
	r := 0
	for _, v := range errs {
		if v == nil {
			r++
		}
	}
	return len(errs) - r
}

func (c *Controller) Main() {
	router := gin.Default()

	router.GET("/layout/:layoutname", c.layout_Playoutname)

	router.Run(":8080")
}
