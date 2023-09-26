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
var m *minify.M

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

	parsed, err := html.New("ALL").Funcs(functions()).ParseGlob("./res/templates/**/*.html")
	if err != nil {
		return err
	}

	comp = parsed
	m = minify.New()
	m.AddFunc("text/html", minifyHTML.Minify)
	return err
}

func minifyOutput(bf bytes.Buffer) (bytes.Buffer, error) {
	if m == nil {
		if err := ready(); err != nil {
			return bytes.Buffer{}, err
		}
	}
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

func (tb *TemplateBuilder) Build(layouts []string) ([]string, []error) {
	if comp == nil || m == nil {
		if err := ready(); err != nil {
			e := make([]error, 1)
			e[0] = err
			return make([]string, 0), e
		}
	}

	output := make([]string, len(layouts))
	errList := make([]error, len(layouts))
	for index, layout := range layouts {
		texthtml, err := tb.buildPage(layout)
		if err != nil {
			errList[index] = err
			continue
		}
		output[index] = texthtml
	}
	return output, errList
}

func (tb *TemplateBuilder) buildPage(layout string) (string, error) {
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
	return mini.String(), nil
}
