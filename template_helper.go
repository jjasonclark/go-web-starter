package main

import (
	"html/template"
	"mime"
	"net/http"
	"path"
)

const (
	// JSONContentType Return type for JSON responses
	JSONContentType = "application/json;charset=utf-8"
)

func handleAsTemplateFile(w http.ResponseWriter, statusCode int, templatePath string, data interface{}) error {
	t, err := fetchTemplate(templatePath)
	if err != nil {
		return err
	}
	setContentType(w, templatePath)
	w.WriteHeader(statusCode)
	return t.Execute(w, data)
}

func fetchTemplate(templatePath string) (*template.Template, error) {
	f, err := Asset(templatePath)
	if err != nil {
		return nil, err
	}
	return template.New("anonymous").Parse(string(f))
}

func setContentType(w http.ResponseWriter, templatePath string) {
	mimeType := mime.TypeByExtension(path.Ext(templatePath))
	if mimeType != "" {
		w.Header().Set("CONTENT-TYPE", mimeType)
	}
}
