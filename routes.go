package main

import "net/http"

func createRouter() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/app", appHandler())
	mux.Handle("/", rootHandler())

	return middlewareStack(mux)
}
