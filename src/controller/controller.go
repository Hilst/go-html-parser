package controller

import (
	"net/http"
	"regexp"

	s "github.com/Hilst/go-ui-html-template/services"
	t "github.com/Hilst/go-ui-html-template/services/templates"

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

func (c *Controller) Main() {
	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(adapter.Wrap(c.m.Middleware))

	router.SetHTMLTemplate(c.ts.GetTemplate())
	router.StaticFile("/", "./res/static/index.html")
	router.StaticFS("static", http.Dir("./res/static/"))

	// FUNCTIONAL ENDPOINTS
	router.GET(c.generatePath(layoutPath), c.get_layout_layoutname)
	router.PATCH(c.generatePath(layoutPath, testPath), c.patch_layout_test)

	// WEB PAGE ENDPOINTS
	router.GET(c.generatePath(tabPath, samplePath), c.get_tab_sample)
	router.GET(c.generatePath(tabPath, testPath), c.get_tab_test)

	router.Run(":8080")
}
