package functions

import (
	html "html/template"

	"github.com/google/uuid"
	"github.com/qri-io/jsonpointer"
)

var funcs = html.FuncMap{
	"array":      array,
	"currency":   currencyFormat,
	"dateformat": dateFormat,
	"decimal":    decimalFormat,
	"filter":     filter,
	"first":      first,
	"float":      float,
	"get":        get,
	"integer":    integer,
	"mask":       mask,
	"now":        now,
	"number":     numberFormat,
	"pad":        pad,
	"percent":    percentFormat,
	"props":      props,
	"split":      split,
	"string":     stringfy,
	"timedate":   timedate,
	"uuid":       genuuid,
}

func Map() html.FuncMap {
	return funcs
}

func Add(key string, fun any) {
	funcs[key] = fun
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

func genuuid() string {
	return uuid.NewString()
}
