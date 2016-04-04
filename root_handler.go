package main

import (
	"fmt"
	"net/http"

	"github.com/elazarl/go-bindata-assetfs"
)

func rootHandler(prefix string) http.Handler {
	return rootPages{
		prefix: prefix,
		assetHandler: http.FileServer(&assetfs.AssetFS{
			Asset: Asset,
			AssetDir: func(name string) ([]string, error) {
				return nil, fmt.Errorf("Asset %s not found", name)
			},
			AssetInfo: AssetInfo,
			Prefix:    prefix,
		}),
	}
}

type rootPages struct {
	prefix       string
	assetHandler http.Handler
}

func (h rootPages) haveStaticAsset(path string) bool {
	_, found := _bindata[h.prefix+path]
	return found
}

func getIndexHandlerFunc(w http.ResponseWriter, r *http.Request) {
	err := handleAsTemplateFile(w, http.StatusOK, "templates/index.html", struct{ Title string }{AppName})
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
	}
}

func (h rootPages) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" && r.Method == "GET" {
		getIndexHandlerFunc(w, r)
	} else if h.haveStaticAsset(r.URL.Path) {
		h.assetHandler.ServeHTTP(w, r)
	} else {
		notFoundHandlerFunc(w, r)
	}
}
