package main

import "flag"

var appConfig struct {
	host   *string
	port   *int
	dbHost *string
	dbPort *int
	dbName *string
}

func init() {
	appConfig.host = flag.String("host", "", "Host interface")
	appConfig.port = flag.Int("port", 3000, "Port to bind to")
	appConfig.dbHost = flag.String("dbhost", "localhost", "DB host interface")
	appConfig.dbPort = flag.Int("dbport", 28015, "DB port")
	appConfig.dbName = flag.String("dbname", "app", "DB name")
	flag.Parse()
}
