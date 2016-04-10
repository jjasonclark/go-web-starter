package main

import (
	"encoding/json"
	"net/http"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func decodeJSONRequest(r *http.Request, formData interface{}) error {
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(formData)
}
