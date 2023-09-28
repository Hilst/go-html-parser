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

type TestRequest struct {
	LayoutHTML string         `json:"html"`
	Data       map[string]any `json:"data"`
}

func NewController(builder *b.TemplateBuilder, service *s.Service) *Controller {
	return &Controller{
		builder,
		service,
	}
}

func (c *Controller) get_layout_layoutname(ctx *gin.Context) {
	layoutName := ctx.Param("layoutname")
	layouts := c.service.RequestLayout(layoutName)
	data := c.service.RequestData("zero.json")
	htmls, errs := c.builder.Build(layouts, data)
	status := http.StatusOK
	readyStatus(&status, errs, len(layouts))
	response := LayoutResponse{
		LayoutsHtmls: htmls,
		Errors:       errs,
	}
	ctx.JSON(status, response)
}

func (c *Controller) patch_layout_test(ctx *gin.Context) {
	var body *TestRequest
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.Data(http.StatusInternalServerError, "text/plain", []byte(err.Error()))
	}

	htmls, errs := c.builder.Build([]string{body.LayoutHTML}, body.Data)
	if len(htmls) != 1 && len(errs) != 1 {
		ctx.Data(http.StatusInternalServerError, "", []byte{})
	}
	status := http.StatusOK
	readyStatus(&status, errs, 1)
	if status == http.StatusOK {
		ctx.Data(status, "text/html", []byte(htmls[0]))
	} else {
		ctx.Data(status, "text/plain", []byte(errs[0].Error()))
	}
}

func (c *Controller) Main() {
	router := gin.Default()

	router.GET("/layout/:layoutname", c.get_layout_layoutname)
	router.PATCH("/layout/test", c.patch_layout_test)

	router.Run(":8080")
}
