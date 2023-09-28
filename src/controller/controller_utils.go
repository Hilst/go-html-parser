package controller

import (
	"net/http"
)

func numberErrors(errs []error) int {
	r := 0
	for _, v := range errs {
		if v == nil {
			r++
		}
	}
	return len(errs) - r
}

func readyStatus(status *int, errs []error, totalSize int) {
	ne := numberErrors(errs)
	if ne > 0 && ne < totalSize {
		*status = http.StatusPartialContent
	}
	if ne == totalSize {
		*status = http.StatusNotFound
	}
}
