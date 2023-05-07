package main

import (
	"github.com/gorilla/handlers"
	"go.uber.org/zap"
	"net/http"
	"os"
)

var Logger *zap.SugaredLogger

func HandlersLogging(handler http.Handler) http.Handler {
	return handlers.CombinedLoggingHandler(os.Stdout, handler)
}

// ConfigureLogging is used to configure the global logger with or without debugging
func ConfigureLogging(debug bool) (*zap.Logger, error) {
	// configure logging
	logConfig := zap.NewProductionConfig()
	logConfig.OutputPaths = []string{"stdout"}
	logConfig.DisableStacktrace = true
	if debug {
		logConfig.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
		logConfig.DisableStacktrace = false
	}

	logger, _ := logConfig.Build()
	defer logger.Sync()

	return logger, nil
}
