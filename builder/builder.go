package builder

import (
	"bytes"
	html "html/template"

	"github.com/tdewolff/minify/v2"
	minifyHTML "github.com/tdewolff/minify/v2/html"
)

type JSON = map[string]any

type TemplateBuilder struct {
	data Data
	all  *html.Template
}

type Data struct {
	Json JSON
	Id   string
}

func NewBuilder(data JSON) (*TemplateBuilder, error) {
	tb := TemplateBuilder{
		Data{
			data,
			"",
		},
		nil,
	}
	err := tb.readyAll()
	if err != nil {
		return nil, err
	}
	return &tb, nil
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
	if err != nil {
		return err
	}
	parsed, err = parsed.New("LAYOUTS").Funcs(tb.functions()).ParseGlob("./screens/*.html")
	if err == nil {
		tb.all = parsed
	}
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

func (tb *TemplateBuilder) Build(id string) (string, error) {
	tb.data.Id = id
	tmpBf := bytes.NewBuffer([]byte{})
	err := tb.all.ExecuteTemplate(tmpBf, "MAIN", tb.data)
	if err != nil {
		return "", err
	}

	mini, err := tb.minifyOutput(*tmpBf)
	if err != nil {
		return "", err
	}
	return mini.String(), err
}
