package main

import (
	"github.com/getsentry/sentry-go"
	"log"
)

func InitAPM(dsn, env string) {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:           dsn,
		Debug:         true,
		EnableTracing: true,
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,
		Environment:      env,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
}
