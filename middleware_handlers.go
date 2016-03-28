package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
)

func middlewareStack(next http.Handler) http.Handler {
	return timeoutHandler(loggingHandler(next))
}

func timeoutHandler(next http.Handler) http.Handler {
	return http.TimeoutHandler(next, 60*time.Second, "Service unavailable")
}

func loggingHandler(next http.Handler) http.Handler {
	return handlers.LoggingHandler(os.Stdout, next)
}
