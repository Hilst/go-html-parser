package templates

import (
	html "html/template"

	"github.com/google/uuid"
	opt "github.com/moznion/go-optional"
	"github.com/qri-io/jsonpointer"
)

var funcs = html.FuncMap{
	"get":  get,
	"uuid": genuuid,

	"array":      array,
	"props":      props,
	"solvearray": resolveArray,
	"emptyarray": emptyArray,

	"integer":  integer,
	"float":    float,
	"currency": currencyFormat,
	"decimal":  decimalFormat,
	"number":   numberFormat,
	"percent":  percentFormat,

	"string":        stringfy,
	"solvestring":   resolveString,
	"invalidstring": invalidString,
	"mask":          mask,
	"pad":           pad,

	"now":        now,
	"timedate":   timedate,
	"dateformat": dateFormat,
}

func FuncMap() html.FuncMap {
	return funcs
}

func AddFunction(key string, fun any) {
	funcs[key] = fun
}

func get(path string, from any) opt.Option[any] {
	p, perr := jsonpointer.Parse(path)
	if perr == nil {
		c, cerr := p.Eval(from)
		if c != nil && cerr == nil {
			return opt.Some[any](c)
		}
	}
	return opt.None[any]()
}

func genuuid() string {
	return uuid.NewString()
}
