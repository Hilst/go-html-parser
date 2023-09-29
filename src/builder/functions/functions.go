package functions

import (
	html "html/template"

	"github.com/qri-io/jsonpointer"
)

func Map() html.FuncMap {
	return html.FuncMap{
		"array":      array,
		"currency":   currencyformat,
		"dateformat": dateformat,
		"decimal":    decimalformat,
		"filter":     filter,
		"first":      first,
		"float":      float,
		"get":        get,
		"integer":    integer,
		"mask":       mask,
		"now":        now,
		"number":     numberformat,
		"pad":        pad,
		"percent":    percentformat,
		"props":      props,
		"split":      split,
		"string":     stringfy,
		"timedate":   timedate,
	}
}

func get(path string, from any) any {
	p, perr := jsonpointer.Parse(path)
	if perr == nil {
		c, cerr := p.Eval(from)
		if cerr == nil {
			return c
		}
	}
	return nil
}
