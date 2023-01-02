package utils

import "net/http"

func ParseHTTPErrorStatusCode(err error) int {
	switch err {
	default:
		return http.StatusInternalServerError
	}
}
