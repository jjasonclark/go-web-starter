package main

import "net/http"

func appHandler() http.Handler {
	return http.HandlerFunc(appPage)
}

func appPage(w http.ResponseWriter, r *http.Request) {
	handleAsTemplateFile(w, r, "templates/app.html", struct {
		Title string
	}{
		Title: AppName,
	})
}
