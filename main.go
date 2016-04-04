package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	serverAddress := fmt.Sprintf("%s:%d", *appConfig.host, *appConfig.port)

	fmt.Println(versionDisplay())

	fmt.Printf("HTTP server Listening to %s\n", serverAddress)
	if err := http.ListenAndServe(serverAddress, createRouter()); err != nil {
		log.Fatalln(err.Error())
	}
}
