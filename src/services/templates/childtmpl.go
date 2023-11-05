package templates

import (
	"bytes"
	html "html/template"
)

var children map[int]*html.Template

func (ts *TemplateService) Load(tmplName string, index int) bool {
	if children == nil {
		children = make(map[int]*html.Template)
	}
	child := ts.execute.Lookup(tmplName)
	if child == nil {
		return false
	}
	children[index] = child
	return true
}

func (ts *TemplateService) Retrieve(data any, index int) html.HTML {
	child, ok := children[index]
	tmpBf := bytes.NewBuffer([]byte{})
	err := child.Execute(tmpBf, data)
	if err != nil {
		return html.HTML("")
	}
	if !ok {
		return html.HTML("")
	}
	return html.HTML(tmpBf.String())
}

func Restart() html.HTML {
	children = nil
	return html.HTML("")
}
