package controller

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	mdl "github.com/Hilst/go-ui-html-template/models"
	s "github.com/Hilst/go-ui-html-template/services"
	t "github.com/Hilst/go-ui-html-template/services/templates"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	adapter "github.com/gwatts/gin-adapter"
	mini "github.com/tdewolff/minify/v2"
	miniCSS "github.com/tdewolff/minify/v2/css"
	miniHTML "github.com/tdewolff/minify/v2/html"
	miniJS "github.com/tdewolff/minify/v2/js"
	miniJSON "github.com/tdewolff/minify/v2/json"
	miniSVG "github.com/tdewolff/minify/v2/svg"
	miniXML "github.com/tdewolff/minify/v2/xml"
)

type Controller struct {
	service *s.Service
	ts      *t.TemplateService
	m       *mini.M
}

type TestRequest struct {
	LayoutHTML string         `json:"html"`
	Data       map[string]any `json:"data"`
}

func NewController(service *s.Service, ts *t.TemplateService) *Controller {
	return &Controller{
		service,
		ts,
		minifyNew(),
	}
}

func minifyNew() *mini.M {
	m := mini.New()
	m.AddFunc("text/css", miniCSS.Minify)
	m.AddFunc("text/html", miniHTML.Minify)
	m.AddFunc("image/svg+xml", miniSVG.Minify)
	m.AddFuncRegexp(regexp.MustCompile("^(application|text)/(x-)?(java|ecma)script$"), miniJS.Minify)
	m.AddFuncRegexp(regexp.MustCompile("[/+]json$"), miniJSON.Minify)
	m.AddFuncRegexp(regexp.MustCompile("[/+]xml$"), miniXML.Minify)
	return m
}

func (c *Controller) get_layout_layoutname(ctx *gin.Context) {
	layoutName := ctx.Query("name")

	layoutch := make(chan mdl.LayoutResponse)
	go c.service.RequestLayout(layoutName, layoutch)

	datach := make(chan mdl.DataResponse, 1)
	go c.service.RequestData(layoutName, datach)

	var builder strings.Builder
	var hiddenNotation string
	var s string
	i := 0
	for layout := range layoutch {
		if layout.Error != nil {
			ctx.HTML(http.StatusNotFound, "404.html", layout.Error.Error())
			return
		}
		if i == 0 {
			hiddenNotation = ""
		} else {
			hiddenNotation = "hidden"
		}
		s = fmt.Sprintf("<div id=\"page_%s\" %s>\n%s\n</div>\n", layout.Ok.Name, hiddenNotation, layout.Ok.Tmpl)
		builder.WriteString(s)
		i++
	}
	combinedLayout := builder.String()

	c.ts.ParseLayout(combinedLayout)
	ctx.HTML(http.StatusOK, "MAIN", <-datach)
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
	router.Use(adapter.Wrap(c.m.Middleware))

	router.SetHTMLTemplate(c.ts.GetTemplate())
	router.StaticFile("/", "./res/static/index.html")
	router.StaticFS("static", http.Dir("./res/static/"))

	router.GET(c.generatePath(layoutPath), c.get_layout_layoutname)
	router.PATCH(c.generatePath(layoutPath, testPath), c.patch_layout_test)

	router.Run(":8080")
}
