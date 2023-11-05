package templates

import (
	html "html/template"
)

type TemplateService struct {
	parsed      *html.Template
	execute     html.Template
	entryLayout string
}

func NewTemplateService() *TemplateService {
	ts := &TemplateService{}
	return ts
}

func (ts *TemplateService) Ready() {
	ts.addInternalFuncs()
	ts.parseALL()
	ts.execute = *html.Must(ts.parsed.Clone())
}

func (ts *TemplateService) addInternalFuncs() {
	AddFunction("child", ts.Retrieve)
	AddFunction("loadchild", ts.Load)
	AddFunction("clearchildren", Restart)
	AddFunction("self", func() html.HTML { return html.HTML(ts.entryLayout) })
}

func (ts *TemplateService) parseALL() {
	ts.parsed = html.Must(html.New("ALL").Funcs(FuncMap()).ParseGlob("./res/templates/**/*.html"))
}

func (ts *TemplateService) ParseLayout(layoutTemplate string) {
	ts.entryLayout = layoutTemplate
	comps := html.Must(ts.parsed.Clone())
	ts.execute = *html.Must(comps.New("LAYOUT").Funcs(FuncMap()).Parse(layoutTemplate))
}

func (ts *TemplateService) GetTemplate() *html.Template {
	return &ts.execute
}
