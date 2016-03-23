package main

import (
	"fmt"
	"net/http"

	"github.com/elazarl/go-bindata-assetfs"
)

func main() {
	fmt.Println(versionDisplay())
	rootHandler := &RootHandler{
		assets: http.FileServer(&assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: AssetInfo, Prefix: "public"}),
	}
	http.Handle("/", rootHandler)
	http.ListenAndServe("0.0.0.0:3000", nil)
}
