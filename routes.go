package main

import "net/http"

func createRouter() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/app", appHandler())
	mux.Handle("/api/", http.StripPrefix("/api", apiRoutes()))
	mux.Handle("/", rootHandler("public"))

	return middlewareStack(mux)
}

func apiRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", notFoundHandlerFunc)
	return mux
}

func notFoundHandlerFunc(w http.ResponseWriter, r *http.Request) {
	err := handleAsTemplateFile(w, http.StatusNotFound, "templates/404.html", struct{ Title string }{AppName})
	if err != nil {
		http.Error(w, "Template error", http.StatusNotFound)
	}
}
