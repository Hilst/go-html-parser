package builder

import (
	"bytes"
	html "html/template"

	"github.com/tdewolff/minify/v2"
	minifyHTML "github.com/tdewolff/minify/v2/html"
)

type JSON = map[string]any

type TemplateBuilder struct {
	data   JSON
	layout string
	all    *html.Template
}

func NewBuilder(data JSON, layout string) TemplateBuilder {
	return TemplateBuilder{
		data,
		layout,
		nil,
	}
}

func (tb *TemplateBuilder) functions() html.FuncMap {
	return html.FuncMap{
		"props":  props,
		"get":    get,
		"string": stringfy,
		"array":  array,
		"where":  where,
		"mask":   mask,
	}
}

func (tb *TemplateBuilder) readyAll() error {
	parsed, err := html.New("ALL").ParseGlob("./templates/**/*.html")
	if err == nil {
		tb.all = parsed
	}
	return err
}

func (tb *TemplateBuilder) parseLayout() error {
	_, err := tb.all.New("LAYOUT").Funcs(tb.functions()).Parse(tb.layout)
	return err
}

func (tb *TemplateBuilder) minifyOutput(bf bytes.Buffer) (bytes.Buffer, error) {
	const mimetype = "text/html"
	m := minify.New()
	m.AddFunc(mimetype, minifyHTML.Minify)
	mini := bytes.NewBuffer([]byte{})
	err := m.Minify(mimetype, mini, &bf)
	if err != nil {
		return bytes.Buffer{}, err
	}
	return *mini, nil
}

func (tb *TemplateBuilder) Build() (string, error) {
	var err error

	err = tb.readyAll()
	if err != nil {
		return "", err
	}

	err = tb.parseLayout()
	if err != nil {
		return "", err
	}

	tmpBf := bytes.NewBuffer([]byte{})
	err = tb.all.ExecuteTemplate(tmpBf, "MAIN", tb.data)
	if err != nil {
		return "", err
	}

	mini, err := tb.minifyOutput(*tmpBf)
	if err != nil {
		return "", err
	}
	return mini.String(), err
}
