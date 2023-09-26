package builder

import (
	"bytes"
	html "html/template"

	"github.com/tdewolff/minify/v2"
	minifyHTML "github.com/tdewolff/minify/v2/html"
)

type JSON = map[string]any

type TemplateBuilder struct {
	data JSON
}

var comp *html.Template

func NewBuilder(data JSON) (*TemplateBuilder, error) {
	tb := TemplateBuilder{
		data,
	}

	err := ready()
	if err != nil {
		return nil, err
	}

	return &tb, nil
}

func ready() error {
	if comp != nil {
		return nil
	}

	parsed, err := html.New("ALL").Funcs(functions()).ParseGlob("./templates/**/*.html")
	if err != nil {
		return err
	}

	comp = parsed
	return err
}

func minifyOutput(bf bytes.Buffer) (bytes.Buffer, error) {
	m := minify.New()
	m.AddFunc("text/html", minifyHTML.Minify)

	mini := bytes.NewBuffer([]byte{})

	err := m.Minify("text/html", mini, &bf)
	if err != nil {
		return bytes.Buffer{}, err
	}

	return *mini, nil
}

func functions() html.FuncMap {
	return html.FuncMap{
		"props":  props,
		"get":    get,
		"string": stringfy,
		"array":  array,
		"where":  where,
		"mask":   mask,
	}
}

func (tb *TemplateBuilder) Build(layout string) (string, error) {
	if comp == nil {
		if err := ready(); err != nil {
			return "", err
		}
	}

	clone, err := comp.Clone()
	if err != nil {
		return "", err
	}
	tmpl, err := clone.New("LAYOUT").Funcs(functions()).Parse(layout)
	if err != nil {
		return "", err
	}

	if err != nil {
		return "", err
	}

	tmpBf := bytes.NewBuffer([]byte{})
	err = tmpl.ExecuteTemplate(tmpBf, "MAIN", tb.data)
	if err != nil {
		return "", err
	}

	mini, err := minifyOutput(*tmpBf)
	if err != nil {
		return "", err
	}
	return mini.String(), err
}
