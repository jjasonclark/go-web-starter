package main

import "net/http"

func appHandler() http.Handler {
	return http.HandlerFunc(appHandlerFunc)
}

func appHandlerFunc(w http.ResponseWriter, r *http.Request) {
	handleAsTemplateFile(w, r, "templates/app.html", struct {
		Title string
	}{
		Title: AppName,
	})
}
