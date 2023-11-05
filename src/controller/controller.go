package controller

import (
	"fmt"
	"net/http"
	"strings"

	s "github.com/Hilst/go-ui-html-template/services"
	t "github.com/Hilst/go-ui-html-template/services/templates"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	m "github.com/tdewolff/minify/v2"
	mHTML "github.com/tdewolff/minify/v2/html"
)

type Controller struct {
	service *s.Service
	ts      *t.TemplateService
	m       *m.M
}

type TestRequest struct {
	LayoutHTML string         `json:"html"`
	Data       map[string]any `json:"data"`
}

func NewController(service *s.Service, ts *t.TemplateService) *Controller {
	m := m.New()
	m.AddFunc("text/html", mHTML.Minify)
	return &Controller{
		service,
		ts,
		m,
	}
}

func (c *Controller) get_index(ctx *gin.Context) {
	ctx.Data(http.StatusOK, "text/plain", []byte("OK"))
}

func (c *Controller) get_layout_layoutname(ctx *gin.Context) {
	layoutName := ctx.Param(c.clearVariablePath(nameVariablePath))
	layouts := c.service.RequestLayout(layoutName)

	data := c.service.RequestData(layoutName)

	var builder strings.Builder
	for _, layout := range layouts {
		builder.WriteString(fmt.Sprintf("<div id=\"page_%s\">%s</div>", layout.Name, layout.Tmpl))
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
	router.Use(c.ginMinifyHTML())

	router.GET(c.generatePath(), c.get_index)
	router.GET(c.generatePath(layoutPath, nameVariablePath), c.get_layout_layoutname)
	router.PATCH(c.generatePath(layoutPath, testPath), c.patch_layout_test)

	router.Run(":8080")
}
