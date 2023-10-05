package builder

import (
	"bytes"
	html "html/template"

	fn "github.com/Hilst/go-ui-html-template/builder/functions"
)

var children map[int]*html.Template

func startChilds() {
	children = make(map[int]*html.Template, 0)
}

func (tb *TemplateBuilder) loadChild(tmplName string, index int) bool {
	startChilds()
	child := tb.parsed.Funcs(fn.Map()).Lookup(tmplName)
	if child == nil {
		return false
	}
	children[index] = child
	return true
}

func (tb *TemplateBuilder) retrieveChild(data any, index int) html.HTML {
	child, ok := children[index]
	if !ok {
		return html.HTML("")
	}
	tmpBf := bytes.NewBuffer([]byte{})
	err := child.Execute(tmpBf, data)
	if err != nil {
		return html.HTML("")
	}
	return html.HTML(tmpBf.String())
}

func clearChildren() html.HTML {
	children = nil
	return html.HTML("")
}
