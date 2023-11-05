package templates

import (
	"fmt"
)

// TYPE ENFORCER
func stringfy(el any) string {
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
