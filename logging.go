package main

import (
	"github.com/gorilla/handlers"
	"net/http"
	"os"
)

func HandlersLogging(handler http.Handler) http.Handler {
	return handlers.CombinedLoggingHandler(os.Stdout, handler)
}
