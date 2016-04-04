package main

import (
	"net/http"

	"github.com/elazarl/go-bindata-assetfs"
)

func rootHandler() http.Handler {
	return rootPages{
		assets: http.FileServer(&assetfs.AssetFS{
			Asset:     Asset,
			AssetDir:  AssetDir,
			AssetInfo: AssetInfo,
			Prefix:    "public",
		}),
	}
}

type rootPages struct {
	assets http.Handler
}

func (h rootPages) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" && r.Method == "GET" {
		handleAsTemplateFile(w, http.StatusOK, "templates/index.html", struct {
			Title string
		}{
			AppName,
		})
	} else {
		h.assets.ServeHTTP(w, r)
	}
}
