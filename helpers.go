package main

import (
	"encoding/json"
	"net/http"
)

type envelope map[string]interface{}

func writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(payload)

	return nil
}
