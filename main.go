package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	serverAddress := fmt.Sprintf("%s:%d", *appConfig.host, *appConfig.port)
	dbAddress := fmt.Sprintf("%s:%d", *appConfig.dbHost, *appConfig.dbPort)
	currentDBConfig := dbConfig{
		Address: dbAddress,
		DBName:  *appConfig.dbName,
	}

	fmt.Println(versionDisplay())

	fmt.Printf("Connecting to RethinkDB at %s\n", dbAddress)
	if err := currentDBConfig.dial(); err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Printf("HTTP server Listening to %s\n", serverAddress)
	if err := http.ListenAndServe(serverAddress, createRouter()); err != nil {
		log.Fatalln(err.Error())
	}
}
