package builder

import (
	"strconv"

	"github.com/qri-io/jsonpointer"
)

func props(els ...any) []any {
	return els
}

// GET
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

// TYPES
func stringfy(el any) string {
	s, sok := el.(string)
	if sok {
		return s
	}
	f, fok := el.(float64)
	if fok {
		return strconv.FormatFloat(f, 'G', -1, 64)
	}
	return el.(string)
}
func array(parr any) []any {
	return parr.([]any)
}

// WHERE
func where(path string, comp any, array []any) any {
	for _, item := range array {
		base := get(path, item)
		base = stringfy(base)
		comp = stringfy(comp)
		if base == comp {
			return item
		}
	}
	return nil
}

// MASK
func mask(pattern string, value string) string {
	result := ""
	vidx := 0
	for _, rune := range pattern {
		char := string(rune)
		if char == "#" {
			result += string(value[vidx])
			vidx += 1
		} else {
			result += char
		}
	}
	return result
}
