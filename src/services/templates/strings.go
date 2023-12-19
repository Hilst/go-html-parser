package templates

import (
	"fmt"

	opt "github.com/moznion/go-optional"
)

// TYPE ENFORCER
func stringfy(el opt.Option[any]) opt.Option[string] {
	if el, err := el.Take(); err == nil {
		return opt.Some[string](el.(string))
	}
	return opt.None[string]()
}

// VALIDATOR
const invalid_string = "<invalid>"

func resolveString(op opt.Option[string]) string {
	return op.TakeOr(invalid_string)
}

func invalidString(s string) bool {
	return s == invalid_string
}

// TRANSFORMERS
func mask(pattern string, value opt.Option[string]) opt.Option[string] {
	if value, err := value.Take(); err == nil {
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
		return opt.Some[string](result)
	}
	return opt.None[string]()
}

func pad(left bool, end int, pattern string, value opt.Option[string]) opt.Option[string] {
	if value, err := value.Take(); err == nil {
		if end <= len(value) {
			return opt.Some[string](value)
		}

		f := "%"
		if !left {
			f += "-"
		}
		f += fmt.Sprintf("%d", end-len(value))
		f += "s"
		return opt.Some[string](fmt.Sprintf(f, value))
	}
	return opt.None[string]()
}
