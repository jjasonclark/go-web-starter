package main

import (
	"html/template"
	"net/http"
)

const (
	// HTMLContentType Return type for generated web pages
	HTMLContentType = "text/html;utf-8"
)

type RootHandler struct {
	assets http.Handler
}

func (h *RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" && r.Method == "GET" {
		handleAsTemplateFile(w, r, "templates/index.html", struct {
			Title string
		}{
			AppName,
		})
	} else {
		h.assets.ServeHTTP(w, r)
	}
}

func handleAsTemplateFile(w http.ResponseWriter, r *http.Request, n string, data interface{}) error {
	f, err := Asset(n)
	if err != nil {
		return err
	}
	w.Header().Add("CONTENT-TYPE", HTMLContentType)
	t, err := template.New("anonymous").Parse(string(f))
	if err != nil {
		return err
	}
	return t.Execute(w, data)
}
