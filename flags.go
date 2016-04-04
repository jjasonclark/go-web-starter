package main

import "flag"

var appConfig struct {
	host   *string
	port   *int
}

func init() {
	appConfig.host = flag.String("host", "", "Host interface")
	appConfig.port = flag.Int("port", 3000, "Port to bind to")
	flag.Parse()
}
