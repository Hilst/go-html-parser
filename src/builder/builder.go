package builder

import (
	"bytes"
	html "html/template"

	"github.com/tdewolff/minify/v2"
	minifyHTML "github.com/tdewolff/minify/v2/html"

	fn "github.com/Hilst/go-ui-html-template/builder/functions"
)

type TemplateBuilder struct {
	parsed *html.Template
}

var comp *html.Template
var m *minify.M

func NewBuilder() (*TemplateBuilder, error) {
	tb := TemplateBuilder{}
	err := ready()
	if err != nil {
		return nil, err
	}
	tb.addDepenFuncs()

	return &tb, nil
}

func ready() error {
	if comp != nil {
		return nil
	}

	parsed, err := html.New("ALL").Funcs(fn.Map()).ParseGlob("./res/templates/**/*.html")
	if err != nil {
		return err
	}

	comp = parsed
	m = minify.New()
	m.AddFunc("text/html", minifyHTML.Minify)
	return err
}

func (tb *TemplateBuilder) addDepenFuncs() {
	fn.Add("child", tb.retrieveChild)
	fn.Add("loadchild", tb.loadChild)
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

func (tb *TemplateBuilder) Build(layouts []string, data map[string]any) ([]string, []error) {
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
		texthtml, err := tb.buildPage(layout, data)
		if err != nil {
			errList[index] = err
			continue
		}
		output[index] = texthtml
	}
	return output, errList
}

func (tb *TemplateBuilder) buildPage(layout string, data map[string]any) (string, error) {
	clone, err := comp.Clone()
	if err != nil {
		return "", err
	}
	tmpl, err := clone.New("LAYOUT").Funcs(fn.Map()).Parse(layout)
	if err != nil {
		return "", err
	}
	tb.parsed, err = tmpl.Clone()
	if err != nil {
		return "", err
	}

	tmpBf := bytes.NewBuffer([]byte{})
	err = tmpl.ExecuteTemplate(tmpBf, "MAIN", data)
	if err != nil {
		return "", err
	}

	mini, err := minifyOutput(*tmpBf)
	if err != nil {
		return "", err
	}
	return mini.String(), nil
}
