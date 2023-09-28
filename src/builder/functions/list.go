package functions

import "strings"

// GENERATORS
func props(els ...any) []any {
	return els
}

func split(sep string, value string) []string {
	return strings.Split(value, sep)
}

// TYPE ENFORCER
func array(parr any) []any {
	return parr.([]any)
}

// REDUCERS
func first(path string, comp any, array []any) any {
	var base string
	for _, item := range array {
		base = stringfy(get(path, item))
		comp = stringfy(comp)
		if base == comp {
			return item
		}
	}
	return nil
}

func filter(path string, comp any, array []any) []any {
	res := make([]any, 0)
	var base string
	for _, item := range array {
		base = stringfy(get(path, item))
		comp = stringfy(comp)
		if base == comp {
			res = append(res, item)
		}
	}
	return res
}
