package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println(versionDisplay())
	serveAddress := ""
	port := 3000
	http.ListenAndServe(fmt.Sprintf("%s:%d", serveAddress, port), createRouter())
}
