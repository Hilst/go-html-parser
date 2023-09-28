package functions

import "strings"

func mapDateFormatter(java string) (golang string) {
	golang = strings.ReplaceAll(java, "yyyy", "2006")
	golang = strings.ReplaceAll(golang, "yy", "06")
	golang = strings.ReplaceAll(golang, "MM", "01")
	golang = strings.ReplaceAll(golang, "dd", "02")
	golang = strings.ReplaceAll(golang, "HH", "15")
	golang = strings.ReplaceAll(golang, "hh", "03")
	golang = strings.ReplaceAll(golang, "mm", "04")
	golang = strings.ReplaceAll(golang, "ss", "05")
	golang = strings.ReplaceAll(golang, "SSS", ".000")
	golang = strings.ReplaceAll(golang, "a", "PM")
	return golang
}
