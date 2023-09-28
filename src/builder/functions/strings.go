package functions

import (
	"fmt"
	"strconv"
)

// TYPE ENFORCER
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

// TRANSFORMERS
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

func pad(left bool, end int, pattern string, value string) string {
	if end <= len(value) {
		return value
	}

	f := "%"
	if !left {
		f += "-"
	}
	f += fmt.Sprintf("%d", end-len(value))
	f += "s"
	return fmt.Sprintf(f, value)
}
