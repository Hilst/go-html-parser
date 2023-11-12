package controller

import (
	"fmt"
	"net/http"
	"strings"

	s "github.com/Hilst/go-ui-html-template/services"
	t "github.com/Hilst/go-ui-html-template/services/templates"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	service *s.Service
	ts      *t.TemplateService
}

type TestRequest struct {
	LayoutHTML string         `json:"html"`
	Data       map[string]any `json:"data"`
}

func NewController(service *s.Service, ts *t.TemplateService) *Controller {
	return &Controller{
		service,
		ts,
	}
}

func (c *Controller) get_layout_layoutname(ctx *gin.Context) {
	layoutName := ctx.Query("name")
	layouts := c.service.RequestLayout(layoutName)

	data := c.service.RequestData(layoutName)

	var builder strings.Builder
	var hiddenNotation string
	var s string
	for i, layout := range layouts {
		if i == 0 {
			hiddenNotation = ""
		} else {
			hiddenNotation = "hidden"
		}
		s = fmt.Sprintf("<div id=\"page_%s\" %s>\n%s\n</div>\n", layout.Name, hiddenNotation, layout.Tmpl)
		builder.WriteString(s)
	}
	combinedLayout := builder.String()

	c.ts.ParseLayout(combinedLayout)
	ctx.HTML(http.StatusOK, "MAIN", data)
}

func (c *Controller) patch_layout_test(ctx *gin.Context) {
	var body *TestRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.ts.ParseLayout(body.LayoutHTML)
	ctx.HTML(http.StatusOK, "MAIN", body.Data)
}

func (c *Controller) Main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.Use(gin.Recovery())

	router.SetHTMLTemplate(c.ts.GetTemplate())
	router.StaticFile("/", "./res/static/index.html")
	router.StaticFS("static", http.Dir("./res/static/"))

	router.GET(c.generatePath(layoutPath), c.get_layout_layoutname)
	router.PATCH(c.generatePath(layoutPath, testPath), c.patch_layout_test)

	router.Run(":8080")
}
